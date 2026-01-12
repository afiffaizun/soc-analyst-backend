package handlers

import (
    "net/http"
    "github.com/afiffaizun/soc-analyst-backend/internal/database"
    "github.com/afiffaizun/soc-analyst-backend/internal/middlewares"
    "github.com/afiffaizun/soc-analyst-backend/internal/models"

    "github.com/gin-gonic/gin"
)

func GetTeams(c *gin.Context) {
    var teams []models.Team
    lang := middlewares.GetLanguage(c.Request)

    if err := database.DB.Find(&teams).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    } 

    var response []map[string]interface{}
    for _, t := range teams {
        role := t.RoleEn
        bio := t.BioEn
        if lang == middlewares.LangId {
            role = t.RoleId
            bio = t.BioId
        }

        response = append(response, map[string]interface{}{
            "id":       t.ID,
            "name":     t.Name,
            "role":     role,
            "bio":      bio,
            "image_url": t.ImageURL,
        })
    }

    c.JSON(http.StatusOK, gin.H{"data": response})
}