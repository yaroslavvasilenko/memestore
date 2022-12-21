# syntax=docker/dockerfile:1

FROM golang:1.19.4 as builder

COPY . /go/src/github.com/yaroslavvasilenko/memestore/
WORKDIR /go/src/github.com/yaroslavvasilenko/memestore/
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o memestore ./

FROM alpine:3.17.0 as production

RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/yaroslavvasilenko/memestore/memestore ./
COPY --from=builder /go/src/github.com/yaroslavvasilenko/memestore/.env ./
CMD ["./memestore"]