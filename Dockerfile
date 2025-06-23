# 1. 设置模块代理和私有仓库白名单
FROM golang:1.24-alpine3.21 AS build-env
WORKDIR /app

#将当前目录一并拷进去
COPY ./ .
#展示路径
RUN pwd && ls -al  \
    &&  go env -w GOPROXY=https://goproxy.cn,direct  \
    &&  go build -o scaffold . \
    && ls -al


FROM alpine:3.20
WORKDIR /app

#设置时区相关
ENV TZ=Asia/Shanghai
RUN apk add --no-cache tzdata \
    && ln -sf /usr/share/zoneinfo/$TZ /etc/localtime \
    && echo $TZ > /etc/timezone

# 把本地编译好的二进制直接复制进来
# 在宿主机上执行了：$Env:GOOS   = 'linux'
           #$Env:GOARCH = 'amd64'
           #go build -o gateway .

# 复制静态编译的二进制
COPY --from=build-env /app/scaffold /app/scaffold
RUN pwd && ls -al && chmod +x /app/scaffold
EXPOSE 8001
CMD ["./scaffold"]
