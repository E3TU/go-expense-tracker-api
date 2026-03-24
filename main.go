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
		// v1.POST()
		// v1.PUT()
		// v1.DELETE()
	}
	r.Run(":8080")
}
