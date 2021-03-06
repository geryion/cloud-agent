package db

func FindUserByUserName(username string) (*User, error) {
	user := &User{}
	var err error
	err = CloudDB.Where("count=?", username).Find(user).Error
	return user, err
}

func AddUserToDB(user *User) error {
	var err error
	err = CloudDB.Create(user).Error
	return err
}

func DelUserFromDB(user string, passwd string) error {
	var err error
	err = CloudDB.Where("count=? AND pass_wd=?", user, passwd).Delete(&User{}).Error
	return err
}

func UpdatePasswdByUser(user string, passwd string) error {
	var err error
	err = CloudDB.Where("count=?", user).Update("pass_wd", passwd).Error
	return err
}
