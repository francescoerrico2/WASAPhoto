package database

func (db *appdbimpl) GetUsername(user User) (string, error) {

	var username string

	err := db.c.QueryRow(`SELECT username FROM users WHERE id_user = ?`, user.IdUser).Scan(&username)
	if err != nil {
		return "", err
	}
	return username, nil
}

func (db *appdbimpl) ModifyUsername(user User, newUsername Username) error {

	_, err := db.c.Exec(`UPDATE users SET username = ? WHERE id_user = ?`, newUsername.Username, user.IdUser)
	if err != nil {

		return err
	}
	return nil
}
