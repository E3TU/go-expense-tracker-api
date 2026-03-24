package handlers

import (
	"expense-tracker-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var expenses = []models.Expense{}

func GetExpenses(c *gin.Context) {
	c.JSON(http.StatusOK, expenses)
}
