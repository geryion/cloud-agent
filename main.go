package main

import (
	"cloud-agent/common/db"
	"cloud-agent/common/log"
	"cloud-agent/controller/server"
	"cloud-agent/test"
	"flag"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"os"
)

func Usage() {
	fmt.Fprint(os.Stdout, "Usage of ", os.Args[0], ":\n")
	flag.PrintDefaults()
	fmt.Fprint(os.Stdout, "\n")
}

func initServer()  {
	//command line parse params
	flag.Usage = Usage
	//test param
	t := flag.Bool("T", false, "Run test")
	vip := flag.String("vip", "172.30.21.6", "Set vip")
	addr := flag.String("addr", "0.0.0.0:8448", "Set addr")
	mysql := flag.Int("mysql", 3306, "Set mysql port")
	flag.Parse()

	if *t {
		test.RunTest(flag.Args(), flag.NArg())
		return
	}
	log.InitLog()
	db.InitDB(*vip, *mysql)

	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	var transportFactory thrift.TTransportFactory
	transportFactory = thrift.NewTTransportFactory()

	if err := server.RunServer(transportFactory, protocolFactory, *addr, false); err != nil {
		fmt.Println("error running server:", err)
	}
	Usage()
}

func main()  {
	initServer()
}
