
FROM registry.access.redhat.com/ubi9/nginx-120:latest

# Copy config and SSL generation script
COPY nginx.conf /etc/nginx/nginx.conf
COPY generate-ssl.sh /usr/local/bin/generate-ssl.sh

# Generate SSL certs in /etc/nginx/ssl
RUN chmod +x /usr/local/bin/generate-ssl.sh && \
    /usr/local/bin/generate-ssl.sh

# ✅ Run as non-root user (required by OpenShift)
USER 1001

# Expose ports
EXPOSE 80 443

CMD ["nginx", "-g", "daemon off;"]
