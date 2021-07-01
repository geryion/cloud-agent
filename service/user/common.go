package user

import (
	"cloud-agent/common/db"
	"cloud-agent/common/util"
	"fmt"
)

func Login(user string, passwd string) error {
	//find db info match
	passwd = util.EncryptionPasswd(passwd)
	result := util.VerifyPasswd(user, passwd)
	if !result {
		return fmt.Errorf("用户名或密码错误!")
	}
	return nil
}

func Register(user string, passwd string, passwd2 string) error {
	//Register user if exist
	exist := util.FindUserIsExist(user)
	if !exist {
		return fmt.Errorf("用户已存在!")
	}
	//format check
	err := util.CheckPasswdLevel(user, passwd)
	if err != nil {
		return err
	}
	//passwd and passwd2 is same
	if passwd != passwd2 {
		return fmt.Errorf("两次密码不一致!")
	}
	//add new user into db
	passwd = util.EncryptionPasswd(passwd)
	usr := &db.User{
		Count : user,
		PassWd: passwd,
	}
	err = db.AddUserToDB(usr)
	if err != nil {
		return fmt.Errorf("注册用户失败!")
	}
	return nil
}

func CancelLation(user string, passwd string) error {
	//user passwd check
	format := util.CheckUserAndPasswdFormat(user, passwd)
	if !format {
		return fmt.Errorf("用户名或密码长度错误!")
	}
	//delete user from db
	passwd = util.EncryptionPasswd(passwd)
	err := db.DelUserFromDB(user, passwd)
	if err != nil {
		return fmt.Errorf("注销用户失败!")
	}
	return nil
}

func ChangePasswd(user string, passwd string) error {
	//user passwd check
	err := util.CheckPasswdLevel(user, passwd)
	if err != nil {
		return err
	}
	//todo need another param to confirm

	//change user passwd
	passwd = util.EncryptionPasswd(passwd)
	err = db.UpdatePasswdByUser(user, passwd)
	if err != nil {
		return fmt.Errorf("修改密码失败!")
	}
	return nil
}
