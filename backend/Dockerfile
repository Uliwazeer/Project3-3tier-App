# ============================
# Stage 1: Build
# ============================
FROM quay.io/library/golang:1.18-alpine AS builder
WORKDIR /app

# تثبيت git علشان تحميل dependencies
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# ============================
# Stage 2: Runtime
# ============================
FROM quay.io/library/alpine:latest
WORKDIR /app

# تثبيت الشهادات وإعداد المجلدات المطلوبة لـ Kubernetes
RUN apk add --no-cache ca-certificates && update-ca-certificates \
    && mkdir -p /var/run/secrets/kubernetes.io/serviceaccount \
    && chmod 755 /var/run/secrets/kubernetes.io/serviceaccount

# نسخ التطبيق من مرحلة البناء
COPY --from=builder /app/main .
RUN chmod +x /app/main

# تشغيل الحاوية بيوزر غير root (مطلوب في OpenShift)
USER 1001

# الأمر الافتراضي لتشغيل التطبيق
CMD ["./main"]
