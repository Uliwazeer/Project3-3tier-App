# ============================
# Stage 1: Build
# ============================
FROM registry.access.redhat.com/ubi9/go-toolset:1.18 AS builder
WORKDIR /app

# نستخدم المستخدم root طول مرحلة الـ build
USER 0

# نسخ الملفات وضبط الصلاحيات
COPY . .
RUN chmod -R 777 /app

# تنفيذ build
RUN go mod tidy && CGO_ENABLED=0 GOOS=linux go build -o main .

# ============================
# Stage 2: Runtime
# ============================
FROM registry.access.redhat.com/ubi9/ubi-minimal:latest
WORKDIR /app

RUN microdnf install -y ca-certificates && update-ca-trust && \
    mkdir -p /var/run/secrets/kubernetes.io/serviceaccount && \
    chmod 755 /var/run/secrets/kubernetes.io/serviceaccount && \
    microdnf clean all

COPY --from=builder /app/main .

# تشغيل الحاوية بيوزر غير root
USER 1001

CMD ["./main"]
