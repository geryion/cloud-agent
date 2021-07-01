package test

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

func closeTransport(transport thrift.TTransport)  {
	err := recover()
	if err != nil {
		fmt.Println(err)
	}
	if transport != nil {
		transport.Close()
	}
}
