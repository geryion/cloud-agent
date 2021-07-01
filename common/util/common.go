package util

import (
	"cloud-agent/common/db"
	"cloud-agent/common/log"
	"crypto/md5"
	"fmt"
	"regexp"
)

func CheckUserAndPasswdFormat(user string, passwd string) bool {
	//user length is allowed range from 6 to 14, passwd length is allowed range from 8 to 20.
	if (len(user) > 6 && len(user) < 14) && (len(passwd) > 8 && len(passwd) < 20) {
		return true
	}
	return false
}

func CheckPasswdLevel(user, passwd string) error {
	format := CheckUserAndPasswdFormat(user, passwd)
	if !format {
		return fmt.Errorf("用户名或密码长度错误!")
	}
	//password level confirm
	num := `[0-9]{1}`
	a_z := `[a-z]{1}`
	A_Z := `[A-Z]{1}`
	symbol := `[!@#$%^&*()_+/]{1}`
	if b, err := regexp.MatchString(num, passwd); !b || err != nil {
		return fmt.Errorf("请至少输入一个数字!")
	}
	if b, err := regexp.MatchString(a_z, passwd); !b || err != nil {
		return fmt.Errorf("请至少输入一个小写字母!")
	}
	if b, err := regexp.MatchString(A_Z, passwd); !b || err != nil {
		return fmt.Errorf("请至少输入一个大写字母!")
	}
	if b, err := regexp.MatchString(symbol, passwd); !b || err != nil {
		return fmt.Errorf("请至少输入一个特殊字符!")
	}
	return nil
}

func VerifyPasswd(user string, passwd string) bool {
	//To confirm passwd and user
	usr, err := db.FindUserByUserName(user)
	if err != nil {
		return false
	}
	if passwd == usr.PassWd {
		return true
	}
	return false
}

func EncryptionPasswd(passwd string) string {
	md5sum := md5.Sum([]byte(passwd + SALT))
	return fmt.Sprintf("%x", md5sum)
}

func FindUserIsExist(user string) bool {
	usr, err := db.FindUserByUserName(user)
	if err != nil {
		log.CouldLog.Info("" + err.Error())
		return false
	}
	if usr.Count != "" {
		log.CouldLog.Info(usr.Count)
		return false
	}
	return true
}
