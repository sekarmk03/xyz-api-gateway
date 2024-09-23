package routes

import (
	"net/http"
	"xyz-api-gateway/pkg/pb"
	"xyz-api-gateway/pkg/utils"

	"github.com/gin-gonic/gin"
)

func GetTransactionByContractNumber(ctx *gin.Context, c pb.TransactionServiceClient) {
	contractNumber := ctx.Param("contract_number")

	res, err := c.GetTransactionByContractNumber(ctx, &pb.TransactionContractNumberRequest{
		ContractNumber: contractNumber,
	})
	if err != nil {
		errResp := utils.GetGrpcError(err)
		ctx.AbortWithStatusJSON(
			errResp.Code,
			errResp,
		)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
