package database

func (db *appdbimpl) GetFollowers(requestinUser User) ([]User, error) {

	rows, err := db.c.Query("SELECT follower FROM followers WHERE followed = ?", requestinUser.IdUser)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()
	var followers []User
	for rows.Next() {
		var folower User
		err = rows.Scan(&folower.IdUser)
		if err != nil {
			return nil, err
		}
		followers = append(followers, folower)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return followers, nil
}
func (db *appdbimpl) GetFollowing(requestinUser User) ([]User, error) {

	rows, err := db.c.Query("SELECT followed FROM followers WHERE follower = ?", requestinUser.IdUser)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()
	var following []User
	for rows.Next() {
		var folowed User
		err = rows.Scan(&folowed.IdUser)
		if err != nil {
			return nil, err
		}
		following = append(following, folowed)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return following, nil
}

func (db *appdbimpl) FollowUser(follower User, followed User) error {

	_, err := db.c.Exec("INSERT INTO followers (follower,followed) VALUES (?, ?)",
		follower.IdUser, followed.IdUser)
	if err != nil {
		return err
	}

	return nil
}

func (db *appdbimpl) UnfollowUser(follower User, followed User) error {

	_, err := db.c.Exec("DELETE FROM followers WHERE(follower = ? AND followed = ?)",
		follower.IdUser, followed.IdUser)
	if err != nil {
		return err
	}

	return nil
}
