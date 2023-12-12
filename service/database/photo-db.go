package database

func (db *appdbimpl) GetPhotosList(requestingUser User, targetUser User) ([]Photo, error) {

	rows, err := db.c.Query("SELECT * FROM photos WHERE id_user = ? ORDER BY date DESC", targetUser.IdUser)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var photos []Photo
	for rows.Next() {
		var photo Photo
		err = rows.Scan(&photo.PhotoId, &photo.Owner, &photo.Date)
		if err != nil {
			return nil, err
		}

		comments, err := db.GetCompleteCommentsList(requestingUser, targetUser, PhotoId{IdPhoto: int64(photo.PhotoId)}) // Old: GetCommentsLen
		if err != nil {
			return nil, err
		}
		photo.Comments = comments

		likes, err := db.GetLikesList(requestingUser, targetUser, PhotoId{IdPhoto: int64(photo.PhotoId)}) // Old: GetLikesLen
		if err != nil {
			return nil, err
		}
		photo.Likes = likes

		photos = append(photos, photo)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return photos, nil
}
func (db *appdbimpl) GetPhoto(requestinUser User, targetPhoto PhotoId) (Photo, error) {

	var photo Photo
	err := db.c.QueryRow("SELECT * FROM photos WHERE (id_photo = ?) AND id_user NOT IN (SELECT banner FROM banned_user WHERE banned = ?)",
		targetPhoto.IdPhoto, requestinUser.IdUser).Scan(&photo)

	if err != nil {
		return Photo{}, ErrUserBanned
	}

	return photo, nil

}
func (db *appdbimpl) CreatePhoto(p Photo) (int64, error) {

	res, err := db.c.Exec("INSERT INTO photos (id_user,date) VALUES (?,?)",
		p.Owner, p.Date)

	if err != nil {
		return -1, err
	}

	photoId, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	return photoId, nil
}

func (db *appdbimpl) RemovePhoto(owner User, p PhotoId) error {

	_, err := db.c.Exec("DELETE FROM photos WHERE id_user = ? AND id_photo = ? ",
		owner.IdUser, p.IdPhoto)
	if err != nil {
		return err
	}

	return nil
}

func (db *appdbimpl) CheckPhotoExistence(targetPhoto PhotoId) (bool, error) {

	var rows int
	err := db.c.QueryRow("SELECT COUNT(*) FROM photos WHERE (id_photo = ?)", targetPhoto.IdPhoto).Scan(&rows)
	if err != nil {
		return false, err
	}

	if rows == 0 {
		return false, nil
	}
	return true, nil

}
