FROM golang:1.23-alpine AS builder

# Set the working directory
WORKDIR /workspace

# Copy module definition files
COPY go.mod go.sum ./

# Install deps
RUN go mod tidy

# Copy the go source
COPY cmd/ ./cmd/
COPY internal/ ./internal/

# Build Go Program
RUN go build -o junkyard ./cmd/main.go

FROM alpine:latest
WORKDIR /root
COPY --from=builder /workspace/junkyard .

# Copy the internal package which includes static assets
COPY --from=builder /workspace/internal/ /root/internal/

EXPOSE 3333
CMD ["./junkyard"]