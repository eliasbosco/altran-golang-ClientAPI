FROM golang:1.15 as build-env
RUN apt-get update
RUN apt-get install python-dev libsasl2-dev -y

WORKDIR /go/src/app
ADD . /go/src/app

RUN GO111MODULE=off GOPATH=/go/bin/app go get -d -v ./...
RUN GO111MODULE=off GOPATH=/go/bin/app go build -o /go/bin/app

FROM gcr.io/distroless/base
COPY --from=build-env /go/bin/app /
COPY ./ports.json /tmp
CMD ["/app"]
