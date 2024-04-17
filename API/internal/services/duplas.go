package services

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vidacalura/CS-STV/internal/models"
)

func GetRankingDuplas(c *gin.Context) {
    var rank models.Duplas

    statusCode, err := rank.GetRankingDuplas()
    if err != nil {
        c.IndentedJSON(statusCode, gin.H{"error": err.Error()})
        return
    }

    c.IndentedJSON(statusCode, gin.H{
        "ranking": rank,
        "message": "Ranking encontrado com sucesso!",
    })
}

func GetDuplaByID(c *gin.Context) {
    codDupla, err := strconv.Atoi(c.Param("codDupla"))
    if err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{
            "error": "Código de dupla inválido",
        })
        return
    }

    var dupla models.Dupla

    statusCode, err := dupla.GetDuplaByID(codDupla)
    if err != nil {
        c.IndentedJSON(statusCode, gin.H{"error": err.Error()})
        return
    }

    c.IndentedJSON(statusCode, gin.H{
        "dupla": dupla,
        "message": "Dupla encontrada com sucesso!",
    })
}
