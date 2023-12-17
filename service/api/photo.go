package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUserPhotos(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	http.ServeFile(w, r, filepath.Join(photoFolder, ps.ByName("id"), "photos", ps.ByName("photo_id")))
}
func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "applicatio/json")
	auth := extractBearer(r.Header.Get("Authorization"))
	valid := validateRequestingUser(ps.ByName("id"), auth)
	if valid != 0 {
		w.WriteHeader(valid)
		return
	}

	photo := Photo{Owner: auth, Date: time.Now().UTC()}
	data, err := io.ReadAll(r.Body)
	if err != nil {
		ctx.Logger.WithError(err).Error("photo-upload: error reading body content")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	r.Body = io.NopCloser(bytes.NewBuffer(data))
	err = checkFormatPhoto(r.Body, io.NopCloser(bytes.NewBuffer(data)), ctx)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.WithError(err).Error("photo-upload: body contains file that is neither jpg or png")
		_ = json.NewEncoder(w).Encode(JSONErrorMsg{Message: IMG_FORMAT_ERROR_MSG})
		return
	}

	r.Body = io.NopCloser(bytes.NewBuffer(data))
	photoIdInt, err := rt.db.CreatePhoto(photo.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("photo-upload: error executing db function call")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	photoId := strconv.FormatInt(photoIdInt, 10)

	PhotoPath, err := getUserPhotoFolder(auth)
	if err != nil {
		ctx.Logger.WithError(err).Error("photo-upload: error getting user's photo folder")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	out, err := os.Create(filepath.Join(PhotoPath, photoId))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("photo-upload: error creating local photo file")
		return
	}

	_, err = io.Copy(out, r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("photo-upload: error copying body content into file photo")
		return
	}

	out.Close()

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(Photo{
		Comments: nil,
		Likes:    nil,
		Owner:    photo.Owner,
		Date:     photo.Date,
		PhotoId:  int(photoIdInt),
	})

}

func checkFormatPhoto(body io.ReadCloser, newReader io.ReadCloser, ctx reqcontext.RequestContext) error {

	_, errJpg := jpeg.Decode(body)
	if errJpg != nil {

		body = newReader
		_, errPng := png.Decode(body)
		if errPng != nil {
			return errors.New(IMG_FORMAT_ERROR_MSG)
		}
		return nil
	}
	return nil
}

func getUserPhotoFolder(user_id string) (UserPhotoFoldrPath string, err error) {

	photoPath := filepath.Join(photoFolder, user_id, "photos")

	return photoPath, nil

}

func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	bearerAuth := extractBearer(r.Header.Get("Authorization"))
	photoIdStr := ps.ByName("photo_id")
	valid := validateRequestingUser(ps.ByName("id"), bearerAuth)
	if valid != 0 {
		w.WriteHeader(valid)
		return
	}

	photoInt, err := strconv.ParseInt(photoIdStr, 10, 64)
	if err != nil {
		ctx.Logger.WithError(err).Error("photo-delete/ParseInt: error converting photoId to int")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = rt.db.RemovePhoto(
		User{IdUser: bearerAuth}.ToDatabase(),
		PhotoId{IdPhoto: photoInt}.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("photo-delete/RemovePhoto: error coming from database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	pathPhoto, err := getUserPhotoFolder(bearerAuth)
	if err != nil {
		ctx.Logger.WithError(err).Error("photo-delete/getUserPhotoFolder: error with directories")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = os.Remove(filepath.Join(pathPhoto, photoIdStr))
	if err != nil {
		ctx.Logger.WithError(err).Error("photo-delete/os.Remove: photo to be removed is missing")
	}

	w.WriteHeader(http.StatusNoContent)
}
