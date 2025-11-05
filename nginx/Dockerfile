FROM registry.access.redhat.com/ubi9/nginx-120:latest

# Temporarily switch to root to copy and chmod
USER 0

# Copy config and SSL generation script
COPY nginx.conf /etc/nginx/nginx.conf
COPY generate-ssl.sh /usr/local/bin/generate-ssl.sh

# Generate SSL certs in /etc/nginx/ssl
RUN chmod +x /usr/local/bin/generate-ssl.sh && \
    mkdir -p /etc/nginx/ssl && \
    /usr/local/bin/generate-ssl.sh

# ✅ Switch back to non-root (required by OpenShift)
USER 1001

EXPOSE 80 443

CMD ["nginx", "-g", "daemon off;"]
