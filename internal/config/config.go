package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

// Struct untuk seluruh akses konfigurasi aplikasi
type Config struct {
	AppName    string
	AppEnv     string
	AppPort    string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
	JWTExpire  string
}

// LoadConfig membaci env variables dan mengembalikan ke Config struct
// Fungsi yang dipanggil pertama kali saat aplikasi dijalankan
func LoadConfig() (*Config, error) {
	// Load file .env kedalam environtment variables
	// godotenv.Load() akan membaca file .env dan set sebagai env variables
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	// Membuat instance Config
	// os.Getenv() untuk membaca nilai environment variable
	config := &Config{
		AppName:    os.Getenv("APP_NAME"),
		AppEnv:     os.Getenv("APP_ENV"),
		AppPort:    os.Getenv("APP_PORT"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		JWTSecret:  os.Getenv("JWT_SECRET"),
		JWTExpire:  os.Getenv("JWT_EXPIRE"),
	}

	return config, nil
}

// GetDSN menghasilkan Data Source Name untuk koneksi MySQL
// DSN adalah string koneksi yang berisi info host, port, user, password, dan database
func (c *Config) GetDSN() string {
	// Format DSN untuk MySQL: user:password@tcp(host:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	// parseTime=True: mengkonversi DATE/DATETIME dari MySQL ke time.Time di Go
	// charset=utf8mb4: support untuk emoji dan karakter unicode lengkap
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.DBUser,
		c.DBPassword,
		c.DBHost,
		c.DBPort,
		c.DBName,
	)
}
