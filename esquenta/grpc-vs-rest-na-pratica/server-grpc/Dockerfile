FROM golang:1.16.0-stretch

WORKDIR /go/src/

ENV PATH="/go/bin:${PATH}"
ENV CGO_ENABLED=0

RUN apt-get update && \
    apt-get install build-essential protobuf-compiler -y && \
    go get google.golang.org/grpc/cmd/protoc-gen-go-grpc && \
    go get google.golang.org/protobuf/cmd/protoc-gen-go && \
    go get github.com/spf13/cobra/cobra

CMD ["tail", "-f", "/dev/null"]