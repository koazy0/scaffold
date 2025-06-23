package utils

import (
	"context"
	rand "crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/golang-jwt/jwt/v5"
	_ "scaffold/internal/logic/logs"
	"scaffold/internal/service"
	"time"
)

type (
	sJwt struct{}
)

func init() {
	var ctx = gctx.New()
	var err error
	secret, err := g.Cfg().Get(ctx, "utils.jwt_secret")
	if err != nil {
		service.Logs().Fatal(err.Error())
	}
	jwtSecret = secret.Bytes()
	service.Logs().Info("Init jwts success!")
	service.RegisterJwt(Jwt())
}

var insJwt = sJwt{}

func Jwt() *sJwt {
	return &insJwt
}

var jwtSecret = []byte("")

// CustomClaims 可根据需要扩展字段
type CustomClaims struct {
	UserID string `json:"user_id"`
	//Password string `json:"password"` #只存ID就行了，其余都通过ID查出来
	jwt.RegisteredClaims
}

// GenerateToken 生成一个带有 userID的 JWT，默认有效期 12 小时
func (s *sJwt) GenerateToken(ctx context.Context, userID string) (string, error) {
	// 构造自定义声明
	expiresTimeYaml, err := g.Cfg().Get(ctx, "utils.jwt_expire")
	expiresTimeHour := expiresTimeYaml.Int()
	if err != nil || expiresTimeHour == 0 {
		expiresTimeHour = 12
	}

	claims := CustomClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expiresTimeHour) * time.Hour)), // 24 小时后过期
			//ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Second)), // 测试用
			IssuedAt: jwt.NewNumericDate(time.Now()),
			//Issuer:    "your-app-name",
			//Subject: userID,
		},
	}

	// 创建 token 对象，指定签名方法和声明
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 用密钥字符串进行签名，得到完整的 JWT 字符串
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s *sJwt) ParseToken(tokenString string) (UserID string, err error) {

	// 告诉jwt库用自己的CustomClaims解析
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 校验签名方法是否为 HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("签名方法不为HMAC")
		}
		//告诉 JWT 库：用这个 jwtSecret 去验证签名
		return jwtSecret, nil
	})
	if err != nil {
		return "", err
	}
	// 断言为自己的类型，通过token.Valid校验token合法性，如过期时间，签发组织等；
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims.UserID, nil
	}
	return "", errors.New("无效的token")
}

// HashPassword 用 SHA-256(salt + password) 方式加密
func (s *sJwt) HashPassword(password, salt string) (passwordEncrypt string) {
	h := sha256.Sum256([]byte(salt + password))
	return hex.EncodeToString(h[:])
}

// GenerateSalt 生成随机盐值，长度固定为16个字节
func (s *sJwt) GenerateSalt() (string, error) {

	seed := make([]byte, 32)
	_, err := rand.Read(seed)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(seed), nil
}
