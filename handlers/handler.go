package handlers

import (
	"expense-tracker-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

var expenses = []models.Expense{}

func GetExpenses(c *gin.Context) {
	c.JSON(http.StatusOK, expenses)
}

func CreateExpense(c *gin.Context) {
	var input models.Expense
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var guid = xid.New()

	input.ID = guid.String()
	expenses = append(expenses, input)

	c.JSON(http.StatusCreated, input)
}

func UpdateExpense(c *gin.Context) {
	id := c.Param("id")
	var input models.Expense
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, expense := range expenses {

		if expense.ID == id {
			expenses[i].Description = input.Description
			expenses[i].Amount = input.Amount
			c.JSON(http.StatusOK, expenses[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "expense not found"})
}

func DeleteExpense(c *gin.Context) {
	id := c.Param("id")

	for i, expense := range expenses {

		if expense.ID == id {
			expenses = append(expenses[:i], expenses[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Expense deleted successfully"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "expense not found"})
}
