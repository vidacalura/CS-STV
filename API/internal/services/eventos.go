package services

import (
	"net/http"
	"strconv"

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

func GetEventosByDuplaID(c *gin.Context) {
    codDupla, err := strconv.Atoi(c.Param("codDupla"))
    if err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{
            "error": "Código de dupla inválido",
        })
        return
    }

    var eventos models.Eventos

    statusCode, err := eventos.GetEventosByTimeID(codDupla)
    if err != nil {
        c.IndentedJSON(statusCode, gin.H{"error": err.Error()})
        return
    }

    c.IndentedJSON(statusCode, gin.H{
        "eventos": eventos,
        "message": "Eventos encontrados com sucesso!",
    })
}
