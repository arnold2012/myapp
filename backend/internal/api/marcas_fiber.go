package api

import (
	"backend/internal/db"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"database/sql"
)

type MarcaHandler struct {
	Repo *db.MarcaRepo
}

func RegisterMarcaRoutes(app *fiber.App, repo *db.MarcaRepo) {
	handler := &MarcaHandler{Repo: repo}
	api := app.Group("/api/marcas")

	api.Get("/", handler.GetAll)
	api.Get("/:id", handler.GetByID)
	api.Post("/", handler.Create)
	api.Put("/:id", handler.Update)
	api.Delete("/:id", handler.Delete)
}

func (h *MarcaHandler) GetAll(c *fiber.Ctx) error {
	data, err := h.Repo.GetAll()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(data)
}

func (h *MarcaHandler) GetByID(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	data, err := h.Repo.GetByID(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	if data == nil {
		return c.Status(404).JSON(fiber.Map{"error": "Marca no encontrada"})
	}
	return c.JSON(data)
}

func (h *MarcaHandler) Create(c *fiber.Ctx) error {
	var input db.Marca
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "JSON inválido"})
	}
	id, err := h.Repo.Create(input)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	input.IDMarca = id
	return c.Status(201).JSON(input)
}

func (h *MarcaHandler) Update(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	var input db.Marca
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "JSON inválido"})
	}
	input.IDMarca = id
	if err := h.Repo.Update(input); err != nil {
		if err == sql.ErrNoRows {
			return c.Status(404).JSON(fiber.Map{"error": "Marca no encontrada"})
		}
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(input)
}

func (h *MarcaHandler) Delete(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	if err := h.Repo.Delete(id); err != nil {
		if err == sql.ErrNoRows {
			return c.Status(404).JSON(fiber.Map{"error": "Marca no encontrada"})
		}
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(204)
}
