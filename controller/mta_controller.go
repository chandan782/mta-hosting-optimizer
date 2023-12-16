package controller

import (
	"net/http"
	"os"
	"strconv"

	"github.com/chandan782/mta-hosting-optimizer.git/service"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	mtaService := service.NewMtaService()

	router.GET("/getHostnames", func(c *gin.Context) {
		threshold, err := strconv.Atoi(os.Getenv("THRESHOLD"))
		if err != nil {
			threshold = 1 // Default threshold
		}

		hostnames := mtaService.GetHostnamesWithLessOrEqualActiveIPs(threshold)
		c.JSON(http.StatusOK, hostnames)
	})
}
