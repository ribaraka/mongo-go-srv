ARG GO_VERSION=1.16.2
FROM golang:${GO_VERSION} AS builder

ARG ARCH=amd64
ARG GO111MODULE=on

EXPOSE 8080

WORKDIR $GOPATH/src/github.com/ribaraka/go-srv-example/

COPY go.mod go.sum ./

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=${ARCH} go build \
		-o ./app \
		-mod=vendor \
		-a ./cmd

FROM alpine:latest
COPY --from=builder /go/src/github.com/ribaraka/go-srv-example/app /app

ENTRYPOINT ["/app"]