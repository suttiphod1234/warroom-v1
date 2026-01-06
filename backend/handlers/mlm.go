package handlers

import (
	"warroom-backend/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type PurchaseRequest struct {
	UserID string  `json:"user_id"`
	Amount float64 `json:"amount"` // In real app, this might be product items
}

func Purchase(c *fiber.Ctx) error {
	req := new(PurchaseRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid User ID"})
	}

	// Call MLM Service
	if err := services.AddPV(userID, req.Amount); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "Purchase successful. PV added and GV distributed.",
		"amount":  req.Amount,
	})
}
