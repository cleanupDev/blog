# Stage 1: Build
FROM golang:1.23.2 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2: Run
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy the binary and necessary files
COPY --from=build /app/main .
COPY --from=build /app/templates ./templates
COPY --from=build /app/static ./static

EXPOSE 3000

CMD ["./main"]
