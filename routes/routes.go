package routes

import (
	"desafio-transacoes/controllers"
	"net/http"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
    api := router.Group("/api")
    {	api.GET("/accounts", controllers.GetAllAccounts)
		api.GET("/accounts/:accountId", controllers.GetAccount)

		api.GET("/operations",controllers.GetAllOperations)
		api.GET("/operations/:operationId",controllers.GetOperation)

		api.GET("/accounts/:accountId/transactions", controllers.GetTransactions)
		api.GET("/accounts/:accountId/transactions/:transactionId", controllers.GetTransaction)

		api.POST("/accounts", controllers.CreateAccount)
		api.POST("/operations",controllers.CreateOperation)
		api.POST("/transactions",controllers.CreateTransaction)
    }
	router.NoRoute(func(c *gin.Context) {
        c.JSON(http.StatusNotFound, gin.H{"error": "Route not found"})
    })
}