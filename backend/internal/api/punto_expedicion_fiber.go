package api

import (
	"backend/internal/db"
	"github.com/gofiber/fiber/v2"
)

type PuntoExpedicionHandler struct {
	repo *db.PuntoRepo
}

func NewPuntoExpedicionHandler(repo *db.PuntoRepo) *PuntoExpedicionHandler {
	return &PuntoExpedicionHandler{repo: repo}
}

func (h *PuntoExpedicionHandler) RegisterRoutes(app *fiber.App) {
	group := app.Group("/api/puntos")
	group.Get("/", h.GetAll)
}

func (h *PuntoExpedicionHandler) GetAll(c *fiber.Ctx) error {
	puntos, err := h.repo.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al obtener puntos de expedición",
		})
	}
	return c.JSON(puntos)
}
