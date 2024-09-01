# 第一阶段：使用 GCC 镜像编译 C++ 程序
FROM golang:latest as builder
WORKDIR /app
COPY . /app
RUN go build -o cache_server ./cache/node/main.go
RUN go build -o cache_client ./cache/read/client.go
RUN go build -o center_server ./center/node/main.go
RUN go build -o center_client ./center/write/client.go

# 第二阶段：使用轻量级的基础镜像
FROM alpine:latest
WORKDIR /app
# 从构建者阶段复制编译好的可执行文件
COPY --from=builder /app/cache_server /app
COPY --from=builder /app/cache_client /app
COPY --from=builder /app/center_server /app
COPY --from=builder /app/center_client /app

