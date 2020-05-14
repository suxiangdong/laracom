package main

import (
	"context"
	"log"

	"github.com/micro/go-micro"
	"github.com/suxiangdong/laracom/demo-service/proto/demo"
)

func main() {
	service := micro.NewService(micro.Name("laracom.demo.cli"))

	service.Init()

	client := demo.NewDemoServiceClient("laracom", service.Client())

	rsp, err := client.SayHello(context.TODO(), &demo.DemoRequest{Name: "sxd"})
	if err != nil {
		log.Fatal("请求失败: ", err)
		return
	}

	log.Println(rsp.Text)
}
