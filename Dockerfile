ARG BINARY_NAME="server"
ARG PORT="50051"
# Builder image
FROM golang:1.23.4-bookworm as builder
# Update and install make
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    make && \
    rm -rf /var/lib/apt/lists/*
# Set working directory
WORKDIR /app
# Copy source files
COPY . .
# Build the binary
RUN make build

# Release image
FROM debian:bookworm-slim as release

# Set working directory
WORKDIR /app
# Update and install ca-certificates
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*
# Copy Environment file
COPY --from=builder /app/.env .
# Export environment variables from .env file
RUN set -o allexport; source .env; set +o allexport
RUN export $(grep -v '^#' .env | xargs -d '\n')
# Copy the binary from the builder image
COPY --from=builder /app/bin/ .
# Run the binary
ARG BINARY_NAME
ARG PORT
ENV BINARY_NAME=${BINARY_NAME}
ENV PORT=${PORT}
# Set the binary as executable
RUN chmod +x /app/$BINARY_NAME
# Set container entrypoint to the environment variable
ENTRYPOINT /app/$BINARY_NAME
