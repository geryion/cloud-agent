package handler

import (
	"cloud-agent/common/log"
	"cloud-agent/common/ret"
	user2 "cloud-agent/service/user"
	"fmt"
)

func (cc *CloudHandler) Login(user string, passwd string) (string, error) {
	prefix := fmt.Sprintf("CLoud->Login: %v.", user)
	log.CouldLog.Info(prefix)
	retc := &ret.RetCommon{}
	err := user2.Login(user, passwd)
	if err != nil {
		retc.Failed(err.Error())
	} else {
		retc.Success(ret.LOGINSUCCESS)
	}
	return ret.Ret2Json(retc)
}

/*注册接口handler*/
func (cc *CloudHandler) Register(user string, passwd string, passwd2 string) (string, error) {
	prefix := fmt.Sprintf("CLoud->Register: %v.", user)
	log.CouldLog.Info(prefix)
	retc := &ret.RetCommon{}
	err := user2.Register(user, passwd, passwd2)
	if err != nil {
		retc.Failed(err.Error())
	} else {
		retc.Success(ret.REGISTERSUCCESS)
	}
	return ret.Ret2Json(retc)
}

/*注销接口handler*/
func (cc *CloudHandler) CancelLation(user string, passwd string) (string, error) {
	prefix := fmt.Sprintf("CLoud->CancelLation: %v.", user)
	log.CouldLog.Info(prefix)
	retc := &ret.RetCommon{}
	err := user2.CancelLation(user, passwd)
	if err != nil {
		retc.Failed(err.Error())
	} else {
		retc.Success(ret.CANCELLATIONSUCCESS)
	}
	return ret.Ret2Json(retc)
}

//todo change passwd need to confirm and check
func (cc *CloudHandler) ChangePasswd(user string, passwd string) (string, error) {
	prefix := fmt.Sprintf("CLoud->ChangePasswd: %v.", passwd)
	log.CouldLog.Info(prefix)
	retc := &ret.RetCommon{}
	err := user2.ChangePasswd(user, passwd)
	if err != nil {
		retc.Failed(err.Error())
	} else {
		retc.Success(ret.CHANGEPASSWDSUCCESS)
	}
	return ret.Ret2Json(retc)
}

func (cc *CloudHandler) ResetPasswd(Type string, info string, code string) (string, error) {
	//todo select confirm and reset user's passwd
	return "", nil
}