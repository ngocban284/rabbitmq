FROM rabbitmq:3.9.7-management-alpine

# Install apk manually
RUN apk add --no-cache apk-tools

# Install additional dependencies
RUN apk add --no-cache \
    curl \
    wget \
    bash

# Set environment variables
ENV RABBITMQ_DEFAULT_USER=admin \
    RABBITMQ_DEFAULT_PASS=password \
    RABBITMQ_DEFAULT_VHOST=/ \
    RABBITMQ_LOGS=- \
    RABBITMQ_SASL_LOGS=-

# Expose ports
EXPOSE 5672 15672

# Start RabbitMQ server
CMD ["rabbitmq-server"]
