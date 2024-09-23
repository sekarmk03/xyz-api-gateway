package routes

import (
	"context"
	"net/http"
	"strconv"
	"xyz-api-gateway/pkg/pb"
	"xyz-api-gateway/pkg/utils"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
)

type CreateConsumerLimitRequestBody struct {
	Tenor       uint32 `json:"tenor"`
	LimitAmount uint64 `json:"limit_amount"`
}

func CreateConsumerLimit(ctx *gin.Context, c pb.ConsumerLimitServiceClient) {
	authorizationHeader := ctx.GetHeader("Authorization")
	grpcCtx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs("authorization", authorizationHeader))

	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		errResp := utils.NewErrorResponse(http.StatusBadRequest, "Bad Request", "Failed to convert id to int")
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			errResp,
		)
		return
	}

	b := CreateConsumerLimitRequestBody{}
	if err := ctx.BindJSON(&b); err != nil {
		errResp := utils.NewErrorResponse(http.StatusBadRequest, "Bad Request", "Invalid request body")
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			errResp,
		)
		return
	}

	res, err := c.CreateConsumerLimit(grpcCtx, &pb.ConsumerLimit{
		ConsumerId:  id,
		Tenor:       b.Tenor,
		LimitAmount: b.LimitAmount,
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
