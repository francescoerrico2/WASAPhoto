package database

func (db *appdbimpl) GetCompleteCommentsList(requestingUser User, requestedUser User, photo PhotoId) ([]CompleteComment, error) {

	rows, err := db.c.Query("SELECT * FROM comments WHERE id_photo = ? AND id_user NOT IN (SELECT banned FROM banned_users WHERE banner = ? OR banner = ?) "+
		"AND id_user NOT IN (SELECT banner FROM banned_users WHERE banned = ?)",
		photo.IdPhoto, requestingUser.IdUser, requestedUser.IdUser, requestingUser.IdUser)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()
	var comments []CompleteComment
	for rows.Next() {
		var comment CompleteComment
		err = rows.Scan(&comment.IdComment, &comment.IdPhoto, &comment.IdUser, &comment.Comment)
		if err != nil {
			return nil, err
		}
		nickname, err := db.GetNickname(User{IdUser: comment.IdUser})
		if err != nil {
			return nil, err
		}
		comment.Nickname = nickname

		comments = append(comments, comment)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return comments, nil
}
func (db *appdbimpl) CommentPhoto(p PhotoId, u User, c Comment) (int64, error) {

	res, err := db.c.Exec("INSERT INTO comments (id_photo,id_user,comment) VALUES (?, ?, ?)",
		p.IdPhoto, u.IdUser, c.Comment)
	if err != nil {
		return -1, err
	}

	commentId, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	return commentId, nil
}
func (db *appdbimpl) UncommentPhoto(p PhotoId, u User, c CommentId) error {

	_, err := db.c.Exec("DELETE FROM comments WHERE (id_photo = ? AND id_user = ? AND id_comment = ?)",
		p.IdPhoto, u.IdUser, c.IdComment)
	if err != nil {
		return err
	}

	return nil
}
func (db *appdbimpl) UncommentPhotoAuthor(p PhotoId, c CommentId) error {

	_, err := db.c.Exec("DELETE FROM comments WHERE (id_photo = ? AND id_comment = ?)",
		p.IdPhoto, c.IdComment)
	if err != nil {
		return err
	}

	return nil
}
