package api

import (
	"net/http"

	"goapi/src/settings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ServerGin() {
	r := gin.Default()

	ginConf := cors.DefaultConfig()
	ginConf.AllowOrigins = []string{"http://localhost:4200"}

	r.Use(cors.New(ginConf))

	r.StaticFS("/", http.Dir(settings.SettingsIns.HlsPath))

	r.POST("/upload", uploadHandler)

	r.Run(":8080")
}
