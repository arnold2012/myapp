package api

import (
    "backend/internal/db"
    "github.com/gofiber/fiber/v2"
    "strconv"
)

type CategoriaHandler struct {
    repo db.CategoriaRepo
}

func NewCategoriaHandler(repo db.CategoriaRepo) *CategoriaHandler {
    return &CategoriaHandler{repo: repo}
}

func (h *CategoriaHandler) RegisterRoutes(app *fiber.App) {
    g := app.Group("/api/categorias")
    g.Get("/", h.GetAll)
    g.Get("/:id", h.GetByID)
    g.Post("/", h.Create)
    g.Put("/:id", h.Update)
    g.Delete("/:id", h.Delete)
}

func (h *CategoriaHandler) GetAll(c *fiber.Ctx) error {
    cats, err := h.repo.GetAll()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(cats)
}

func (h *CategoriaHandler) GetByID(c *fiber.Ctx) error {
    id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
    cat, err := h.repo.GetByID(id)
    if err == db.ErrCategoriaNotFound {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Categoría no encontrada"})
    }
    return c.JSON(cat)
}

func (h *CategoriaHandler) Create(c *fiber.Ctx) error {
    var cat db.Categoria
    if err := c.BodyParser(&cat); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Datos inválidos"})
    }

    id, err := h.repo.Create(cat)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    cat.IDCategoria = id
    return c.Status(fiber.StatusCreated).JSON(cat)
}

func (h *CategoriaHandler) Update(c *fiber.Ctx) error {
    id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
    var cat db.Categoria
    if err := c.BodyParser(&cat); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Datos inválidos"})
    }

    cat.IDCategoria = id
    if err := h.repo.Update(cat); err != nil {
        if err == db.ErrCategoriaNotFound {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Categoría no encontrada"})
        }
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return c.JSON(cat)
}

func (h *CategoriaHandler) Delete(c *fiber.Ctx) error {
    id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
    if err := h.repo.Delete(id); err != nil {
        if err == db.ErrCategoriaNotFound {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Categoría no encontrada"})
        }
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return c.SendStatus(fiber.StatusNoContent)
}
