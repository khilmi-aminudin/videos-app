FROM golang:1.18.2-alpine AS builder

WORKDIR /app

RUN apk update && apk add git

COPY . .

RUN go mod tidy

RUN go build -o videos-app

FROM alpine:3.13

WORKDIR /app

COPY --from=builder /app/videos-app /app/videos-app

EXPOSE 8080

CMD ["./videos-app", "--port=8080", "--host=0.0.0.0"]



