package tests

import (
	"bytes"
	"desafio-transacoes/controllers"
	"desafio-transacoes/db"
	"desafio-transacoes/models"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func setupTestDB() {
	err := db.InitDB()
	if err != nil {
		panic("Database connection has failed")
	}
}

func TestCreateAccount(t *testing.T) {
	gin.SetMode(gin.TestMode)
	setupTestDB()

	t.Run("Criar conta", func(t *testing.T) {
	
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		body := `{"document_number": "111207019"}`

		req, _ := http.NewRequest("POST", "/accounts", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		c.Request = req

		controllers.CreateAccount(c)

		assert.Equal(t, http.StatusCreated, w.Code)

		var account models.Account
		db.DB.Where("document_number = ?", "111207051").First(&account)

		assert.Equal(t, "111207051", account.DocumentNumber)
	})
}

func TestCreateTransaction(t *testing.T) {

	gin.SetMode(gin.TestMode)
	setupTestDB()

	accountID := 1
	operationTypeID := 4

	var account models.Account
	if err := db.DB.First(&account, accountID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		t.Fatalf("Account with ID %d not found", accountID)
	}

	var operationType models.OperationsType
	if err := db.DB.First(&operationType, operationTypeID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		t.Fatalf("Operation type with ID %d not found", operationTypeID)
	}

	transactionData := models.Transaction{
		AccountID:       int64(accountID),
		OperationTypeID: int64(operationTypeID),
		Amount:          100.00,
		EventDate:       time.Now(),
	}

	jsonData, err := json.Marshal(transactionData)
	assert.NoError(t, err, "Error when serializing JSON transaction")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req, err := http.NewRequest("POST", "/transactions", bytes.NewBuffer(jsonData))
	assert.NoError(t, err, "Error when creating HTTP request")
	req.Header.Set("Content-Type", "application/json")

	c.Request = req

	controllers.CreateTransaction(c)

	assert.Equal(t, http.StatusCreated, w.Code, "Status Returned: %d", w.Code)

	assert.NotEmpty(t, w.Body.String(), "Empty response")
}
