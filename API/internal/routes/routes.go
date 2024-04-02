package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vidacalura/CS-STV/internal/models"
	"github.com/vidacalura/CS-STV/internal/services"
	"github.com/vidacalura/CS-STV/internal/utils"
)

func NewRouter() *gin.Engine {
	env := &models.Env{
		DB: utils.ConectarBD(),
	}

	models.E = env

	r := gin.Default()

	r.Use(CORSMiddleFunc())

	dupla := r.Group("/api/duplas")
	{
		dupla.GET("/ranking", services.GetRankingDuplas)
	}

	evnt := r.Group("/api/eventos")
	{
		evnt.GET("/recentes", services.GetEventosRecentes)
	}

	part := r.Group("/api/partidas")
	{
		part.GET("/recentes", services.GetPartidasRecentes)
	}

	poftw := r.Group("/api/player-of-the-week")
	{
		poftw.GET("", services.GetPlayerOfTheWeek)
	}

	return r
}

func CORSMiddleFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		//c.Writer.Header().Set("Access-Control-Allow-Origin", "https://site-do-ngc-ai.com")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:4000")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
