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
		// POST /api/timbrado
app.Post("/api/timbrado", func(c *fiber.Ctx) error {
	var t db.Timbrado
	if err := c.BodyParser(&t); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "JSON inválido"})
	}
	if t.FechaAutorizacion.IsZero() || t.FechaInicioVigencia.IsZero() {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Fechas requeridas"})
	}
	id, err := db.CrearTimbrado(context.Background(), dbConn, t)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{"id_timbrado": id})
})

// PUT /api/timbrado/:id
app.Put("/api/timbrado/:id", func(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "ID inválido"})
	}
	var t db.Timbrado
	if err := c.BodyParser(&t); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "JSON inválido"})
	}
	t.IDTimbrado = int64(id)
	if err := db.ActualizarTimbrado(context.Background(), dbConn, t); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(http.StatusNoContent)
})

}


