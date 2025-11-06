# ============================
# Stage 1: Build
# ============================
FROM registry.access.redhat.com/ubi9/go-toolset:1.18 AS builder
WORKDIR /app

COPY . .
RUN go mod tidy && CGO_ENABLED=0 GOOS=linux go build -o main .

# ============================
# Stage 2: Runtime
# ============================
FROM registry.access.redhat.com/ubi9/ubi-minimal:latest
WORKDIR /app

# تثبيت الشهادات وإعداد المجلدات المطلوبة لـ Kubernetes
RUN microdnf install -y ca-certificates && update-ca-trust && \
    mkdir -p /var/run/secrets/kubernetes.io/serviceaccount && \
    chmod 755 /var/run/secrets/kubernetes.io/serviceaccount && \
    microdnf clean all

# نسخ التطبيق من مرحلة البناء
COPY --from=builder /app/main .

# تشغيل الحاوية بيوزر غير root (مطلوب في OpenShift)
USER 1001

# الأمر الافتراضي لتشغيل التطبيق
CMD ["./main"]
