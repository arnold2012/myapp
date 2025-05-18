package api

import (
  "backend/internal/db"
  "database/sql"
  "github.com/gofiber/fiber/v2"
)

func RegisterDocContableRoutes(app *fiber.App, dbConn *sql.DB) {
  // Establecimientos
  estRepo := db.NewEstablecimientoRepo(dbConn)
  app.Get("/api/establecimientos", func(c *fiber.Ctx) error {
    list, err := estRepo.GetAll()
    if err != nil {
      return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(list)
  })

  // Puntos de expedición
  puntoRepo := db.NewPuntoRepo(dbConn)
  app.Get("/api/puntos_expedicion", func(c *fiber.Ctx) error {
    list, err := puntoRepo.GetAll()
    if err != nil {
      return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(list)
  })
}
