package utils

import (
	"moyu/internal/common"
	"moyu/internal/service"
)

type (
	sExample struct{}
)

func init() {
	service.RegisterExample(Example())
	common.Logs().Info("Init Examples success")
}

var insExample = sExample{}

func Example() *sExample {
	return &insExample
}

func (s *sExample) Example() {

}
