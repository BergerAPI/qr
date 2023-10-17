FROM golang:1.21.3-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o qr

EXPOSE 8080

CMD [ "qr" ]