package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/afiffaizun/soc-analyst-backend/internal/database"
	"github.com/afiffaizun/soc-analyst-backend/internal/handlers"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
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

func main() {
	// #region agent log
	debugLog("main.go:14", "main function entry", map[string]interface{}{}, "D")
	// #endregion
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

	// #region agent log
	debugLog("main.go:41", "before database.Connect", map[string]interface{}{
		"dbHost": dbHost, "dbUser": dbUser, "dbName": dbName, "dbPort": dbPort,
	}, "C")
	// #endregion
	database.Connect(dbHost, dbUser, dbPassword, dbName, dbPort)
	// #region agent log
	debugLog("main.go:41", "after database.Connect", map[string]interface{}{"success": true}, "C")
	// #endregion

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
	// #region agent log
	debugLog("main.go:57", "before server start", map[string]interface{}{"port": port}, "D")
	// #endregion
	if err := r.Run(":" + port); err != nil {
		// #region agent log
		debugLog("main.go:59", "server start error", map[string]interface{}{"error": err.Error()}, "D")
		// #endregion
		log.Fatal(err)
	}
}
