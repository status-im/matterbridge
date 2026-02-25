FROM golang:1.23-bullseye AS builder

RUN apt update
RUN apt install -y llvm clang

ADD go.mod go.sum ./
RUN go mod download

COPY . /go/src/github.com/42wim/matterbridge

WORKDIR /go/src/github.com/42wim/matterbridge

ENV GOPATH=/go
ENV CC=clang
ENV CXX=clang++

RUN go build -ldflags=-checklinkname=0 -o /bin/matterbridge

FROM debian:bullseye-slim
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates && rm -rf /var/lib/apt/lists/*
COPY --from=builder /bin/matterbridge /bin/matterbridge

ENTRYPOINT ["/bin/matterbridge"]