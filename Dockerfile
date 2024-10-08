FROM golang:1.22.0 AS builder

# Set destination for COPY
WORKDIR /app

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY cmd/done/main.go ./cmd/done/
COPY pkg/ ./pkg/

# Download Go modules
COPY go.mod .
COPY go.sum .
RUN go mod tidy

# Build
RUN go build -o /app/done ./cmd/done

# Run stage
FROM debian:trixie-slim

# Set dest for COPY
WORKDIR /app

# Install software
RUN apt-get update
RUN apt-get install -y curl sqlite3

# Copy the executables
COPY --from=builder /app/done /app/done

# Copy templates
COPY templates/ ./templates/

# Copy entrypoint script
# COPY entrypoint.sh ./entrypoint.sh
# RUN chmod +x ./entrypoint.s

# This is for documentation purposes only.
# To actually open the port, runtime parameters
# must be supplied to the docker command.
EXPOSE 3000

HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 CMD [ "curl", "-f", "http://localhost:3000/marco" ]

# Run
ENTRYPOINT ["/app/done"]
