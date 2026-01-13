package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
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

// findAvailablePort tries to find an available port, starting with preferredPort
// If preferredPort is not available, it tries ports 8080-8090 as fallback
func findAvailablePort(preferredPort string) (string, error) {
	// #region agent log
	debugLog("main.go:findAvailablePort", "checking preferred port", map[string]interface{}{
		"preferredPort": preferredPort,
	}, "A")
	// #endregion

	// First, try the preferred port
	if isPortAvailable(preferredPort) {
		// #region agent log
		debugLog("main.go:findAvailablePort", "preferred port available", map[string]interface{}{
			"port": preferredPort,
		}, "A")
		// #endregion
		return preferredPort, nil
	}

	// #region agent log
	debugLog("main.go:findAvailablePort", "preferred port not available, trying fallback", map[string]interface{}{
		"preferredPort": preferredPort,
	}, "A")
	// #endregion

	// If preferred port is not available, try fallback ports (8080-8090)
	preferredPortInt, err := strconv.Atoi(preferredPort)
	if err != nil {
		preferredPortInt = 8080
	}

	// Try ports starting from 8080, or preferredPort+1 if preferredPort is numeric
	startPort := 8080
	if preferredPortInt >= 8080 && preferredPortInt < 8090 {
		startPort = preferredPortInt + 1
	}

	for port := startPort; port <= 8090; port++ {
		portStr := strconv.Itoa(port)
		// #region agent log
		debugLog("main.go:findAvailablePort", "checking fallback port", map[string]interface{}{
			"port": portStr, "attempt": port - startPort + 1,
		}, "A")
		// #endregion
		if isPortAvailable(portStr) {
			// #region agent log
			debugLog("main.go:findAvailablePort", "fallback port available", map[string]interface{}{
				"port": portStr, "preferredPort": preferredPort,
			}, "A")
			// #endregion
			log.Printf("Port %s is in use, using fallback port %s", preferredPort, portStr)
			return portStr, nil
		}
	}

	return "", fmt.Errorf("no available port found (tried %s and 8080-8090)", preferredPort)
}

// isPortAvailable checks if a port is available by attempting to listen on it
func isPortAvailable(port string) bool {
	// #region agent log
	debugLog("main.go:isPortAvailable", "checking port availability", map[string]interface{}{
		"port": port,
	}, "A")
	// #endregion

	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		// #region agent log
		debugLog("main.go:isPortAvailable", "port not available", map[string]interface{}{
			"port": port, "error": err.Error(),
		}, "A")
		// #endregion
		return false
	}
	ln.Close()
	// #region agent log
	debugLog("main.go:isPortAvailable", "port available", map[string]interface{}{
		"port": port,
	}, "A")
	// #endregion
	return true
}

func main() {
	// #region agent log
	debugLog("main.go:14", "main function entry", map[string]interface{}{}, "D")
	// #endregion
	// Load DB config from env (with sensible defaults)
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost" // Changed from "db" to "localhost" for running outside Docker
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

	// Get port from env or use default
	preferredPort := os.Getenv("PORT")
	if preferredPort == "" {
		preferredPort = "5433" // default API port sesuai permintaan
	}

	// #region agent log
	debugLog("main.go:87", "port configuration", map[string]interface{}{
		"preferredPort": preferredPort, "portEnv": os.Getenv("PORT"), "isDefault": os.Getenv("PORT") == "",
	}, "A")
	// #endregion

	// Try to find an available port
	port, err := findAvailablePort(preferredPort)
	if err != nil {
		// #region agent log
		debugLog("main.go:95", "failed to find available port", map[string]interface{}{
			"error": err.Error(), "preferredPort": preferredPort,
		}, "E")
		// #endregion
		log.Fatal("Failed to find available port:", err)
	}

	// #region agent log
	debugLog("main.go:100", "port selected", map[string]interface{}{
		"port": port, "preferredPort": preferredPort, "isFallback": port != preferredPort,
	}, "B")
	// #endregion

	log.Printf("Server starting on port %s", port)
	// #region agent log
	debugLog("main.go:105", "before server start", map[string]interface{}{
		"port": port, "address": ":" + port,
	}, "C")
	// #endregion

	// #region agent log
	debugLog("main.go:110", "attempting to bind port", map[string]interface{}{
		"port": port, "address": ":" + port,
	}, "D")
	// #endregion

	if err := r.Run(":" + port); err != nil {
		// #region agent log
		debugLog("main.go:115", "server start error", map[string]interface{}{
			"error": err.Error(), "port": port, "errorType": fmt.Sprintf("%T", err),
		}, "F")
		// #endregion
		log.Fatal(err)
	}
}
