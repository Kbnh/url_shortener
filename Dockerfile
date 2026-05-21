FROM golang:1.25-alpine AS builder

RUN apk add --no-cache gcc musl-dev

WORKDIR /build

# dep cache
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -ldflags="-s -w" -o app ./cmd/url-shortener

FROM alpine:latest

WORKDIR /app

# cert / future
RUN apk --no-cache add ca-certificates

# bin copy
COPY --from=builder /build/app .

EXPOSE 8080

CMD ["./app"]
