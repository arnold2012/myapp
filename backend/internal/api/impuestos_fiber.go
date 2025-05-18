package api

import (
	"backend/internal/db"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type ImpuestosHandler struct {
	repo *db.ImpuestosRepo
}

func NewImpuestosHandler(repo *db.ImpuestosRepo) *ImpuestosHandler {
	return &ImpuestosHandler{repo: repo}
}

func (h *ImpuestosHandler) RegisterRoutes(app *fiber.App) {
	api := app.Group("/api/impuestos")
	api.Get("/", h.getAll)
	api.Get("/:id", h.getByID)
	api.Post("/", h.create)
	api.Put("/:id", h.update)
	api.Delete("/:id", h.delete)
}

// GET /api/impuestos
func (h *ImpuestosHandler) getAll(c *fiber.Ctx) error {
	impuestos, err := h.repo.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al obtener los impuestos",
		})
	}
	return c.JSON(impuestos)
}

// GET /api/impuestos/:id
func (h *ImpuestosHandler) getByID(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil || id <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID inválido",
		})
	}
	impuesto, err := h.repo.GetByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al buscar el impuesto",
		})
	}
	if impuesto == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Impuesto no encontrado",
		})
	}
	return c.JSON(impuesto)
}

// POST /api/impuestos
func (h *ImpuestosHandler) create(c *fiber.Ctx) error {
	var impuesto db.Impuestos
	if err := c.BodyParser(&impuesto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "JSON inválido",
		})
	}
	id, err := h.repo.Insert(impuesto)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al insertar el impuesto",
		})
	}
	impuesto.ID = id
	return c.Status(fiber.StatusCreated).JSON(impuesto)
}

// PUT /api/impuestos/:id
func (h *ImpuestosHandler) update(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil || id <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID inválido",
		})
	}
	var impuesto db.Impuestos
	if err := c.BodyParser(&impuesto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "JSON inválido",
		})
	}
	impuesto.ID = id
	if err := h.repo.Update(impuesto); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al actualizar el impuesto",
		})
	}
	return c.JSON(impuesto)
}

// DELETE /api/impuestos/:id
func (h *ImpuestosHandler) delete(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil || id <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID inválido",
		})
	}
	if err := h.repo.Delete(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al eliminar el impuesto",
		})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
