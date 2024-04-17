package routes

import (
	"time"

	"github.com/gin-contrib/cors"
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
		dupla.GET("/:codDupla", services.GetDuplaByID)
	}

	evnt := r.Group("/api/eventos")
	{
		evnt.GET("/recentes", services.GetEventosRecentes)
		evnt.GET("/dupla/:codDupla", services.GetEventosByDuplaID)
	}

	notc := r.Group("/api/noticias")
	{
		notc.GET("/feed", services.GetFeedNoticias)
		notc.GET("/:codNotc", services.GetNoticiaByID)
	}

	part := r.Group("/api/partidas")
	{
		part.GET("/recentes", services.GetPartidasRecentes)
		part.GET("/dupla/:codDupla", services.GetPartidasDupla)
	}

	poftw := r.Group("/api/player-of-the-week")
	{
		poftw.GET("", services.GetPlayerOfTheWeek)
	}

	return r
}

func CORSMiddleFunc() gin.HandlerFunc {
	return cors.New(cors.Config{
		//AllowOrigins:     []string{"http://127.0.0.1:5500"},
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
