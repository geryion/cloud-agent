package server

import (
	"cloud-agent/controller/gen-go/cloud"
	"cloud-agent/controller/handler"
	"crypto/tls"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

func RunServer(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string, secure bool) error {
	var transport thrift.TServerTransport
	var err error
	if secure {
		//TODO SSL待测试
		cfg := new(tls.Config)
		if cert, err := tls.LoadX509KeyPair("server.crt", "server.key"); err == nil {
			cfg.Certificates = append(cfg.Certificates, cert)
		} else {
			return err
		}
		transport, err = thrift.NewTSSLServerSocket(addr, cfg)
	} else {
		transport, err = thrift.NewTServerSocket(addr)
	}

	if err != nil {
		return err
	}
	fmt.Printf("%T\n", transport)

	cloud_handler := handler.NewCloudHandler()
	cloud_processor := cloud.NewCloudProcessor(cloud_handler)

	server := thrift.NewTSimpleServer4(cloud_processor, transport, transportFactory, protocolFactory)

	fmt.Println("Starting the simple server... on ", addr)
	return server.Serve()
}
