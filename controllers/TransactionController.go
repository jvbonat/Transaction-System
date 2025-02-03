package controllers

import (
	"desafio-transacoes/db"
	"desafio-transacoes/models"
	"time"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateTransaction(c *gin.Context) {
	var transaction models.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	if transaction.Amount == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "transaction amount cannot be zero"})
		return
	}

	var account models.Account
	if err := db.DB.First(&account, "Account_ID = ?", transaction.AccountID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}

	var operation models.OperationsType
	if err := db.DB.First(&operation, transaction.OperationTypeID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Operation type not found"})
		return
	}

	if isDebitOperation(operation.OperationTypeID) && transaction.Amount > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Debit transactions must have negative values"})
		return
	}

	if isCreditOperation(operation.OperationTypeID) && transaction.Amount < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Credit transactions must have positive values"})
		return
	}

	transaction.EventDate = time.Now()

	if err := db.DB.Create(&transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error when creating transaction"})
		return
	}

	c.JSON(http.StatusCreated, "Transaction registered")
}

func GetTransactions(c *gin.Context) {
	var transactions []models.Transaction
	accountId := c.Param("accountId")
	if accountId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing parameter accountId"})
	}
	if err := db.DB.Preload("Account").Preload("OperationsType").Where("account_id = ?", accountId).Find(&transactions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error when searching for transactions"})
		return
	}
	if len(transactions) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No transactions have been found"})
		return
	}
	c.JSON(http.StatusOK, transactions)
}

func GetTransaction(c *gin.Context) {
	var transaction models.Transaction
	accountId := c.Param("accountId")
	transactionId := c.Param("transactionId")
	if accountId == "" || transactionId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing parameters transactionId and accountId"})
		return
	}
	if err := db.DB.Preload("Account").Preload("OperationsType").
		Where("Account_ID = ? AND Transaction_ID = ?", accountId, transactionId).
		First(&transaction).Error; err != nil {

		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return

	}
	c.JSON(http.StatusOK, transaction)
}

func isDebitOperation(operationTypeID int64) bool {
	return operationTypeID == 1 || operationTypeID == 2 || operationTypeID == 3
}

func isCreditOperation(operationTypeID int64) bool {
	return operationTypeID == 4
}
