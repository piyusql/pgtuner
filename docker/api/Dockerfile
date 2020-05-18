FROM golang:1.14 AS builder
WORKDIR /go/src/github.com/piyusgupta/pgtuner/
COPY backend .
RUN  CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /data/pgtuner
COPY --from=builder /go/src/github.com/piyusgupta/pgtuner/app .
CMD ["./app"]
