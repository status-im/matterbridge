FROM golang:1.23-bullseye AS builder

RUN apt update && apt install -y llvm clang

WORKDIR /build
COPY . .

ENV GO111MODULE=on
ENV CC=clang
ENV CXX=clang++

RUN go build -ldflags=-checklinkname=0 -o /bin/matterbridge

FROM debian:bullseye-slim
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates && rm -rf /var/lib/apt/lists/*
COPY --from=builder /bin/matterbridge /bin/matterbridge
ENTRYPOINT ["/bin/matterbridge"]