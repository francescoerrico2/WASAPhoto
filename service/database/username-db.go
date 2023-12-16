package database

func (db *appdbimpl) GetUsername(user User) (string, error) {

	var Username string

	err := db.c.QueryRow(`SELECT nickname FROM users WHERE id_user = ?`, user.IdUser).Scan(&Username)
	if err != nil {
		return "", err
	}
	return Username, nil
}

func (db *appdbimpl) ModifyUsername(user User, newUsername Username) error {

	_, err := db.c.Exec(`UPDATE users SET nickname = ? WHERE id_user = ?`, newUsername.Username, user.IdUser)
	if err != nil {

		return err
	}
	return nil
}
