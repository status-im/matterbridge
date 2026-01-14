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

FROM gcr.io/distroless/static-debian11

COPY --from=builder /bin/matterbridge /bin/matterbridge

ENTRYPOINT ["/bin/matterbridge"]
