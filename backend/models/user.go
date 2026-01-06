package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRole string

const (
	RoleAdmin UserRole = "admin"
	RoleMP    UserRole = "mp"
	RoleAgent UserRole = "agent"
	RoleVoter UserRole = "voter"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Username     string    `gorm:"unique;not null"`
	PasswordHash string    `gorm:"not null"`
	FullName     string    `gorm:"not null"`
	Role         UserRole  `gorm:"type:user_role;default:'voter'"`
	PhoneNumber  string
	LineID       string

	// Hierarchy
	SponsorID   *uuid.UUID `gorm:"type:uuid"`
	Sponsor     *User      `gorm:"foreignKey:SponsorID"`
	PlacementID *uuid.UUID `gorm:"type:uuid"`
	Placement   *User      `gorm:"foreignKey:PlacementID"`

	// Status
	IsActive  bool `gorm:"default:true"`
	IsBanned  bool `gorm:"default:false"`
	TierLevel int  `gorm:"default:1"` // 1=Bronze, 2=Silver, etc.

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// Hook to ensure UUID is generated before creation if skipped by DB
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return
}
