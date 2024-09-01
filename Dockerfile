# 第一阶段：使用 Go 官方镜像编译 Go 程序
FROM golang:1.23.0 as builder
WORKDIR /app

# 复制应用源代码到工作目录
COPY . /app

# 列出当前目录的内容，用于调试
RUN ls -al ./ && echo "Listing contents of the current directory"

# 禁用 CGO 并编译应用程序，确保静态链接
ENV CGO_ENABLED=0
RUN go build -ldflags '-extldflags "-static"' -o cache_server ./cache/node/main.go
RUN go build -ldflags '-extldflags "-static"' -o cache_client ./cache/read/client.go
RUN go build -ldflags '-extldflags "-static"' -o center_server ./center/node/main.go
RUN go build -ldflags '-extldflags "-static"' -o center_client ./center/write/client.go

# 第二阶段：使用轻量级的基础镜像
FROM alpine:latest
WORKDIR /app

# 从构建者阶段复制编译好的可执行文件
COPY --from=builder /app/cache_server /app/
COPY --from=builder /app/cache_client /app/
COPY --from=builder /app/center_server /app/
COPY --from=builder /app/center_client /app/

