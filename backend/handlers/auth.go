package handlers

import (
	"time"
	"warroom-backend/db"
	"warroom-backend/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Define secret key (In production, load from ENV)
var jwtSecret = []byte("super_secret_warroom_key")

// DTOs
type RegisterRequest struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	FullName    string `json:"full_name"`
	Role        string `json:"role"`
	SponsorID   string `json:"sponsor_id"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(c *fiber.Ctx) error {
	req := new(RegisterRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Hash Password
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not hash password"})
	}

	// Create User Struct
	user := models.User{
		Username:     req.Username,
		PasswordHash: string(hash),
		FullName:     req.FullName,
		Role:         models.UserRole(req.Role),
	}

	// Handle Sponsor
	if req.SponsorID != "" {
		parsedID, err := uuid.Parse(req.SponsorID)
		if err == nil {
			user.SponsorID = &parsedID
		}
	}

	// Save to DB
	if result := db.DB.Create(&user); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "User created successfully",
		"user_id": user.ID,
	})
}

func Login(c *fiber.Ctx) error {
	req := new(LoginRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	var user models.User
	if result := db.DB.Where("username = ?", req.Username).First(&user); result.Error != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User not found"})
	}

	// Check Password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid password"})
	}

	// Generate JWT
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString(jwtSecret)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not login"})
	}

	return c.JSON(fiber.Map{
		"token": t,
	})
}
