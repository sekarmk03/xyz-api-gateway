package main

import (
	"log"
	"time"
	"xyz-api-gateway/pkg/config"

	"xyz-api-gateway/pkg/modules/consumer/consumer"
	"xyz-api-gateway/pkg/modules/consumer/consumer_limit"
	"xyz-api-gateway/pkg/modules/transaction/transaction"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Could not load config:", err)
	}

	r := gin.Default()

	r.RedirectTrailingSlash = false

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{c.FeOriginUrl},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	consumer.RegisterRoutes(r, &c)
	consumer_limit.RegisterRoutes(r, &c)
	transaction.RegisterRoutes(r, &c)

	r.Run(c.Port)

	// if err := r.RunTLS(c.Port, c.SSLCert, c.SSLKey); err != nil {
	// 	log.Fatalln("ERROR: Could not start server:", err)
	// }
}
