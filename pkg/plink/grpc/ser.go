package grpc

import (
	"chatcser/pkg/plink/config"
	service "chatcser/pkg/plink/grpc/proto"
	"chatcser/pkg/plink/iface"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

type GrpcService struct {
	Config  *config.Config
	PServer iface.IServer

	Router iface.IRouter

	//当前Server的消息管理模块，用来绑定MsgId和对应的处理方法
	MsgHandler iface.IMsgHandle
	//当前Server的链接管理器
	ConnMgr iface.IConnManager
	//该Server的连接创建时Hook函数
	OnConnStart func(conn iface.IConnection)
	//该Server的连接断开时的Hook函数
	OnConnStop func(conn iface.IConnection)

	server *grpc.Server

	*service.UnimplementedProdServiceServer
}

func NewGrpcServer(s iface.IServer) *GrpcService {
	grpc := &GrpcService{
		Config:     s.GetConfig(),
		PServer:    s,
		Router:     s.GetRouter(),
		MsgHandler: s.GetMsgHandler(),
		ConnMgr:    s.GetConnMgr(),
		server:     grpc.NewServer(),
	}
	return grpc
}

func (p GrpcService) GetProdSocket(stream service.ProdService_GetProdSocketServer) error {

	log.Println("start of stream")
	for {
		//接收客户端的流
		recv, err := stream.Recv()
		if err == io.EOF {
			log.Println("客户端流接收完毕，跳出循环")
			break
		}
		//打印客户端的流
		log.Printf("接收到客户端的流 %s", recv.GetProdId())
		//服务端接收的值打印
		//log.Printf("The prodid is %s", recv.GetProdId())

		//向客户端发送流  (+100  表示服务端逻辑 对值进行处理后返回)
		err = stream.Send(&service.ProductResponse{ProdSocket: recv.ProdId + 100})
		if err != nil {
			log.Fatal(err.Error())
		}

	}

	return nil
}

func Start(g *GrpcService) error {
	// Start gRPC servergo
	lis, err := net.Listen("tcp", g.Config.GrpcConfig.GrpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}
	fmt.Printf("[START] Grpc Server name: %s,listenner at port: %s, is starting\n", g.Config.Name, g.Config.GrpcConfig.GrpcPort)

	service.RegisterProdServiceServer(g.server, g)
	log.Println("Starting gRPC server...")
	if err := g.server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		return err
	}
	return nil
}
