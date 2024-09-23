package routes

import (
	"net/http"
	"xyz-api-gateway/pkg/pb"
	"xyz-api-gateway/pkg/utils"

	"github.com/gin-gonic/gin"
)

type CreateTransactionRequestBody struct {
	ConsumerId  uint64 `json:"consumer_id"`
	Tenor       uint32 `json:"tenor"`
	Otr         uint64 `json:"otr"`
	AdminFee    uint64 `json:"admin_fee"`
	Installment uint64 `json:"installment"`
	Interest    uint64 `json:"interest"`
	AssetName   string `json:"asset_name"`
}

func CreateTransaction(ctx *gin.Context, c pb.TransactionServiceClient) {
	b := CreateTransactionRequestBody{}
	if err := ctx.BindJSON(&b); err != nil {
		errResp := utils.NewErrorResponse(http.StatusBadRequest, "Bad Request", "Invalid request body")
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			errResp,
		)
		return
	}

	res, err := c.CreateTransaction(ctx, &pb.Transaction{
		ConsumerId:  b.ConsumerId,
		Tenor:       b.Tenor,
		Otr:         b.Otr,
		AdminFee:    b.AdminFee,
		Installment: b.Installment,
		Interest:    b.Interest,
		AssetName:   b.AssetName,
	})
	if err != nil {
		errResp := utils.GetGrpcError(err)
		ctx.AbortWithStatusJSON(
			errResp.Code,
			errResp,
		)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
