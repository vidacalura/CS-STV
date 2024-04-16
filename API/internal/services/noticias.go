package services

import (
	"net/http"
	"strconv"

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

// Retorna uma notícia com base em seu ID
func GetNoticiaByID(c *gin.Context) {
	codNotc, err := strconv.Atoi(c.Param("codNotc"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Código de notícia inválido.",
		})
		return
	}

	var noticia models.Noticia
	statusCode, err := noticia.GetNoticiaByID(codNotc)
	if err != nil {
		c.IndentedJSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(statusCode, gin.H{
		"noticia": noticia,
		"message": "Notícia encontrada com sucesso!",
	})
}
