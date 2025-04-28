FROM golang:alpine AS builder

# 构建环境配置
ENV CGO_ENABLED=0
WORKDIR /build

# 优化 APK 镜像源
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

# 如果你用 go mod，就直接复制 go.mod、go.sum 优化构建缓存
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main

# -------------------

FROM alpine

WORKDIR /app

# 安装调试工具（只调试时使用，后续可删）
RUN apk add --no-cache tzdata bash curl mysql-client

COPY --from=builder /build/main /app/
COPY settings.yaml /app/
# COPY var /app/var/  # 确保 data_dir 存在
# COPY uploads /app/uploads/  # 如果你用到静态资源

CMD ["./main"]