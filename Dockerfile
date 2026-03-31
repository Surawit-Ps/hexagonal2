# multi-stage build for Go app
FROM golang:1.20-alpine AS builder

WORKDIR /app

# cache deps
COPY go.mod go.sum ./
RUN go mod download

# copy source
COPY . .

# build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /usr/local/bin/hgo .

# final image
FROM alpine:3.18
RUN apk add --no-cache ca-certificates

COPY --from=builder /usr/local/bin/hgo /usr/local/bin/hgo

WORKDIR /app

EXPOSE 3000

ENTRYPOINT ["/usr/local/bin/hgo"]
