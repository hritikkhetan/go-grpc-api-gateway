package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hritikkhetan/go-grpc-api-gateway/pkg/time/pb"
)

func CurrTime(ctx *gin.Context, c pb.TimeServiceClient) {

	res, err := c.CurrTime(context.Background(), &pb.CurrTimeRequest{})

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusAccepted, &res)
}
