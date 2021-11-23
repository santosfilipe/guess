FROM golang:latest

WORKDIR /

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o /guess

EXPOSE 3000

CMD [ "/guess" ]