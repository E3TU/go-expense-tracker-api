package main

import (
	"expense-tracker-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/expenses", handlers.GetExpenses)
		v1.POST("/expenses", handlers.CreateExpense)
		v1.PUT("/expenses/:id", handlers.UpdateExpense)
		v1.DELETE("/expenses/:id", handlers.DeleteExpense)
	}
	r.Run(":8080")
}
