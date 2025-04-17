# 🐦 Manuk POS Backend

**Manuk POS Backend** adalah sistem backend untuk aplikasi Point of Sale (POS) yang dibangun menggunakan bahasa pemrograman **Go** (Golang). Backend ini menyediakan API untuk mengelola pengguna, produk, transaksi, dan fitur POS lainnya. Cocok digunakan sebagai sistem kasir untuk usaha kecil hingga menengah.

---

## 🚀 Tech Stack

| Package                     | Deskripsi                                      |
|----------------------------|-----------------------------------------------|
| [Gin](https://github.com/gin-gonic/gin) | Web framework ringan dan cepat                 |
| [JWT](https://github.com/golang-jwt/jwt) | Autentikasi berbasis token (JWT v5)           |
| [GORM](https://gorm.io/)   | ORM untuk database MySQL                       |
| [godotenv](https://github.com/joho/godotenv) | Load konfigurasi dari file `.env`            |
| [x/crypto](https://pkg.go.dev/golang.org/x/crypto) | Utility kriptografi dari Go                |

---

## 📁 Struktur Project


---

## ⚙️ Instalasi & Menjalankan

### 1. Clone repository

```bash
git clone https://github.com/ibudiana/manuk-pos-backend.git
cd manuk-pos-backend
```

### Buat file .env

```env
DB_USER=root
DB_PASS=password
DB_HOST=localhost
DB_PORT=3306
DB_NAME=manuk_pos

SERVER_PORT=:8080

JWT_SECRETE=manuk_pos_secret_key_123
JWT_EXPIRES_IN_HOUR=24
```

## Install dependency
```
go mod tidy
```

## Run Server
```
go run main.go
```

## Authentication

Authorization: Bearer <token>
