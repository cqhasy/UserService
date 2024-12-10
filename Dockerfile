FROM golang:1.23 AS builder

COPY . /src
WORKDIR /src

RUN GOPROXY=https://goproxy.cn make build

FROM debian:stable-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src/bin /app

WORKDIR /app

EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf

CMD ["./UserService", "-conf", "/data/conf"]

# 如果不用docker-compose统一部署，而是数据库和服务分开，这样子部署
# 这里用的绝对路径，也可以相对路径，但是要注意相对路径是相对于docker build的路径
# docker build -t userservice .
# docker run -d -p 8000:8000 -p 9000:9000 -v /c/Users/24933/Documents/Code/Go/src/UserService/configs:/data/conf  --name userservice userservice