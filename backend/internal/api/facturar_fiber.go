package api

import (
	"backend/internal/db"
	"backend/internal/utils"
	"database/sql"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

func RegisterFacturaRoutes(app *fiber.App, dbConn *sql.DB) {
	// Ruta para generar factura
	app.Post("/api/facturar", func(c *fiber.Ctx) error {
		var params db.FacturaParams
		if err := json.Unmarshal(c.Body(), &params); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "JSON inválido: " + err.Error(),
			})
		}
		if len(params.Items) == 0 {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "Debe incluir al menos un ítem para la factura.",
			})
		}
		if err := db.InsertFactura(c.Context(), dbConn, params); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": "No se pudo insertar la factura: " + err.Error(),
			})
		}
		return c.Status(http.StatusCreated).JSON(fiber.Map{
			"message": "Factura generada con éxito.",
		})
	})

	// Ruta para buscar una orden de venta por número
	app.Get("/api/order/search/:numero", func(c *fiber.Ctx) error {
		numStr := c.Params("numero")
		numero, err := strconv.ParseInt(numStr, 10, 64)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "Número de orden inválido",
			})
		}
		detalles, err := db.SearchOrderByNumero(c.Context(), dbConn, numero)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error al buscar la orden: " + err.Error(),
			})
		}
		if len(detalles) == 0 {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"error": "Orden no encontrada",
			})
		}
		return c.JSON(detalles)
	})

	// Ruta para obtener el próximo número de factura
	app.Get("/api/facturar/next-numero/:establecimiento/:expedicion", func(c *fiber.Ctx) error {
		numEst := c.Params("establecimiento")
		numExp := c.Params("expedicion")

		numero, err := utils.ObtenerProximoNumeroFactura(dbConn, numEst, numExp)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": "No se pudo obtener el número de factura: " + err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"next_numero_factura": numero,
		})
	})
}
