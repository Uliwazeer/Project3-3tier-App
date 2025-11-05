#!/bin/sh

SSL_DIR=/etc/nginx/ssl

mkdir -p $SSL_DIR

openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
  -keyout $SSL_DIR/selfsigned.key \
  -out $SSL_DIR/selfsigned.crt \
  -subj "/C=EG/ST=Cairo/L=Cairo/O=MyCompany/OU=IT/CN=localhost"

echo "✅ SSL certificate generated in $SSL_DIR successfully!"
