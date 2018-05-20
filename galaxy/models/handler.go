package models

import "errors"

type (
	HandlerRsp struct {
		Context string
		Res     string
		Err     error
	}

	Handler interface {
		Validate(context string, g *Guider) error
		Handle(context string, g *Guider) HandlerRsp
	}
)

var (
	NotMatchErr = errors.New("not match err")
	UnknownErr  = errors.New("I have no idea what you are talking about")
	CalcErr     = errors.New("calc alias to decimal err")
)
