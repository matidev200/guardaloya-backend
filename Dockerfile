# Etapa 1: Compilación
FROM golang:1.23.3 AS builder

WORKDIR /app

COPY . .

WORKDIR /app/cmd/app

RUN go build -o /app/myapp .

FROM golang:1.20
WORKDIR /app

COPY --from=builder /app/myapp .

EXPOSE 8080

CMD ["/app/myapp"]