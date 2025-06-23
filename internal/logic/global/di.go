package global

import (
	"scaffold/internal/service"
)

// todo 在这个包里完成依赖注入

type Container struct {
	IMiddleware service.IMiddleware
	// ……其他逻辑层接口
}
