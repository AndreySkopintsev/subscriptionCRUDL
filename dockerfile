FROM golang:1.24.4

WORKDIR /effective

COPY . .

RUN go mod download

RUN go build -o effective .

EXPOSE 8080

CMD ["./effective"]