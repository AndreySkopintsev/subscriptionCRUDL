FROM golang:1.25rc2-alpine3.22

WORKDIR /effective

COPY . .

RUN go mod download

RUN go build -o effective .

EXPOSE 8080

ENTRYPOINT ["./effective"]