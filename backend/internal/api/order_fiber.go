package api

import (
	"backend/internal/db"
	"context"
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func RegisterOrderRoutes(app *fiber.App, dbConn *sql.DB) {
	// Endpoint para crear una orden
	app.Post("/api/orders", func(c *fiber.Ctx) error {
		var params db.InsertOrderParams

		if err := c.BodyParser(&params); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "JSON inválido",
			})
		}

		if len(params.Items) == 0 {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "Se requiere al menos un item",
			})
		}

		// Ahora InsertSalesOrder devuelve el ID de la orden
		orderID, err := db.InsertSalesOrder(context.Background(), dbConn, params)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(http.StatusCreated).JSON(fiber.Map{
			"message":  "Orden creada con éxito",
			"id_order": orderID,
		})
	})

	// Endpoint para imprimir una orden (consultar con detalle)
	app.Get("/api/orders/:id/imprimir", func(c *fiber.Ctx) error {
		orderID, err := c.ParamsInt("id")
		if err != nil || orderID <= 0 {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "ID de orden inválido",
			})
		}

		orderDetails, err := db.GetOrderForPrint(context.Background(), dbConn, int64(orderID))
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if len(orderDetails) == 0 {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"error": "Orden no encontrada",
			})
		}

		return c.Status(http.StatusOK).JSON(orderDetails)
	})
}


// package api

// import (
//     "backend/internal/db"
//     "database/sql"
//     "github.com/gofiber/fiber/v2"
//     "context"
//     "net/http"
// )

// func RegisterOrderRoutes(app *fiber.App, dbConn *sql.DB) {
//     app.Post("/api/orders", func(c *fiber.Ctx) error {
//         var params db.InsertOrderParams

//         if err := c.BodyParser(&params); err != nil {
//             return c.Status(http.StatusBadRequest).JSON(fiber.Map{
//                 "error": "JSON inválido",
//             })
//         }

//         if len(params.Items) == 0 {
//             return c.Status(http.StatusBadRequest).JSON(fiber.Map{
//                 "error": "Se requiere al menos un item",
//             })
//         }

//         err := db.InsertSalesOrder(context.Background(), dbConn, params)
//         if err != nil {
//             return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
//                 "error": err.Error(),
//             })
//         }

//         return c.Status(http.StatusCreated).JSON(fiber.Map{
//             "message": "Orden creada con éxito",
//         })
//     })
// }
