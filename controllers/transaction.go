package controllers

import (
	"API/database"
	"API/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ExibeTodasTransactions(c *gin.Context) {
	var transactions []models.Transaction
	database.DB.Find(&transactions)
	c.JSON(200, transactions)

}

func NewTransaction(c *gin.Context) {
	var transaction models.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&transaction)
	c.JSON(http.StatusOK, transaction)
}

func BuscaTransactionPorID(c *gin.Context) {
	var transaction models.Transaction
	id := c.Params.ByName("id")
	database.DB.First(&transaction, id)
	if transaction.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "transaction não encontrada!"})
		return
	}
	c.JSON(http.StatusOK, transaction)
}

func DeletaTransaction(c *gin.Context) {
	var transaction models.Transaction
	id := c.Params.ByName("id")
	database.DB.Delete(&transaction, id)
	c.JSON(http.StatusOK, gin.H{"data": "Transaction deletado com sucesso!"})
}

func EditaTransaction(c *gin.Context) {
	var transaction models.Transaction
	id := c.Params.ByName("id")
	database.DB.First(&transaction, id)

	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Model(&transaction).UpdateColumns(transaction)
	c.JSON(http.StatusOK, transaction)
}

func Transaction(c *gin.Context) {
	var account1 models.Account
	var transaction1 models.Transaction

	transaction1.RunningBalance = account1.CurrentBalance - transaction1.DebitedAmount
	transaction1.RunningBalance = account1.CurrentBalance + transaction1.CreditedAmount
	id := c.Params.ByName("id")
	database.DB.First(&transaction1.ID, id)

	if err := c.ShouldBindJSON(&transaction1); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Model(&transaction1.RunningBalance).UpdateColumns(transaction1.RunningBalance)
	c.JSON(http.StatusOK, transaction1)
}
