package middleware

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"net/http"
	"scaffold/internal/service"
	"strconv"
	"strings"
)

type (
	sMiddleware struct{}
)

func init() {
	service.RegisterMiddleware(Middleware())
}

var insMiddleware = sMiddleware{}

func Middleware() *sMiddleware {
	return &insMiddleware
}

// DefaultHandlerResponse is the default implementation of HandlerResponse.

type DefaultHandlerResponse struct {
	Message   string      `json:"message"`
	Result    interface{} `json:"result"`
	Status    int         `json:"status"`
	Code      string      `json:"code"`
	Timestamp int64       `json:"timestamp"`
}

// MiddlewareHandlerResponse is the default middleware handling handler response object and its error.
func (s *sMiddleware) MiddlewareHandlerResponse(r *ghttp.Request) {
	r.Middleware.Next()

	// There's custom buffer content, it then exits current handler.
	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		msg  string
		err  = r.GetError()
		res  = r.GetHandlerResponse()
		code = gerror.Code(err)
	)
	if err != nil {
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
		msg = err.Error()
	} else {
		if r.Response.Status > 0 && r.Response.Status != http.StatusOK {
			msg = http.StatusText(r.Response.Status)
			switch r.Response.Status {
			case http.StatusNotFound:
				code = gcode.CodeNotFound
			case http.StatusForbidden:
				code = gcode.CodeNotAuthorized
			default:
				code = gcode.CodeUnknown
			}
			// It creates error as it can be retrieved by other middlewares.
			err = gerror.NewCode(code, msg)
			r.SetError(err)
		} else {
			code = gcode.CodeOK
		}
	}
	if msg == "" {
		msg = "success"
	}
	r.Response.WriteJson(DefaultHandlerResponse{
		Code:      strconv.Itoa(code.Code()),
		Status:    r.Response.Status,
		Message:   msg,
		Result:    res,
		Timestamp: gtime.Now().UnixMilli(),
	})
}

// CORS allows Cross-origin resource sharing.
func (s *sMiddleware) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

// AccessKeyAuth 通过AK进行认证,后面会进行修改
func (s *sMiddleware) AccessKeyAuth(r *ghttp.Request) {
	// 从请求头获取 AK 值
	ak := r.GetHeader("Access-Key") // 或 r.Header.Get("AK")
	configAK, _ := g.Cfg().Get(context.Background(), "server.AccessKey")

	// 如果 AK 为空或不匹配
	if ak != configAK.String() {
		r.Response.WriteStatusExit(
			http.StatusUnauthorized,
			ghttp.DefaultHandlerResponse{
				Code:    http.StatusUnauthorized,
				Message: "Invalid access key",
			},
		)
		return
	}

	// 验证通过，继续后续处理
	r.Middleware.Next()
}

// JWTAuth 使用 JWT 进行认证
func (s *sMiddleware) JWTAuth(r *ghttp.Request) {
	// 从请求头获取 AK 值
	authHeader := r.GetHeader("Authorization") // 或 r.Header.Get("AK")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		//不用下面默认的回复方式
		//r.Response.WriteStatusExit(
		//	http.StatusUnauthorized,
		//	ghttp.DefaultHandlerResponse{
		//		Code:    http.StatusUnauthorized,
		//		Message: "Missing or invalid Authorization header",
		//	},
		//)
		err := gerror.NewCode(gcode.CodeNotAuthorized, "缺少或非法的 Authorization 头")
		r.SetError(err)
		r.Response.WriteStatus(http.StatusUnauthorized)
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	userID, err := service.Jwt().ParseToken(tokenString)
	if err != nil {
		err = gerror.NewCode(gcode.CodeNotAuthorized, err.Error())
		r.SetError(err)
		r.Response.WriteStatus(http.StatusUnauthorized)
		return
	}

	// 认证通过，设置用户信息上下文（可选）
	r.SetCtxVar("user_id", userID)
	r.Middleware.Next()
}
