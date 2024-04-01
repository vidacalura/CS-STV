package services

import(
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
