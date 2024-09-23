package routes

import (
	"net/http"
	"strconv"
	"xyz-api-gateway/pkg/pb"
	"xyz-api-gateway/pkg/utils"

	"github.com/gin-gonic/gin"
)

func GetConsumerLimitsByConsumerId(ctx *gin.Context, c pb.ConsumerLimitServiceClient) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		errResp := utils.NewErrorResponse(http.StatusBadRequest, "Bad Request", "Failed to convert id to int")
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			errResp,
		)
		return
	}

	res, err := c.GetConsumerLimitsByConsumerId(ctx, &pb.ConsumerRequest{
		ConsumerId: id,
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
