package services

import (
	"github.com/gin-gonic/gin"
	"github.com/vidacalura/CS-STV/internal/models"
)

func GetPartidasRecentes(c *gin.Context) {
    var partidas models.Partidas

    statusCode, err := partidas.GetPartidasRecentes()
    if err != nil {
        c.IndentedJSON(statusCode, gin.H{"error": err.Error()})
        return
    }

    c.IndentedJSON(statusCode, gin.H{
        "partidas": partidas,
        "message":  "Partidas recentes encontradas com sucesso!",
    })
}
