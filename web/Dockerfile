FROM golang:1.12.9-alpine3.10 as builder
COPY . .
RUN go build -o /web .

FROM alpine:3.10
CMD ["./web"]
COPY --from=builder /web .