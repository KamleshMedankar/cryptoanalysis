// routes/routes.go
package routes

import (
	"crypto_analysis/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/analyze", controllers.AnalyzeWallet)
	router.GET("/status", controllers.CheckStatus)

}