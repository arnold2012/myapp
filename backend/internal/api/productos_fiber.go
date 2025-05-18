// backend/internal/api/productos_fiber.go
package api

import (
	"backend/internal/db"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type ProductoHandler struct {
	Repo *db.ProductoRepo
}

func NewProductoHandler(repo *db.ProductoRepo) *ProductoHandler {
	return &ProductoHandler{Repo: repo}
}

func (h *ProductoHandler) RegisterRoutes(app *fiber.App) {
	api := app.Group("/api/productos")
	api.Post("/", h.CreateProducto)
	api.Get("/", h.GetAllProductos)
	api.Get("/search", h.SearchProducto)
	api.Get("/next-code", h.GetNextCodItem) // Nuevo endpoint
	api.Get("/:id", h.GetProductoByID)
	api.Put("/:id", h.UpdateProducto)
	api.Delete("/:id", h.DeleteProducto)
}

func (h *ProductoHandler) CreateProducto(c *fiber.Ctx) error {
	var p db.Producto
	if err := c.BodyParser(&p); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Datos inválidos"})
	}
	if err := h.Repo.Create(p); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Producto creado con éxito"})
}

func (h *ProductoHandler) GetAllProductos(c *fiber.Ctx) error {
	productos, err := h.Repo.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(productos)
}

func (h *ProductoHandler) GetProductoByID(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID inválido"})
	}
	producto, err := h.Repo.GetByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Producto no encontrado"})
	}
	return c.JSON(producto)
}

func (h *ProductoHandler) UpdateProducto(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID inválido"})
	}
	var p db.Producto
	if err := c.BodyParser(&p); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Datos inválidos"})
	}
	if err := h.Repo.Update(id, p); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Producto actualizado con éxito"})
}

func (h *ProductoHandler) DeleteProducto(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID inválido"})
	}
	if err := h.Repo.Delete(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Producto eliminado con éxito"})
}

func (h *ProductoHandler) SearchProducto(c *fiber.Ctx) error {
	query := c.Query("q")
	productos, err := h.Repo.Search(query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(productos)
}

func (h *ProductoHandler) GetNextCodItem(c *fiber.Ctx) error {
	code, err := h.Repo.GetNextCodItem()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pudo generar el siguiente código",
		})
	}
	return c.JSON(fiber.Map{"next_cod_item": code})
}
