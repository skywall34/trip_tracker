# Stage 1: Building the application
FROM golang:1.24.1-alpine AS builder
# Install c compiler
RUN apk add --no-cache gcc musl-dev
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# Binary was compiled with 'CGO_ENABLED=0', go-sqlite3 requires cgo to work. This is a stub
# This is since we're using github.com/mattn/go-sqlite3, using a pure GO SQLite driver might work better
RUN CGO_ENABLED=1 CGO_CFLAGS="-DSQLITE_ENABLE_MATH_FUNCTIONS" go build -tags sqlite_math_functions -o trip_tracker

# Stage 2: Run the application
FROM alpine:latest
RUN apk add --no-cache tzdata gcc musl-dev
WORKDIR /root/
COPY --from=builder /app/trip_tracker .
COPY --from=builder /app/static ./static
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/internal/database/database.db ./internal/database/database.db
EXPOSE 3000
CMD ["./trip_tracker"]