package handlers

import (
    "net/http"
    
	"github.com/afiffaizun/soc-analyst-backend/internal/database"
	"github.com/afiffaizun/soc-analyst-backend/internal/middlewares"
	"github.com/afiffaizun/soc-analyst-backend/internal/models"
	"github.com/gin-gonic/gin"
)

func GetServices(c *gin.Context) {
    var services []models.Service
    lang := middlewares.GetLanguage(c.Request)

    if err := database.DB.Find(&services).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Transform response based on language
    var response []map[string]interface{}
    for _, s := range services {
        title := s.TitleEn
        desc := s.DescriptionEn
        if lang == middlewares.LangId {
            title = s.TitleId
            desc = s.DescriptionId
        }

        response = append(response, map[string]interface{}{
            "id":          s.ID,
            "title":       title,
            "description": desc,
            "icon_url":    s.IconURL,
        })
    }

    c.JSON(http.StatusOK, gin.H{"data": response})
}