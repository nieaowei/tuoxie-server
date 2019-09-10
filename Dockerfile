FROM golang:latest
RUN mkdir tuoxie-user-handle-service &&\
	go env -w GO111MODULE=on &&\
	go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct 
ADD . tuoxie-user-handle-service
WORKDIR tuoxie-user-handle-service
RUN go build main.go
EXPOSE 8080
COPY libtaos.so.1 /usr/lib
ENTRYPOINT ["./main"]