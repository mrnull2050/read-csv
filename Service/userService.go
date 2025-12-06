package service

import (
	"context"
	err "poj/MyError"
	repo "poj/Repo"
	"poj/check"
	"poj/model"
)

type UserService struct {
	Repo repo.UserRepo
}

func (us *UserService) SaveUser(ctx context.Context, u model.UserModel) (interface{}, error) {
	//1 cleaning
	u = check.Cleaning(ctx, u)
	//2 validation
	if err := check.Validation(ctx, u); err != nil {
		return nil, err
	}
	if err := us.Repo.Save(ctx, u); err != nil {
		return nil, err
	}
	return &err.ErrorInfo{
		Code:    200,
		Sussecc: true,
		Result:  "User Created",
	}, nil

}
