package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Wallet struct {
	UserID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	User              User      `gorm:"foreignKey:UserID"`
	PVBalance         float64   `gorm:"default:0.00;type:decimal(15,2)"` // Personal Volume
	GVBalance         float64   `gorm:"default:0.00;type:decimal(15,2)"` // Group Volume
	CommissionBalance float64   `gorm:"default:0.00;type:decimal(15,2)"` // Withdrawable Cash
	LastUpdated       time.Time `gorm:"autoUpdateTime"`
}

// Hook to create wallet when user is created (can be handled in User hook too)
func (w *Wallet) BeforeCreate(tx *gorm.DB) (err error) {
	return
}
