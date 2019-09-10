FROM golang:latest
RUN mkdir tuoxie-user-handle-service &&\
	go env -w GO111MODULE=on &&\
	go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct 
ADD . tuoxie-user-handle-service
WORKDIR tuoxie-user-handle-service
RUN go build .
EXPOSE 8080
ENTRYPOINT ["./tuoxie-user-handle-service/main"]