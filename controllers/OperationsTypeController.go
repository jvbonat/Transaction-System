package controllers

import (
	"desafio-transacoes/db"
	"desafio-transacoes/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateOperation(c *gin.Context) {
	var operation models.OperationsType

	if err := c.ShouldBindJSON(&operation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	if operation.Description == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Description is mandatory"})
		return
	}

	var existing models.OperationsType
	result := db.DB.Where("description = ?", operation.Description).First(&existing)

	if result.Error == nil {
		c.JSON(http.StatusConflict, gin.H{
			"error":             "Operation with this description already exists",
			"operation_type_id": existing.OperationTypeID,
			"description":       existing.Description,
		})
		return
	}
	if err := db.DB.Create(&operation).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating operation"})
		return
	}

	c.JSON(http.StatusCreated, operation)
}

func GetAllOperations(c *gin.Context) {
	var operations []models.OperationsType
	if err := db.DB.Find(&operations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error when searching for operations"})
		return
	}
	if len(operations) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No operations have been found"})
		return
	}
	c.JSON(http.StatusOK, operations)
}

func GetOperation(c *gin.Context) {
	operationId := c.Param("operationId")
	if operationId == "" {
		GetAllOperations(c)
		return
	}
	var operation models.OperationsType
	if err := db.DB.First(&operation,operationId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Operation not found"})
		return
	}
	c.JSON(http.StatusOK,operation)
}
