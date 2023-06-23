package service

import "github.com/mohsalsaleem/go-starter/internal/service/example"

type Service interface {
	Run() error
}

func New() Service {
	return example.NewExample()
}
