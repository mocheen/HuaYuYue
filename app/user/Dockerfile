FROM m.daocloud.io/docker.io/library/golang:1.24.2-alpine AS builder

WORKDIR /app
COPY . .
ENV GOPROXY=https://goproxy.cn,direct
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /user ./cmd/main.go

FROM m.daocloud.io/docker.io/library/alpine:latest

WORKDIR /
COPY --from=builder /user /user
COPY --from=builder /app/conf /configs

EXPOSE 10002
ENTRYPOINT ["/user"]