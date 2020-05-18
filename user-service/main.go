package main

import (
	"log"

	"github.com/suxiangdong/laracom/user-service/handler"

	"github.com/micro/go-micro"

	repository "github.com/suxiangdong/laracom/user-service/repo"

	pb "github.com/suxiangdong/laracom/user-service/proto/user"

	database "github.com/suxiangdong/laracom/user-service/db"
)

func main() {
	db, err := database.CreateConnection()
	defer func() {
		err = db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	db.AutoMigrate(&pb.User{})

	repo := repository.UserRepository{Db: db}

	service := micro.NewService(micro.Name("laracom.user.service"), micro.Version("v1"))

	service.Init()

	pb.RegisterUserServiceHandler(service.Server(), &handler.UserService{Repo: repo})

	err = service.Run()

	if err != nil {
		log.Fatalf("service start failed")
	}
}
