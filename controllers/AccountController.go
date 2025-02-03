package controllers

import (
	"desafio-transacoes/db"
	"desafio-transacoes/models"
	"net/http"
	"unicode"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func CreateAccount(c *gin.Context) {
	var account models.Account

	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	isValidDocumentNumber := isValidDocumentNumber(account.DocumentNumber)

	if !isValidDocumentNumber {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid document number"})
		return
	}

	var existingAccount models.Account
	result := db.DB.Where("Document_Number = ?", account.DocumentNumber).First(&existingAccount)

	if result.Error == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Account already exists"})
		return
	}

	if result.Error == gorm.ErrRecordNotFound {
		if err := db.DB.Create(&account).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error when creating account"})
			return
		}
		c.JSON(http.StatusCreated, account)
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})

}

func isNumericText(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func isValidDocumentNumber(documentNumber string) bool {

	if documentNumber == "" {
		return false
	}

	if !isNumericText(documentNumber) {
		return false
	}

	if len(documentNumber) != 11 {
		return false
	}

	return true
}

func GetAccount(c *gin.Context) {
	accountId := c.Param("accountId")
	if accountId == "" {
		GetAllAccounts(c)
		return
	}

	var account models.Account
	if err := db.DB.First(&account, accountId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"account_id":      account.AccountID,
		"document_number": account.DocumentNumber,
	})
}

func GetAllAccounts(c *gin.Context) {
	var accounts []models.Account
	if err := db.DB.Find(&accounts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error when searching for account"})
		return
	}

	if len(accounts) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No accounts have been found"})
		return
	}

	c.JSON(http.StatusOK, accounts)
}
