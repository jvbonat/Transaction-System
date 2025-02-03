package models

import (
	"time"
)

type Account struct {
	AccountID      int64  `gorm:"primaryKey;autoIncrement;column:Account_ID" json:"account_id"`
	DocumentNumber string `gorm:"type:varchar(20);not null;unique;column:Document_Number" json:"document_number"`
}

type OperationsType struct {
	OperationTypeID int64  `gorm:"primaryKey;autoIncrement;column:OperationTypeID" json:"operation_type_id"`
	Description     string `gorm:"type:varchar(255);not null" json:"description"`
}

type Transaction struct {
	TransactionID   int64          `gorm:"primaryKey;autoIncrement;column:Transaction_ID" json:"transaction_id"`
	AccountID       int64          `gorm:"not null;column:Account_ID" json:"account_id"`
	OperationTypeID int64          `gorm:"not null;column:OperationTypeID" json:"operation_type_id"`
	Amount          float64        `gorm:"type:decimal(10,2);not null" json:"amount"`
	EventDate       time.Time      `gorm:"type:datetime(6);not null" json:"event_date"`
	Account         Account        `gorm:"foreignKey:AccountID;references:AccountID" json:"account"`
	OperationsType  OperationsType `gorm:"foreignKey:OperationTypeID;references:OperationTypeID" json:"operation_type"`
}
