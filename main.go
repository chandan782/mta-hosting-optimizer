package main

import (
	"github.com/chandan782/mta-hosting-optimizer.git/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	controller.RegisterRoutes(router)

	router.Run(":8080")
}
