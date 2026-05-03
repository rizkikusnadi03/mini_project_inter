# Backend Golang Mini Project

Proyek ini adalah aplikasi backend (REST API) yang dibangun menggunakan bahasa pemrograman Go (Golang) dan mengimplementasikan pola **Clean Architecture**.

## 🛠️ Teknologi & Library yang Digunakan

Berikut adalah daftar teknologi dan pustaka (library) utama yang digunakan dalam proyek ini:

1. **[Go (Golang)](https://go.dev/) v1.22**: Bahasa pemrograman utama yang digunakan untuk membangun sistem backend yang cepat dan efisien.
2. **[Go Fiber (v2)](https://docs.gofiber.io/)**: Framework web untuk Go yang terinspirasi dari Express.js. Sangat cepat, minim alokasi memori, dan digunakan untuk membuat routing serta menangani request/response HTTP.
3. **[GORM](https://gorm.io/)**: Library Object Relational Mapping (ORM) untuk Go. Digunakan untuk berinteraksi dengan database secara lebih mudah menggunakan struktur objek (struct) Go.
4. **MySQL Driver (`gorm.io/driver/mysql`)**: Driver koneksi agar GORM dapat terhubung dan berkomunikasi dengan database MySQL.
5. **JWT (`github.com/golang-jwt/jwt/v5`)**: Library untuk membuat, menandatangani, dan memvalidasi JSON Web Tokens (JWT). Digunakan untuk sistem autentikasi dan otorisasi pengguna.
6. **Godotenv (`github.com/joho/godotenv`)**: Library untuk memuat variabel lingkungan (environment variables) dari file `.env` ke dalam environment sistem (seperti konfigurasi database, port, dan secret key JWT).
7. **Bcrypt (`golang.org/x/crypto`)**: Library kriptografi yang digunakan secara khusus untuk melakukan *hashing* password pengguna agar tersimpan dengan aman di database.

## 🏗️ Struktur Arsitektur

Proyek ini menggunakan **Clean Architecture** untuk memisahkan *concern* pada kode. Anda akan menemukan direktori `internal` yang terbagi ke dalam beberapa lapisan (layer):
- **Port (`internal/core/port`)**: Mendefinisikan antarmuka (interface) dari *repository* dan *usecase*.
- **Usecase (`internal/core/usecase`)**: Berisi logika bisnis (business logic) dari aplikasi.
- **Handler/Controller (`internal/adapter/handler`)**: Menangani HTTP Request dari klien, memanggil *usecase*, dan mengembalikan HTTP Response.
- **Repository (`internal/adapter/repository`)**: Bertanggung jawab atas interaksi data (Query, Insert, Update, Delete) langsung ke database.
- **Router (`internal/adapter/router`)**: Tempat mendaftarkan semua rute/endpoint API.

## 🚀 Cara Menjalankan Proyek

### 1. Persyaratan Sistem (Prerequisites)
Pastikan Anda sudah menginstal beberapa software berikut di komputer Anda:
- **Go** (versi 1.22 atau lebih baru)
- **MySQL Server** (sedang berjalan)
- **Git** (opsional)

### 2. Konfigurasi Database & Environment
1. Buat database baru di MySQL server Anda (misal: `notion_db`).
2. Buat file bernama `.env` di root direktori proyek.
3. Salin konfigurasi di bawah ini ke dalam file `.env` yang baru saja dibuat, dan sesuaikan nilai `DB_USER`, `DB_PASSWORD`, `DB_NAME` dengan konfigurasi MySQL lokal Anda:

```env
DB_HOST=127.0.0.1
DB_PORT=3306
DB_USER=root
DB_PASSWORD=mypassword
DB_NAME=notion_db
JWT_SECRET=evermos_secret_key
PORT=8000
```

### 3. Mengunduh Dependencies
Buka terminal di root direktori proyek, lalu jalankan perintah berikut untuk mengunduh semua pustaka yang dibutuhkan (berdasarkan `go.mod`):
```bash
go mod tidy
```

### 4. Menjalankan Aplikasi
Jalankan aplikasi dengan perintah:
```bash
go run main.go
```

Jika berhasil, server akan berjalan dan dapat diakses pada port yang didefinisikan di dalam file `.env` (secara default `http://localhost:8000`).

### 5. Menguji API
Untuk menguji Endpoint API yang tersedia, Anda bisa mengimpor file `Internship.postman_collection.json` ke dalam aplikasi **Postman** Anda.
