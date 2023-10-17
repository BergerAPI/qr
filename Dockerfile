FROM golang:1.21.3

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY * ./

RUN go get qr
RUN go build -o qr

EXPOSE 8080

CMD [ "./qr" ]