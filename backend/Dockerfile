# Stage 1: build
FROM golang:1.18-alpine AS builder
WORKDIR /app
RUN apk add --no-cache git
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2: runtime
FROM alpine:latest
WORKDIR /app

# ضروري: تثبيت ca-certificates وخلق المجلدات التي يحتاجها K8s
RUN apk add --no-cache ca-certificates && update-ca-certificates \
    && mkdir -p /var/run/secrets/kubernetes.io/serviceaccount \
    && chmod 755 /var/run/secrets/kubernetes.io/serviceaccount

COPY --from=builder /app/main .
RUN chmod +x /app/main

# ✅ إضافة السطر المطلوب لتشغيل الحاوية بيوزر غير root
USER 1001

CMD ["./main"]
