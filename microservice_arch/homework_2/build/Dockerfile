FROM golang:latest

WORKDIR /api
COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

ENV PORT=8000

RUN go build -o /otus-helthcheck

EXPOSE 8000

CMD ["/otus-helthcheck"]
