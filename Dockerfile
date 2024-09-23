FROM golang:1.23.1-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

COPY pkg/config/envs/dev.env /app/cmd/pkg/config/envs/dev.env

WORKDIR /app/cmd

RUN go build -o /app/main .

EXPOSE 50054

CMD ["/app/main"]
