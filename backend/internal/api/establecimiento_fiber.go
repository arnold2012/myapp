package api

import (
	"backend/internal/db"
	"github.com/gofiber/fiber/v2"
)

type EstablecimientoHandler struct {
	repo *db.EstablecimientoRepo
}

func NewEstablecimientoHandler(repo *db.EstablecimientoRepo) *EstablecimientoHandler {
	return &EstablecimientoHandler{repo: repo}
}

func (h *EstablecimientoHandler) RegisterRoutes(app *fiber.App) {
	group := app.Group("/api/establecimientos")
	group.Get("/", h.GetAll)
}

// GET /api/establecimientos
func (h *EstablecimientoHandler) GetAll(c *fiber.Ctx) error {
	data, err := h.repo.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pudo obtener los establecimientos",
		})
	}
	return c.JSON(data)
}
