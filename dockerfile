# syntax=docker/dockerfile:1

FROM golang:latest as builder

COPY . /go/src/github.com/yaroslavvasilenko/memestore/
WORKDIR /go/src/github.com/yaroslavvasilenko/memestore/
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o app ./

FROM alpine:latest as production

RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/yaroslavvasilenko/memestore/app ./
COPY --from=builder /go/src/github.com/yaroslavvasilenko/memestore/.env ./
CMD ["./app"]