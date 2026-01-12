package main

import (
    "log"
    "os"
    "strconv"

    "github.com/afiffaizun/soc-analyst-backend/internal/database"
    "github.com/afiffaizun/soc-analyst-backend/internal/handlers"
    "github.com/gin-gonic/gin"
    _ "github.com/lib/pq"
)

func main() {
    // Load DB config from env (with sensible defaults)
    dbHost := os.Getenv("DB_HOST")
    if dbHost == "" {
        dbHost = "db"
    }
    dbUser := os.Getenv("DB_USER")
    if dbUser == "" {
        dbUser = "socuser"
    }
    dbPassword := os.Getenv("DB_PASSWORD")
    if dbPassword == "" {
        dbPassword = "socpassword"
    }
    dbName := os.Getenv("DB_NAME")
    if dbName == "" {
        dbName = "soc_db"
    }
    dbPort := 5432
    if p := os.Getenv("DB_PORT"); p != "" {
        if v, err := strconv.Atoi(p); err == nil {
            dbPort = v
        } else {
            log.Printf("invalid DB_PORT %q, using default %d", p, dbPort)
        }
    }

    database.Connect(dbHost, dbUser, dbPassword, dbName, dbPort)

    r := gin.Default()

    api := r.Group("/api/v1")
    {
        api.GET("/services", handlers.GetServices)
        api.GET("/teams", handlers.GetTeams)
        api.GET("/articles", handlers.GetArticles)
    }

    port := os.Getenv("PORT")
    if port == "" {
        port = "5433" // default API port sesuai permintaan
    }

    log.Printf("Server starting on port %s", port)
    if err := r.Run(":" + port); err != nil {
        log.Fatal(err)
    }
}