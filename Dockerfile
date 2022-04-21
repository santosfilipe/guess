FROM golang:latest

WORKDIR /

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o bin/guess pkg/main.go

EXPOSE 8888

CMD [ "bin/guess" ]