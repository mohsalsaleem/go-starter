package example

import "github.com/mohsalsaleem/go-starter/logger"

type Example struct {
}

func NewExample() *Example {
	return &Example{}
}

func (*Example) Run() error {
	logger.Infof("Starting example service")
	return nil
}
