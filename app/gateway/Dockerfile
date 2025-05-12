FROM m.daocloud.io/docker.io/library/golang:1.24.2-alpine AS builder

WORKDIR /app
COPY . .
ENV GOPROXY=https://goproxy.cn,direct
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /gateway ./cmd/main.go

FROM m.daocloud.io/docker.io/library/alpine:latest

WORKDIR /
COPY --from=builder /gateway /gateway
COPY --from=builder /app/conf /configs

EXPOSE 4000
ENTRYPOINT ["/gateway"]