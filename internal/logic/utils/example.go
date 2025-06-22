package utils

import (
	"scaffold/internal/service"
)

type (
	sExample struct{}
)

func init() {
	service.RegisterExample(Example())
	service.Logs().Info("Init Examples success")
}

var insExample = sExample{}

func Example() *sExample {
	return &insExample
}

func (s *sExample) Example() {

}
