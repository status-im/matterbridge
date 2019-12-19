FROM alpine:3.10.0 AS builder

RUN apk update && apk add go git gcc musl-dev linux-headers

COPY . /go/src/github.com/42wim/matterbridge

WORKDIR /go/src/github.com/42wim/matterbridge

ENV GOPATH /go
ENV CGOENABLE 1
ENV GO111MODULE off

RUN go get
RUN go build -x -ldflags "-X main.githash=$(git log --pretty=format:'%h' -n 1)" -o /bin/matterbridge

FROM alpine:latest

RUN apk update && apk add ca-certificates

COPY --from=builder /bin/matterbridge /bin/matterbridge

ENTRYPOINT ["/bin/matterbridge"]
