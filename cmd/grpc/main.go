package main

import (
	service "chatcser/pkg/plink/grpc/proto"
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	//建立无认证链接
	dial, err := grpc.Dial(":8986", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer dial.Close()

	//客户端实例化链接
	client := service.NewProdServiceClient(dial)
	//调用服务端的方法
	stream, _ := client.GetProdSocket(context.Background())
	// 创建goroutine用来向stream中发送message
	ch := make(chan int32, 5)

	go func() {
		prodID := []int32{1, 2, 3, 4, 5}
		for _, v := range prodID {
			log.Printf("客户端发送给服务端的prodid %d\n", v)
			ch <- v
			//发送流到服务端
			_ = stream.Send(&service.ProductRequest{ProdId: v})
			time.Sleep(time.Second)
		}
		// 调用指定次数后主动关闭流
		err = stream.CloseSend()
		if err != nil {
			log.Println(err.Error())
		}
		fmt.Println("客户端结束发送")
		close(ch)
	}()

	// 从stream中接收message
	for {
		prodIDResult, ok := <-ch
		if ok == false {
			log.Println("在通道里取完了数据")
			break
		}
		prodSocket, err := stream.Recv()
		if err == io.EOF {
			log.Println("客户端结束接收，跳出循环")
			break
		}
		log.Printf("客户端发送的值是%d,服务端处理过的值%s \n", prodIDResult, prodSocket.GetProdSocket())
	}

}
