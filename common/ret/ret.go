package ret

const (
	LOGINSUCCESS = "登录成功"
	LOGINFAILED  = "登录失败"

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
