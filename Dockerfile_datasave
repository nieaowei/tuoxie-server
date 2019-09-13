FROM centos:latest
RUN yum install -y kde-l10n-Chinese &&\
	yum reinstall -y glibc-common &&\
	localedef -c -f UTF-8 -i en_US en_US.UFT-8 &&\
	locale -a
ENV  LC_ALL=en_US.UTF-8
ENV LD_LIBRARY_PATH="$LD_LIBRARY_PATH:/usr/lib"

RUN yum install -y wget &&\
	yum install gcc -y &&\
	wget https://dl.google.com/go/go1.13.linux-amd64.tar.gz &&\
	tar -C /usr/local -xzf go1.13.linux-amd64.tar.gz

RUN mkdir tuoxie-user-handle-service &&\
	/usr/local/go/bin/go env -w GO111MODULE=on &&\
	/usr/local/go/bin/go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct &&\
	/usr/local/go/bin/go env -w GOROOT=/usr/local/go

ADD . tuoxie-user-handle-service
WORKDIR tuoxie-user-handle-service

ADD taos.tar.gz .
RUN cd taos-1.6.1.7-linux-2019-08-22-17-35 && ./install_client.sh &&\
	echo "charset UTF-8" >> /etc/taos/taos.cfg

RUN /usr/local/go/bin/go build main.go
EXPOSE 8080

ENTRYPOINT ["./main","-type=datasave"]