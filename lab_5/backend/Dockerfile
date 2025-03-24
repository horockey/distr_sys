FROM golang:1.24.1-alpine
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o backend .
CMD ["./backend"]