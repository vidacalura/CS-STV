package services

import (
    "github.com/gin-gonic/gin"
    "github.com/vidacalura/CS-STV/internal/models"
)

func GetPlayerOfTheWeek(c *gin.Context) {
    var playerOfTheWeek models.PlayerOfTheWeek

    statusCode, err := playerOfTheWeek.GetUltimoPlayerOfTheWeek()
    if err != nil {
        c.IndentedJSON(statusCode, gin.H{"error": err.Error()})
        return
    }

    c.IndentedJSON(statusCode, gin.H{
        "playerData": playerOfTheWeek,
        "message":    "Jogador da semana encontrado com sucesso!",
    })
}
