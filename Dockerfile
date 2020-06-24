FROM golang:1.14 as builder

WORKDIR /

ADD main.go .

RUN CGO_ENABLED=0 GOOS=linux go build -o iptables-logger main.go

FROM scratch

COPY --from=builder /iptables-logger /iptables-logger

ENTRYPOINT [ "/iptables-logger" ]