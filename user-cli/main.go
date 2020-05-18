package main

import (
	"log"
	"os"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
	pb "github.com/suxiangdong/laracom/user-service/proto/user"
	"golang.org/x/net/context"
)

func main() {
	service := micro.NewService(
		micro.Flags(
			cli.StringFlag{
				Name:  "name",
				Usage: "Your Name",
			},
			cli.StringFlag{
				Name:  "email",
				Usage: "Your Email",
			},
			cli.StringFlag{
				Name:  "password",
				Usage: "Your Password",
			},
		),
	)

	client := pb.NewUserServiceClient("laracom.user.cli", service.Client())

	service.Init(
		micro.Action(func(c *cli.Context) {
			name := c.String("name")
			email := c.String("email")
			password := c.String("password")

			log.Println("参数:", name, email, password)

			// 调用用户服务
			r, err := client.Create(context.TODO(), &pb.User{
				Name:     name,
				Email:    email,
				Password: password,
			})
			if err != nil {
				log.Fatalf("创建用户失败: %v", err)
			}
			log.Printf("创建用户成功: %s", r.User.Id)

			getAll, err := client.GetAll(context.Background(), &pb.Request{})
			if err != nil {
				log.Fatalf("获取所有用户失败: %v", err)
			}
			for _, v := range getAll.Users {
				log.Println(v)
			}
			os.Exit(0)
		}),
	)

	if err := service.Run(); err != nil {
		log.Fatalf("用户客户端启动失败: %v", err)
	}
}