package main

import (
    "log"
    "os"
    
	"github.com/afiffaizun/soc-analyst-backend/internal/database"
	"github.com/afiffaizun/soc-analyst-backend/internal/handlers"
	
    "github.com/gin-gonic/gin"
    _ "github.com/lib/pq"
)

func main() {
    // Initialize Database
    // Note: Ideally load these from env variables
    database.ConnectDB("localhost", "socuser", "socpassword", "soc_db", 5432)

    r := gin.Default()

    // API Routes
    api := r.Group("/api/v1")
    {
        api.GET("/services", handlers.GetServices)
        api.GET("/teams", handlers.GetTeams)
        api.GET("/articles", handlers.GetArticles)
    }

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("Server starting on port %s", port)
    if err := r.Run(":" + port); err != nil {
        log.Fatal(err)
    }
}