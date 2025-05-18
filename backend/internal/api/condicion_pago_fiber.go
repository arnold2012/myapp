package api

import (
    "backend/internal/db"
    "github.com/gofiber/fiber/v2"
    "strconv"
)

func RegisterCondicionPagoRoutes(app *fiber.App, repo *db.CondicionPagoRepo) {
    group := app.Group("/api/condiciones_pago")

    group.Post("/", func(c *fiber.Ctx) error {
        var cond db.CondicionPago
        if err := c.BodyParser(&cond); err != nil {
            return c.Status(400).JSON(fiber.Map{"error": "Datos inválidos"})
        }
        if err := repo.Create(&cond); err != nil {
            return c.Status(500).JSON(fiber.Map{"error": err.Error()})
        }
        return c.Status(201).JSON(cond)
    })

    group.Get("/", func(c *fiber.Ctx) error {
        condiciones, err := repo.GetAll()
        if err != nil {
            return c.Status(500).JSON(fiber.Map{"error": err.Error()})
        }
        return c.JSON(condiciones)
    })

    group.Get("/:id", func(c *fiber.Ctx) error {
        id, err := strconv.ParseInt(c.Params("id"), 10, 64)
        if err != nil {
            return c.Status(400).JSON(fiber.Map{"error": "ID inválido"})
        }

        condicion, err := repo.GetByID(id)
        if err != nil {
            return c.Status(500).JSON(fiber.Map{"error": err.Error()})
        }
        if condicion == nil {
            return c.Status(404).JSON(fiber.Map{"error": "Condición no encontrada"})
        }
        return c.JSON(condicion)
    })

    group.Put("/:id", func(c *fiber.Ctx) error {
        id, err := strconv.ParseInt(c.Params("id"), 10, 64)
        if err != nil {
            return c.Status(400).JSON(fiber.Map{"error": "ID inválido"})
        }

        var cond db.CondicionPago
        if err := c.BodyParser(&cond); err != nil {
            return c.Status(400).JSON(fiber.Map{"error": "Datos inválidos"})
        }
        cond.ID = id

        if err := repo.Update(&cond); err != nil {
            return c.Status(500).JSON(fiber.Map{"error": err.Error()})
        }

        return c.JSON(fiber.Map{"message": "Condición actualizada"})
    })

    group.Delete("/:id", func(c *fiber.Ctx) error {
        id, err := strconv.ParseInt(c.Params("id"), 10, 64)
        if err != nil {
            return c.Status(400).JSON(fiber.Map{"error": "ID inválido"})
        }

        if err := repo.Delete(id); err != nil {
            return c.Status(500).JSON(fiber.Map{"error": err.Error()})
        }

        return c.JSON(fiber.Map{"message": "Condición eliminada"})
    })
}
