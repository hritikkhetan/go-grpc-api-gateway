package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hritikkhetan/go-grpc-api-gateway/pkg/config"
	"github.com/hritikkhetan/go-grpc-api-gateway/pkg/time"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	time.RegisterRoutes(r, &c)

	r.Run(c.Port)
}
