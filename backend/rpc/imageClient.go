package rpc

import (
	"identify/backend/common"
	"identify/backend/rpc/image"
	"net"

	"git.apache.org/thrift.git/lib/go/thrift"
	"golang.org/x/net/context"
)

type Image struct {
	Host string
	Port string
}

func (rpc *Image) Identify(appId, imgPath string, ptype common.ProjectType) (uniqueId string, err error) {
	// 定义协议
	transportFactory := thrift.NewTTransportFactory()
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transport, err := thrift.NewTSocket(net.JoinHostPort(rpc.Host, rpc.Port))
	if err != nil {
		return
	}

	useTransport, err := transportFactory.GetTransport(transport)
	// 创建client
	client := image.NewImageServiceClientFactory(useTransport, protocolFactory)
	if err = transport.Open(); err != nil {
		return
	}
	defer transport.Close()
	r := &image.Request{
		ImgPath:     imgPath,
		ProjectType: int32(ptype),
		App:         appId,
	}
	// 根据request获取response
	resp, _ := client.GetIdentify(context.Background(), r)
	uniqueId = resp.BookID
	return uniqueId, nil
}
