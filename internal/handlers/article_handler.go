package handlers

import (
    "net/http"
    "github.com/afiffaizun/soc-analyst-backend/internal/database"
	"github.com/afiffaizun/soc-analyst-backend/internal/middlewares"
	"github.com/afiffaizun/soc-analyst-backend/internal/models"
	"github.com/gin-gonic/gin"
)

func GetArticles(c *gin.Context) {
    var articles []models.Article
    lang := middlewares.GetLanguage(c.Request)

    // Only get published articles
    if err := database.DB.Where("published_at = ?", true).Find(&articles).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

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