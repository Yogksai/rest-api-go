package app

import (
	"log"

	"github.com/Yogksai/rest-api-go/config"
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.New()
	cfg := config.NewConfig()

	if err := r.Run(cfg.HTTPServer.Host + ":" + cfg.HTTPServer.Port); err != nil {
		log.Fatal("failed to run server")
		return
	}
}
