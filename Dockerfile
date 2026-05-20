FROM golang:1.22-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY src/ .

RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-w -s" \
    -o /book-library \
    .


FROM scratch

COPY --from=builder /book-library /book-library

EXPOSE 8080

ENTRYPOINT ["/book-library"]