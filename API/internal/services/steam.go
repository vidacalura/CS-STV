package services

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func GetSteamUserData(c *gin.Context) {
	steamID := c.Param("steamID")
	url := "https://api.steampowered.com/ISteamUser/GetPlayerSummaries/v0002/?key=" +
		os.Getenv("steamKey") +
		"&steamids=" +
		steamID

	steamClient := http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Erro ao retornar dados de usuário da Steam.",
		})
		return
	}

	res, err := steamClient.Do(req)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Erro ao retornar dados de usuário da Steam.",
		})
		return
	}

	if res.Body != nil {
		defer res.Body.Close()
	}
}
