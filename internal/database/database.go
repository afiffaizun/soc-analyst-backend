package database

import (
    "encoding/json"
    "fmt"
    "log"
    "os"
    "time"
    
	"github.com/afiffaizun/soc-analyst-backend/internal/models"
    
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

// #region agent log
func debugLog(location, message string, data map[string]interface{}, hypothesisId string) {
    logEntry := map[string]interface{}{
        "sessionId":    "debug-session",
        "runId":        "run1",
        "hypothesisId": hypothesisId,
        "location":     location,
        "message":      message,
        "data":         data,
        "timestamp":    time.Now().UnixMilli(),
    }
    if jsonData, err := json.Marshal(logEntry); err == nil {
        if f, err := os.OpenFile("/home/smart/Belajar/project/soc-analyst-backend/.cursor/debug.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); err == nil {
            f.WriteString(string(jsonData) + "\n")
            f.Close()
        }
    }
}
// #endregion

var DB *gorm.DB

func Connect(host, user, password, dbname string, port int) {
    // #region agent log
    debugLog("database.go:16", "Connect function entry", map[string]interface{}{
        "host": host, "user": user, "dbname": dbname, "port": port,
    }, "C")
    // #endregion
    // Susun DSN (Data Source Name)
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", host, user, password, dbname, port)
    // #region agent log
    debugLog("database.go:20", "DSN constructed", map[string]interface{}{"dsn": dsn}, "C")
    // #endregion

    var err error
    var maxRetries = 10 // Mencoba sebanyak 10 kali
    
    log.Println("Attempting to connect to database...")

    for i := 0; i < maxRetries; i++ {
        // #region agent log
        debugLog("database.go:28", "before gorm.Open", map[string]interface{}{"attempt": i + 1, "maxRetries": maxRetries}, "C")
        // #endregion
        DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
        if err == nil {
            // Jika koneksi berhasil, keluar dari loop
            log.Println("Database connected successfully!")
            // #region agent log
            debugLog("database.go:32", "database connection success", map[string]interface{}{"attempt": i + 1}, "C")
            // #endregion
            break
        }

        // Log error retry
        log.Printf("Database not ready yet (attempt %d/%d). Retrying in 2 seconds... Error: %v", i+1, maxRetries, err)
        // #region agent log
        debugLog("database.go:38", "database connection retry", map[string]interface{}{
            "attempt": i + 1, "error": err.Error(),
        }, "C")
        // #endregion
        
        // Tunggu 2 detik sebelum mencoba lagi
        time.Sleep(2 * time.Second)
    }

    // Jika setelah 10 kali coba tetap gagal
    if err != nil {
        log.Fatal("Failed to connect to database after several retries:", err)
    }

    // Auto migrate models
    log.Println("Running database migrations...")
    // #region agent log
    debugLog("database.go:47", "before AutoMigrate", map[string]interface{}{}, "C")
    // #endregion
    err = DB.AutoMigrate(&models.Service{}, &models.Team{}, &models.Article{})
    if err != nil {
        // #region agent log
        debugLog("database.go:49", "AutoMigrate error", map[string]interface{}{"error": err.Error()}, "C")
        // #endregion
        log.Fatal("Failed to migrate database:", err)
    }
    // #region agent log
    debugLog("database.go:52", "AutoMigrate success", map[string]interface{}{}, "C")
    // #endregion
    
    log.Println("Database migration completed!")
}