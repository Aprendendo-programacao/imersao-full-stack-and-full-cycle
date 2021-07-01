ENV PATH="/go/bin:${path}"
ENV GO111MODULE=on
ENV CGO_ENABLED=1

RUN apt-get update && \
    apt-get install build-essential protobuf-compile -y && \
    go get google.golang.org/grpc/cmd/protoc-gen-go-grpc && \
    go get google.golang.org/protobuf/cmd/protoc-gen-go && \
    go get github.com/spf13/cobra/cobra

CMD ["tail", "-f", "/dev/null"]