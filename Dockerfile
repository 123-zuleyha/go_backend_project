# Go resmi imajını temel alıyoruz
FROM golang:1.22 AS builder

# Çalışma dizinini ayarlıyoruz
WORKDIR /app

# Go mod dosyalarını kopyalıyoruz
COPY ../go.mod ../go.sum ./

# Bağımlılıkları indiriyoruz
RUN go mod download

# Proje dosyalarını kopyalıyoruz
COPY ../ ./

# Go kodunu derliyoruz
RUN go build -o main .

# İkinci bir aşama başlatıyoruz, daha küçük bir imaj kullanacağız
FROM alpine:latest

# Çalışma dizinini ayarlıyoruz
WORKDIR /app

# Build aşamasından derlenmiş uygulamayı kopyalıyoruz
COPY --from=builder /app/main .

# Uygulamayı çalıştırıyoruz
CMD ["./main"]

# API'nin dinleyeceği portu belirtiriz
EXPOSE 8081
