package database

func (db *appdbimpl) GetNickname(user User) (string, error) {

	var nickname string

	err := db.c.QueryRow(`SELECT nickname FROM users WHERE id_user = ?`, user.IdUser).Scan(&nickname)
	if err != nil {
		return "", err
	}
	return nickname, nil
}

func (db *appdbimpl) ModifyNickname(user User, newNickname Nickname) error {

	_, err := db.c.Exec(`UPDATE users SET nickname = ? WHERE id_user = ?`, newNickname.Nickname, user.IdUser)
	if err != nil {

		return err
	}
	return nil
}
