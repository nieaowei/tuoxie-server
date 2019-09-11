FROM golang:latest
RUN mkdir tuoxie-user-handle-service &&\
	go env -w GO111MODULE=on &&\
	go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct 
ADD . tuoxie-user-handle-service
WORKDIR tuoxie-user-handle-service
ADD taos.tar.gz .
RUN cd taos-1.6.1.7-linux-2019-08-22-17-35 && ./install_client.sh
RUN go build main.go
COPY taos.cfg /etc/taos/taos.cfg
EXPOSE 8080
ENTRYPOINT ["./main"]