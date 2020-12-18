#!/bin/bash
protoc -I portsgrpc/ \
	-I${GOPATH}/src \
	--go_out=plugins=grpc:portsgrpc \
	portsgrpc/portsgrpc.proto
