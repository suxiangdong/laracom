package main

import (
	"context"
	"log"

	"github.com/micro/go-micro"
	pb "github.com/suxiangdong/laracom/demo-service/proto/demo"
)

type S struct {
}

func (s *S) SayHello(ctx context.Context, req *pb.DemoRequest, rsp *pb.DemoResponse) error {
	rsp.Text = "hello " + req.Name
	return nil
}

func main() {
	service := micro.NewService(micro.Name("laracom"))
	service.Init()

	pb.RegisterDemoServiceHandler(service.Server(), &S{})

	if err := service.Run(); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}
