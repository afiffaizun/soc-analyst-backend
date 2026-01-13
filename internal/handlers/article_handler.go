package handlers

import (
    "encoding/json"
    "net/http"
    "os"
    "time"
    "github.com/afiffaizun/soc-analyst-backend/internal/database"
	"github.com/afiffaizun/soc-analyst-backend/internal/middlewares"
	"github.com/afiffaizun/soc-analyst-backend/internal/models"
	"github.com/gin-gonic/gin"
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

func GetArticles(c *gin.Context) {
    // #region agent log
    debugLog("article_handler.go:11", "GetArticles function entry", map[string]interface{}{}, "B")
    // #endregion
    var articles []models.Article
    lang := middlewares.GetLanguage(c.Request)

    // Only get published articles
    // #region agent log
    debugLog("article_handler.go:16", "before database query", map[string]interface{}{
        "query": "published_at = ?", "value": true,
    }, "B")
    // #endregion
    if err := database.DB.Where("published_at = ?", true).Find(&articles).Error; err != nil {
        // #region agent log
        debugLog("article_handler.go:17", "database query error", map[string]interface{}{
            "error": err.Error(),
        }, "B")
        // #endregion
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    // #region agent log
    debugLog("article_handler.go:19", "database query success", map[string]interface{}{
        "articleCount": len(articles),
    }, "B")
    // #endregion

    var response []map[string]interface{}
    for _, a := range articles {
        title := a.TitleEn
        content := a.ContentEn
        if lang == middlewares.LangId {
            title = a.TitleId
            content = a.ContentId
        }

        response = append(response, map[string]interface{}{
            "id":        a.ID,
            "title":     title,
            "content":   content,
            "image_url": a.ImageURL,
            "created_at": a.CreatedAt,
        })
    }

    c.JSON(http.StatusOK, gin.H{"data": response})
}