package test

import (
	"cloud-agent/controller/client"
	"cloud-agent/controller/gen-go/cloud"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"os"
)

//todo 接口测试

var cc *cloud.CloudClient
var ts thrift.TTransport

func RunTest(args []string, nargs int)  {
	var err error
	cc, ts, err = client.Client("172.30.21.6:8448", false)
	defer closeTransport(ts)
	if err != nil {
		fmt.Println(fmt.Sprintf("Init test client failed. Err: %v.", err))
		os.Exit(1)
	}
	switch nargs - 1 {
	case 0:
		runTestArg0(args)
	case 1:
		runTestArg1(args)
	case 2:
		runTestArg2(args)
	case 3:
		runTestArg3(args)
	default:
		Usage()
	}
}

func runTestArg0(args []string)  {

}

func runTestArg1(args []string)  {

}

func runTestArg2(args []string)  {
	switch args[0] {
	case LOGIN:
		TestLogin(args[1], args[2])
	case CHANGEPASSWD:
		TestChangePasswd(args[1], args[2])
	case CANCELLATION:
		TestCancelLation(args[1], args[2])
	default:
		Usage()
	}
}

func runTestArg3(args []string)  {
	switch args[0] {
	case REGISTER:
		TestRegister(args[1], args[2], args[3])
	default:
		Usage()
	}
}

