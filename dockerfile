FROM golang:1.17-alpine

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go build -o server .

RUN mkdir "store" && mkdir "store/document" && mkdir "store/audio"

CMD [ "/app/server" ]
