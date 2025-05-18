package api

import (
	"backend/internal/db"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type HistorialHandler struct {
	Repo *db.HistorialRepo
}

func NewHistorialHandler(repo *db.HistorialRepo) *HistorialHandler {
	return &HistorialHandler{Repo: repo}
}

func (h *HistorialHandler) RegisterRoutes(app *fiber.App) {
	historial := app.Group("/api/historial")
	historial.Get("/:order_id", h.GetHistorialByOrderID)
}

func (h *HistorialHandler) GetHistorialByOrderID(c *fiber.Ctx) error {
	orderID, err := strconv.ParseInt(c.Params("order_id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID inválido"})
	}

	historial, err := h.Repo.GetHistorialByOrderID(orderID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(historial)
}
