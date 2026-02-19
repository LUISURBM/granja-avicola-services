# Cambiamos a la versión 1.25 para que coincida con tu go.mod
FROM golang:1.25-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

# Copiar archivos de módulos
COPY go.mod go.sum ./

# Ahora sí funcionará porque la versión de Go coincide o es superior
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api/main.go

# ETAPA DE EJECUCIÓN
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .

EXPOSE 8080
CMD ["./main"]