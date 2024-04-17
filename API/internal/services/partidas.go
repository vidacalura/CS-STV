package services

import (
	"net/http"
	"strconv"

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

func GetPartidasDupla(c *gin.Context) {
    codTime, err := strconv.Atoi(c.Param("codDupla"))
    if err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{
            "error": "Código de time inválido.",
        })
        return
    }

    var partidas models.Partidas

    statusCode, err := partidas.GetPartidasByTimeID(codTime)
    if err != nil {
        c.IndentedJSON(statusCode, gin.H{"error": err.Error()})
        return
    }

    c.IndentedJSON(statusCode, gin.H{
        "partidas": partidas,
        "message":  "Partidas de time encontradas com sucesso!",
    })
}
