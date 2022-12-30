FROM golang AS build-env
ENV GOPROXY https://goproxy.cn
ADD . /go/src/app
WORKDIR /go/src/app
RUN go mod tidy
RUN GOOS=linux GOARCH=386 go build -v -o /go/src/app/go-hello-world

FROM alpine
COPY --from=build-env /go/src/app/go-hello-world /usr/local/bin/go-hello-world
EXPOSE 8080
CMD [ "go-hello-world" ]