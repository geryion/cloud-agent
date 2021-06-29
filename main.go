package main

import (
	"cloud-agent/common/db"
	"cloud-agent/common/log"
	"cloud-agent/controller/server"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

func initServer()  {
	log.InitLog()
	db.InitDB()
	//命令行参数启动服务在本地

	addr := "0.0.0.0:8448"

	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	var transportFactory thrift.TTransportFactory
	transportFactory = thrift.NewTTransportFactory()

	if err := server.RunServer(transportFactory, protocolFactory, addr, false); err != nil {
		fmt.Println("error running server:", err)
	}
}

func main()  {
	initServer()
}
