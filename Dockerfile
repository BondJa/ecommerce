FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN ls cmd

RUN go build -o /app/bin/ecommerce /app/cmd/main.go

CMD [ "/app/bin/ecommerce" ]