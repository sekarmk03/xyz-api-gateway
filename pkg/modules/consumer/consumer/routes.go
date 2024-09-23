package consumer

import (
	"xyz-api-gateway/pkg/config"
	"xyz-api-gateway/pkg/modules/consumer/consumer/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ConsumerServiceClient {
	svc := &ConsumerServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/consumers")
	routes.GET("/", svc.GetAllConsumers)
	routes.GET("/:id", svc.GetConsumerById)
	routes.POST("/", svc.CreateConsumer)
	routes.PUT("/:id", svc.UpdateConsumer)
	routes.DELETE("/:id", svc.DeleteConsumer)

	return svc
}

func (svc *ConsumerServiceClient) GetAllConsumers(ctx *gin.Context) {
	routes.GetAllConsumers(ctx, svc.Client)
}

func (svc *ConsumerServiceClient) GetConsumerById(ctx *gin.Context) {
	routes.GetConsumerById(ctx, svc.Client)
}

func (svc *ConsumerServiceClient) DeleteConsumer(ctx *gin.Context) {
	routes.DeleteConsumer(ctx, svc.Client)
}

func (svc *ConsumerServiceClient) UpdateConsumer(ctx *gin.Context) {
	routes.UpdateConsumer(ctx, svc.Client)
}

func (svc *ConsumerServiceClient) CreateConsumer(ctx *gin.Context) {
	routes.CreateConsumer(ctx, svc.Client)
}