FROM golang:1.23 as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-app

FROM alpine:latest
COPY --from=builder /go-app /go-app
EXPOSE 8080
ENTRYPOINT ["/go-app"]