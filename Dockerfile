FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY go.mod go.sum* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o amartha-recon-service .

FROM alpine:3.19
LABEL maintainer="Amartha Recon POC"
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /app
COPY --from=builder /app/amartha-recon-service .
RUN mkdir -p storage
USER nobody:nobody
EXPOSE 8080
ENTRYPOINT ["./amartha-recon-service", "serveHttp"]
