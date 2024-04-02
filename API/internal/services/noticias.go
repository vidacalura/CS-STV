package services

import (
	"github.com/gin-gonic/gin"
	"github.com/vidacalura/CS-STV/internal/models"
)

// Pega o feed de notícias com as notícias mais recentes
func GetFeedNoticias(c *gin.Context) {
    var feed models.FeedNoticias

    statusCode, err := feed.GetFeedNoticias()
    if err != nil {
        c.IndentedJSON(statusCode, gin.H{"error": err.Error()})
        return
    }

    c.IndentedJSON(statusCode, gin.H{
        "feed":    feed,
        "message": "Feed de notícias encontrado com sucesso!",
    })
}
