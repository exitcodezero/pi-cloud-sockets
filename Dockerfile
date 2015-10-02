FROM golang

COPY ./app /go/src/app
WORKDIR /go/src/app
RUN go get

EXPOSE 9000
