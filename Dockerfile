# Gunakan image Golang resmi sebagai builder
FROM golang:1.21.4 AS builder

# Set direktori kerja
WORKDIR /app

# Salin go.mod dan go.sum
COPY go.mod go.sum ./

# Unduh semua dependensi
RUN go mod download

# Salin seluruh kode sumber
COPY src/ ./src/

# Set variabel lingkungan untuk cross-compilation
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

# Bangun aplikasi
RUN go build -o myapp ./src/apisimulation/main.go

# Mulai tahap baru dari awal menggunakan Alpine
FROM alpine:latest

# Salin binary yang sudah dibangun dari tahap builder
COPY --from=builder /app/myapp .

# Salin file konfigurasi
COPY --from=builder /app/src/config/config.json ./src/config/config.json

# Berikan izin eksekusi
RUN chmod +x myapp

# Ekspose port 8080
EXPOSE 8080

# Perintah untuk menjalankan aplikasi
CMD ["./myapp"]