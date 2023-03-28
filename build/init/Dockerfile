FROM golang:1.18 as builder

WORKDIR /go/src

RUN echo $GO_MOD
ARG GO_MOD MOD_NAME
RUN ls
RUN echo $GO_MOD
COPY $GO_MOD ./
COPY backend/go.sum ./
RUN rename $MOD_NAME go.mod
RUN pwd
RUN ls
RUN go mod download

COPY backend/cmd/init/main.go ./

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
RUN go build \
    -o /go/bin/main \
    -ldflags '-s -w'

FROM scratch as runner

COPY --from=builder /go/bin/main /app/main

ENTRYPOINT ["/app/main"]
