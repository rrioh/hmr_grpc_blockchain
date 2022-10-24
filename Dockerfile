FROM amazonlinux:2

RUN yum install -y wget tar gzip

RUN wget https://go.dev/dl/go1.19.2.linux-arm64.tar.gz
RUN tar -C /usr/local -xzf go1.19.2.linux-arm64.tar.gz

ENV PATH $PATH:/usr/local/go/bin
ENV GOPATH /root/go

RUN mkdir -p /root/go/pkg
RUN mkdir -p /root/go/bin

WORKDIR /root/go/src/hmr_grpc_blockchain