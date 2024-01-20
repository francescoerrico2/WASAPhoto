package database

func (db *appdbimpl) BanUser(banner User, banned User) error {

	_, err := db.c.Exec("INSERT INTO banned_users (banner,banned) VALUES (?, ?)", banner.IdUser, banned.IdUser)
	if err != nil {
		return err
	}

	return nil
}
func (db *appdbimpl) UnbanUser(banner User, banned User) error {

	_, err := db.c.Exec("DELETE FROM banned_users WHERE (banner = ? AND banned = ?)", banner.IdUser, banned.IdUser)
	if err != nil {
		return err
	}

	return nil
}
func (db *appdbimpl) BannedUserCheck(requestingUser User, targetUser User) (bool, error) {

	var cnt int
	err := db.c.QueryRow("SELECT COUNT(*) FROM banned_users WHERE banned = ? AND banner = ?",
		requestingUser.IdUser, targetUser.IdUser).Scan(&cnt)

	if err != nil {
		return true, err
	}
	if cnt == 1 {
		return true, nil
	}
	return false, nil

}
