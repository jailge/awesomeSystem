# 打包依赖阶段使用 golang 作为基础镜像
#FROM golang:1.17-alpine as builder
FROM golang:alpine as builder

WORKDIR /build
#RUN adduser -u 10001 -D app-runner

# 启用 go module
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

# CGO_ENABLED禁用 cgo 然后指定 OS，go build
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o bigSystem-public ./svc/public/cmd/
RUN go build -o awesomeSystem-user ./app/user/
#RUN go build -o awesomeSystem-weight ./app/weight/

# 运行阶段指定 scratch 作为基础镜像
#FROM alpine:3.10 as final
FROM scratch

#WORKDIR /app
#COPY --from=builder /build/bigSystem-public /app/
#COPY --from=builder /etc/passwd /etc/passwd
##COPY --from=builder /build/config.toml /app/
#COPY --from=builder /build/logs/ /app/logs/

COPY ./logs /logs
COPY ./settings.yaml /settings.yaml
COPY ./config/model.conf /config/model.conf

COPY --from=builder /build/awesomeSystem-user /
#COPY --from=builder /build/awesomeSystem-weight /

#ENV GIN_MODE=release \
#    PORT=8377

EXPOSE 8878
#EXPOSE 8877

#USER app-runner
#ENTRYPOINT ["/app/frApi", "-config=./config.toml"]
ENTRYPOINT ["/awesomeSystem-user"]
#ENTRYPOINT ["/awesomeSystem-weight"]