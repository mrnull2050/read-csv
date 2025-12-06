package middleware

import (
	"poj/MyError"
	"regexp"
)

var Reg = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)


func EmailValidtion(email string) error{
	if !Reg.MatchString(email){
		return &err.ErrorInfo{
			Code: 400,
			Sussecc: false,
			Result: "Email valdtion faild",
		}
	}
	return nil
}