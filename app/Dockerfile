FROM golang:1.12.9-alpine3.10 as builder
COPY . .
RUN go build -o /app .

FROM alpine:3.10
CMD ["./app"]
COPY --from=builder /app .