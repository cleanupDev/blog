# Stage 1: Build
FROM golang:1.23.4 AS build

WORKDIR /app

# Copy files first
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Install and setup templ
RUN go install github.com/a-h/templ/cmd/templ@latest
ENV PATH="$PATH:/go/bin"
RUN templ generate

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2: Run
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy the binary and necessary files
COPY --from=build /app/main .

EXPOSE 3000

CMD ["./main"]
