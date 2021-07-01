package ret

import (
	"cloud-agent/common/log"
	"encoding/json"
	"fmt"
)

const (
	LOGINSUCCESS = "登录成功"
	LOGINFAILED  = "登录失败"
	REGISTERSUCCESS = "注册成功"
	REGISTERFAILED  = "注册失败"
	CANCELLATIONSUCCESS = "注销成功"
	CANCELLATIONFAILED  = "注销失败"
	CHANGEPASSWDSUCCESS = "修改密码成功"
	CHANGEPASSWDFAILED  = "修改密码失败"
)

const (
	SUCCESS = "0"
	FAILED  = "1"
)

type RetCommon struct {
	Status 	 	string
	Message 	string
}

func (retc *RetCommon) Success(msg string) {
	retc.Status = SUCCESS
	retc.Message= msg
}

func (retc *RetCommon) Failed(msg string) {
	retc.Status = FAILED
	retc.Message= msg
}

func Ret2Json(ret interface{}) (string, error) {
	retjson, err := json.Marshal(ret)
	if err != nil {
		log.CouldLog.Error(fmt.Sprintf("Ret 2 Json Failed. Err: %v.", err.Error()))
		return "", err
	}
	return string(retjson), nil
}