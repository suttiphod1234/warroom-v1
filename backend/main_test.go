package main

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	app := fiber.New()
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
			"service": "backend",
		})
	})

	req := httptest.NewRequest("GET", "/health", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)
}
