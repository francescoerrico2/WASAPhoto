package database

func (db *appdbimpl) GetLikesList(requestingUser User, requestedUser User, photo PhotoId) ([]CompleteUser, error) {

	rows, err := db.c.Query("SELECT id_user FROM likes WHERE id_photo = ? AND id_user NOT IN (SELECT banned FROM banned_users WHERE banner = ? OR banner = ?) "+
		"AND id_user NOT IN (SELECT banner FROM banned_users WHERE banned = ?)",
		photo.IdPhoto, requestingUser.IdUser, requestedUser.IdUser, requestingUser.IdUser)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()
	var likes []CompleteUser
	for rows.Next() {
		var user CompleteUser
		err = rows.Scan(&user.IdUser)
		if err != nil {
			return nil, err
		}
		nickname, err := db.GetUsername(User{IdUser: user.IdUser})
		if err != nil {
			return nil, err
		}
		user.Username = nickname

		likes = append(likes, user)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return likes, nil
}

func (db *appdbimpl) LikePhoto(p PhotoId, u User) error {

	_, err := db.c.Exec("INSERT INTO likes (id_photo,id_user) VALUES (?, ?)", p.IdPhoto, u.IdUser)
	if err != nil {
		return err
	}

	return nil
}

func (db *appdbimpl) UnlikePhoto(p PhotoId, u User) error {

	_, err := db.c.Exec("DELETE FROM likes WHERE(id_photo = ? AND id_user = ?)", p.IdPhoto, u.IdUser)
	if err != nil {
		return err
	}

	return nil
}
