package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type ServiceConfiguration struct {
}

var ServiceConfig ServiceConfiguration

func Init() error {
	err := envconfig.Process("", &ServiceConfig)
	if err != nil {
		return fmt.Errorf("could not initialize configurations - %s", err.Error())
	}

	return nil
}
