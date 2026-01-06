package services

import (
	"fmt"
	"warroom-backend/db"
	"warroom-backend/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AddPV: Adds Personal Volume to a user and triggers upline GV distribution
func AddPV(userID uuid.UUID, amount float64) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		// 1. Update User's Wallet (PV)
		var wallet models.Wallet
		if err := tx.Where("user_id = ?", userID).FirstOrCreate(&wallet, models.Wallet{UserID: userID}).Error; err != nil {
			return err
		}

		wallet.PVBalance += amount
		if err := tx.Save(&wallet).Error; err != nil {
			return err
		}

		// 2. Record Transaction
		transaction := models.Transaction{
			UserID:          userID,
			Amount:          amount,
			TransactionType: models.TxPurchasePV,
			Status:          "completed",
			Description:     fmt.Sprintf("Purchase PV: %.2f", amount),
		}
		if err := tx.Create(&transaction).Error; err != nil {
			return err
		}

		// 3. Recursive GV Distribution
		if err := DistributeGV(tx, userID, amount); err != nil {
			return err
		}

		// 4. Check Tier Promotion
		if err := CheckTierPromotion(tx, userID); err != nil {
			return err
		}

		return nil
	})
}

// DistributeGV: Recursively adds Group Volume to uplines
func DistributeGV(tx *gorm.DB, startUserID uuid.UUID, amount float64) error {
	var user models.User
	// Get current user to find PlacementID (Upline)
	if err := tx.First(&user, startUserID).Error; err != nil {
		return err
	}

	// Base case: No placement (Root node)
	if user.PlacementID == nil {
		return nil
	}

	// Update Upline's GV
	var uplineWallet models.Wallet
	if err := tx.Where("user_id = ?", *user.PlacementID).FirstOrCreate(&uplineWallet, models.Wallet{UserID: *user.PlacementID}).Error; err != nil {
		return err
	}

	uplineWallet.GVBalance += amount
	if err := tx.Save(&uplineWallet).Error; err != nil {
		return err
	}

	// Recursive Call: Move up the tree
	return DistributeGV(tx, *user.PlacementID, amount)
}

// CheckTierPromotion: Logic to upgrade user tier based on points
func CheckTierPromotion(tx *gorm.DB, userID uuid.UUID) error {
	var wallet models.Wallet
	if err := tx.Where("user_id = ?", userID).First(&wallet).Error; err != nil {
		return err
	}

	var user models.User
	if err := tx.First(&user, userID).Error; err != nil {
		return err
	}

	// Simple Tier Logic (can be expanded)
	// Tier 1: Bronze (Default)
	// Tier 2: Silver (PV > 1000)
	// Tier 3: Gold (PV > 5000 OR GV > 20000)
	
	newTier := user.TierLevel

	if wallet.PVBalance >= 5000 || wallet.GVBalance >= 20000 {
		newTier = 3
	} else if wallet.PVBalance >= 1000 {
		newTier = 2
	}

	if newTier > user.TierLevel {
		user.TierLevel = newTier
		if err := tx.Save(&user).Error; err != nil {
			return err
		}
		// Log Promotion (Could be a separate transaction or log)
		fmt.Printf("🎉 User %s promoted to Tier %d\n", user.Username, newTier)
	}

	return nil
}
