package database

import (
    "fmt"
    "log"
    "time"
    
	"github.com/afiffaizun/soc-analyst-backend/internal/models"
    
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func Connect(host, user, password, dbname string, port int) {
    // Susun DSN (Data Source Name)
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", host, user, password, dbname, port)

    var err error
    var maxRetries = 10 // Mencoba sebanyak 10 kali
    
    log.Println("Attempting to connect to database...")

    for i := 0; i < maxRetries; i++ {
        DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
        if err == nil {
            // Jika koneksi berhasil, keluar dari loop
            log.Println("Database connected successfully!")
            break
        }

        // Log error retry
        log.Printf("Database not ready yet (attempt %d/%d). Retrying in 2 seconds... Error: %v", i+1, maxRetries, err)
        
        // Tunggu 2 detik sebelum mencoba lagi
        time.Sleep(2 * time.Second)
    }

    // Jika setelah 10 kali coba tetap gagal
    if err != nil {
        log.Fatal("Failed to connect to database after several retries:", err)
    }

    // Auto migrate models
    log.Println("Running database migrations...")
    err = DB.AutoMigrate(&models.Service{}, &models.Team{}, &models.Article{})
    if err != nil {
        log.Fatal("Failed to migrate database:", err)
    }
    
    log.Println("Database migration completed!")
}