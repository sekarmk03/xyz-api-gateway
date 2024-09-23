package consumer_limit

import (
	"xyz-api-gateway/pkg/config"
	"xyz-api-gateway/pkg/modules/consumer/consumer_limit/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ConsumerLimitServiceClient {
	svc := &ConsumerLimitServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/consumers")
	routes.GET("/:id/limits", svc.GetConsumerLimitsByConsumerId)
	routes.POST("/limit", svc.CreateConsumerLimit)

	return svc
}

func (svc *ConsumerLimitServiceClient) GetConsumerLimitsByConsumerId(ctx *gin.Context) {
	routes.GetConsumerLimitsByConsumerId(ctx, svc.Client)
}

func (svc *ConsumerLimitServiceClient) CreateConsumerLimit(ctx *gin.Context) {
	routes.CreateConsumerLimit(ctx, svc.Client)
}
