package api

import (
	"backend/internal/db"
	"context"
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func RegisterTimbradoRoutes(app *fiber.App, dbConn *sql.DB) {
	// GET /api/timbrados/vigentes
	app.Get("/api/timbrados/vigentes", func(c *fiber.Ctx) error {
		lista, err := db.ObtenerTimbradosVigentes(context.Background(), dbConn)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error obteniendo timbrados vigentes: " + err.Error(),
			})
		}
		return c.JSON(lista)
	})

	// GET /api/timbrado/:establecimiento/:expedicion
	app.Get("/api/timbrado/:establecimiento/:expedicion", func(c *fiber.Ctx) error {
		estID, err1 := c.ParamsInt("establecimiento")
		expID, err2 := c.ParamsInt("expedicion")
		if err1 != nil || err2 != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "Parámetros inválidos",
			})
		}

		idTimbrado, err := db.ObtenerTimbradoPorEstablecimientoYExpedicion(c.Context(), dbConn, int64(estID), int64(expID))
		if err != nil {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"id_timbrado": idTimbrado,
		})
	})
}
