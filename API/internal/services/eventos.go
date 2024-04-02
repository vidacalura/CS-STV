package services

import (
	"github.com/gin-gonic/gin"
	"github.com/vidacalura/CS-STV/internal/models"
)

func GetEventosRecentes(c *gin.Context) {
    var eventos models.Eventos

    statusCode, err := eventos.GetEventosRecentes()
    if err != nil {
        c.IndentedJSON(statusCode, gin.H{"error": err.Error()})
        return
    }

    c.IndentedJSON(statusCode, gin.H{
        "eventos": eventos,
        "message": "Eventos encontrados com sucesso!",
    })
}
