package check

import (
	"context"

	myerr "poj/MyError"
	"poj/middleware"
	"poj/model"
	"strings"
)

func Cleaning(ctx context.Context, u model.UserModel) model.UserModel {

	//name fixing
	u.Name = strings.TrimSpace(u.Name)
	u.Name = strings.Title(strings.ToLower(u.Name))
	//email fixing
	u.Email = strings.ToLower(u.Email)
	u.Email = strings.TrimSpace(u.Email)

	return u

}

func Validation(ctx context.Context, u model.UserModel) error {

	if strings.TrimSpace(u.Name) == "" {
		return &myerr.ErrorInfo{
			Code:    400,
			Sussecc: false,
			Result:  "Error for name",
		}
	}
	if err := middleware.EmailValidtion(u.Email); err != nil {
		return err
	}
	return nil
}
