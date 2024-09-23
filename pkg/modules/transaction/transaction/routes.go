package transaction

import (
	"xyz-api-gateway/pkg/config"
	"xyz-api-gateway/pkg/modules/transaction/transaction/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *TransactionServiceClient {
	svc := &TransactionServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/transactions")
	routes.GET("/", svc.GetAllTransactions)
	routes.GET("/:contract_number", svc.GetTransactionByContractNumber)
	routes.POST("/", svc.CreateTransaction)

	routes = r.Group("/consumers")
	routes.GET("/:id/transactions", svc.GetTransactionsByConsumerId)

	return svc
}

func (svc *TransactionServiceClient) GetAllTransactions(ctx *gin.Context) {
	routes.GetAllTransactions(ctx, svc.Client)
}

func (svc *TransactionServiceClient) GetTransactionByContractNumber(ctx *gin.Context) {
	routes.GetTransactionByContractNumber(ctx, svc.Client)
}

func (svc *TransactionServiceClient) GetTransactionsByConsumerId(ctx *gin.Context) {
	routes.GetTransactionsByConsumerId(ctx, svc.Client)
}

func (svc *TransactionServiceClient) CreateTransaction(ctx *gin.Context) {
	routes.CreateTransaction(ctx, svc.Client)
}
