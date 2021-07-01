package test

import (
	"fmt"
)

func TestLogin(user, passwd string)  {
	fmt.Println(fmt.Sprintf("Testing Login"))
	r, err := cc.Login(user, passwd)
	fmt.Println(fmt.Sprintf("Ret: %v \nErr: %v", r, err))
}

func TestRegister(user, passwd, passwd2 string)  {
	fmt.Println(fmt.Sprintf("Testing Register"))
	r, err := cc.Register(user, passwd, passwd2)
	fmt.Println(fmt.Sprintf("Ret: %v \nErr: %v", r, err))
}

func TestChangePasswd(user, passwd string) {
	fmt.Println(fmt.Sprintf("Testing ChangePasswd"))
	r, err := cc.ChangePasswd(user, passwd)
	fmt.Println(fmt.Sprintf("Ret: %v \nErr: %v", r, err))
}

func TestCancelLation(user, passwd string)  {
	fmt.Println(fmt.Sprintf("Testing CancelLation"))
	r, err := cc.CancelLation(user, passwd)
	fmt.Println(fmt.Sprintf("Ret: %v \nErr: %v", r, err))
}