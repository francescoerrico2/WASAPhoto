package database

func (db *appdbimpl) GetStream(user User) ([]Photo, error) {

	rows, err := db.c.Query(`SELECT * FROM photos WHERE id_user IN (SELECT followed FROM followers WHERE follower = ?) ORDER BY date DESC`,
		user.IdUser)
	if err != nil {
		return nil, err
	}

	defer func() { _ = rows.Close() }()

	var res []Photo
	for rows.Next() {
		var photo Photo
		err = rows.Scan(&photo.PhotoId, &photo.Owner, &photo.Date)
		if err != nil {
			return nil, err
		}
		res = append(res, photo)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return res, nil
}

func (db *appdbimpl) CreateUser(u User) error {

	_, err := db.c.Exec("INSERT INTO users (id_user,nickname) VALUES (?, ?)",
		u.IdUser, u.IdUser)

	if err != nil {
		return err
	}

	return nil
}

func (db *appdbimpl) CheckUser(targetUser User) (bool, error) {

	var cnt int
	err := db.c.QueryRow("SELECT COUNT(*) FROM users WHERE id_user = ?",
		targetUser.IdUser).Scan(&cnt)

	if err != nil {

		return true, err
	}

	if cnt == 1 {
		return true, nil
	}
	return false, nil
}
