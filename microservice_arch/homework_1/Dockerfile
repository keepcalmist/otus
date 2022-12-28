FROM golang:latest

WORKDIR /api
COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

ENV PORT=8080

RUN go build -o /otus-helthcheck

EXPOSE 8080

CMD ["/otus-helthcheck"]
