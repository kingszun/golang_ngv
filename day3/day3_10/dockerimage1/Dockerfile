FROM golang:1.20

RUN mkdir /app

ADD . /app

WORKDIR /app

COPY go.mod ./

RUN go mod download 

RUN go build -o main .

CMD ["/app/main"]