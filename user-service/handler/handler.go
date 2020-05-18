package handler

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	"github.com/suxiangdong/laracom/user-service/repo"

	pb "github.com/suxiangdong/laracom/user-service/proto/user"
)

type UserService struct {
	Repo repo.UserRepository
}

func (srv *UserService) Create(ctx context.Context, req *pb.User, rsp *pb.Response) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	req.Password = string(hashedPass)

	if err := srv.Repo.Create(req); err != nil {
		return err
	}

	return nil
}

func (srv *UserService) Get(ctx context.Context, req *pb.User, rsp *pb.Response) error {
	user, err := srv.Repo.Get(req.Id)

	if err != nil {
		return err
	}

	rsp.User = user

	return nil
}

func (srv *UserService) GetAll(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	users, err := srv.Repo.GetAll()

	if err != nil {
		return err
	}

	rsp.Users = users
	return nil
}