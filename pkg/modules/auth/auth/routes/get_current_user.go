package routes

import (
	"net/http"
	"xyz-api-gateway/pkg/pb"
	"xyz-api-gateway/pkg/utils"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/emptypb"
)

func GetCurrentUser(ctx *gin.Context, c pb.AuthServiceClient) {
	res, err := c.GetCurrentUser(ctx, &emptypb.Empty{})

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