package main

import (
	"log"
	"time"
	"xyz-api-gateway/pkg/config"
	// "xyz-api-gateway/pkg/modules/auth/auth"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Could not load config:", err)
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	auth.RegisterRoutes(r, &c)

	// r.Run(c.Port)

	if err := r.RunTLS(c.Port, c.SSLCert, c.SSLKey); err != nil {
		log.Fatalln("ERROR: Could not start server:", err)
	}
}
