FROM golang

COPY . /go/src/github.com/exitcodezero/picloud
WORKDIR /go/src/github.com/exitcodezero/picloud
RUN go get

EXPOSE 9000
