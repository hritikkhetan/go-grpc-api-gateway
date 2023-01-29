package time

import (
	"github.com/gin-gonic/gin"
	"github.com/hritikkhetan/go-grpc-api-gateway/pkg/config"
	"github.com/hritikkhetan/go-grpc-api-gateway/pkg/time/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/time")
	routes.GET("/currTime", svc.CurrTime)

}

func (svc *ServiceClient) CurrTime(ctx *gin.Context) {
	routes.CurrTime(ctx, svc.Client)
}
