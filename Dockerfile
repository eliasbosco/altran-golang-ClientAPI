FROM golang:1.15 as build-env

WORKDIR /go/src/app
ADD . /go/src/app

RUN GO111MODULE=off GOPATH=/go/bin/app go get -d -v ./...
RUN GO111MODULE=off GOPATH=/go/bin/app go build -ldflags "-s -w" -o /go/bin/app

FROM gcr.io/distroless/base-debian10
COPY --from=build-env /go/bin/app /
COPY ./ports.json /tmp
CMD ["/app"]
