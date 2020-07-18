FROM golang:1.14-alpine as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .

FROM alpine:latest
COPY --from=builder /app/app /usr/bin/app
ENTRYPOINT ["sh", "-c", "app"]