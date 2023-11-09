FROM golang:1.19-alpine AS builder

WORKDIR /build

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/main.go


FROM ubuntu:16.04

RUN mkdir -p kubeedge

COPY --from=builder /build/main kubeedge/

WORKDIR kubeedge
