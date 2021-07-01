package client

import (
	"cloud-agent/common/log"
	"cloud-agent/controller/gen-go/cloud"
	"crypto/tls"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"time"
)

func Client(addr string, secure bool) (*cloud.CloudClient, thrift.TTransport, error) {
	return ClientWithTimeout(addr, 0, secure)
}

func ClientWithTimeout(addr string, timeout time.Duration, secure bool) (*cloud.CloudClient, thrift.TTransport, error) {
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transportFactory := thrift.NewTTransportFactory()
	var transport thrift.TTransport
	var err error
	if secure {
		cfg := new(tls.Config)
		cfg.InsecureSkipVerify = true
		transport, err = thrift.NewTSSLSocket(addr, cfg)
	} else {
		transport, err = thrift.NewTSocketTimeout(addr, timeout)
	}
	if err != nil {
		log.CouldLog.Error(fmt.Sprintf("Error opening socket: %v.", err))
		return nil, nil, err
	}
	transport = transportFactory.GetTransport(transport)
	if err := transport.Open(); err != nil {
		return nil, transport, err
	}
	return cloud.NewCloudClientFactory(transport, protocolFactory), transport, nil
}