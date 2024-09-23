package auth

import (
	"xyz-api-gateway/pkg/config"
	"xyz-api-gateway/pkg/modules/auth/auth/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *AuthServiceClient {
	svc := &AuthServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/auth")
	routes.POST("/login", svc.Login)
	routes.GET("/whoami", svc.GetCurrentUser)

	return svc
}

func (svc *AuthServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}

func (svc *AuthServiceClient) GetCurrentUser(ctx *gin.Context) {
	routes.GetCurrentUser(ctx, svc.Client)
}
