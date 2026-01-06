package models

import (
	"time"

	"github.com/google/uuid"
)

type TransactionType string

const (
	TxDeposit    TransactionType = "deposit"
	TxWithdrawal TransactionType = "withdrawal"
	TxCommission TransactionType = "commission_payout"
	TxPurchasePV TransactionType = "purchase_pv"
	TxTransfer   TransactionType = "transfer"
)

type Transaction struct {
	ID              uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID          uuid.UUID       `gorm:"type:uuid;not null"`
	Amount          float64         `gorm:"type:decimal(15,2);not null"`
	TransactionType TransactionType `gorm:"type:varchar(20);not null"`
	Status          string          `gorm:"default:'pending';type:varchar(20)"` // pending, completed, failed
	Description     string          `gorm:"type:text"`
	CreatedAt       time.Time
}
