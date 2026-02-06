FROM golang:1.24-alpine as build

WORKDIR /app

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go

FROM alpine:latest

WORKDIR /app
COPY --from=build /app/main .
COPY .env .

EXPOSE 8080

CMD ["./main"]
