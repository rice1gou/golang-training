FROM golang:1.18 as builder

WORKDIR /go/src

COPY $GO_MOD go.sum ./
RUN go mod download

COPY cmd/identity/main.go ./

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
RUN go build \
    -o /go/bin/main \
    -ldflags '-s -w'

FROM scratch as runner

COPY --from=builder /go/bin/main /app/main

ENTRYPOINT ["/app/main"]
