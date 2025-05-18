package router

import (
    "backend/internal/api"
    "backend/internal/db"
    "database/sql"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupFiberRoutes(database *sql.DB) *fiber.App {
    app := fiber.New()

    app.Use(cors.New(cors.Config{
        AllowOrigins: "*",
        AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
        AllowHeaders: "Content-Type",
    }))

    // Contactos
    contactRepo := db.NewContactRepo(database)
    contactHandler := api.NewContactHandler(contactRepo)
    contactHandler.RegisterRoutes(app)

    //  Agregar el registro de órdenes
    api.RegisterOrderRoutes(app, database)

    // Agregar el registro condicion de pagos
    condicionRepo := db.NewCondicionPagoRepo(database)
    api.RegisterCondicionPagoRoutes(app, condicionRepo)

    // Agregar el registro de productos
    productoRepo := db.NewProductoRepo(database)
    productoHandler := api.NewProductoHandler(productoRepo)
    productoHandler.RegisterRoutes(app)

    //  Facturación
    api.RegisterFacturaRoutes(app, database)
    api.RegisterDocContableRoutes(app, database)

    //  Timbrado
    api.RegisterTimbradoRoutes(app, database)

    //Impuestos
    impuestoRepo := db.NewImpuestoRepo(database)
    impuestoHandler := api.NewImpuestosHandler(impuestoRepo)
    impuestoHandler.RegisterRoutes(app)
    //Categoria
    categoriaRepo := db.NewCategoriaRepo(database)
    categoriaHandler := api.NewCategoriaHandler(categoriaRepo)
    categoriaHandler.RegisterRoutes(app)

        // Marcas
    marcaRepo := db.NewMarcaRepo(database)
    api.RegisterMarcaRoutes(app, marcaRepo)

        // Historial de órdenes
    ordenesRepo := db.NewOrdenesRepo(database)
    ordenesHandler := api.NewOrdenesHandler(ordenesRepo)
    ordenesHandler.RegisterRoutes(app)    

    // Historial de órdenes
    historialRepo := db.NewHistorialRepo(database)
    historialHandler := api.NewHistorialHandler(historialRepo)
    historialHandler.RegisterRoutes(app)

    return app
}

// package router

// import (
//     "backend/internal/api"
//     "backend/internal/db"
//     "database/sql"
//     "github.com/gofiber/fiber/v2"
//     "github.com/gofiber/fiber/v2/middleware/cors"
// )

// func SetupFiberRoutes(database *sql.DB) *fiber.App {
//     app := fiber.New()

//     app.Use(cors.New(cors.Config{
//         AllowOrigins: "*",
//         AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
//         AllowHeaders: "Content-Type",
//     }))

//     // Contactos
//     contactRepo := db.NewContactRepo(database)
//     contactHandler := api.NewContactHandler(contactRepo)
//     contactHandler.RegisterRoutes(app)

//     //  Agregar el registro de órdenes
//     api.RegisterOrderRoutes(app, database)

//     // Agregar el registro condicion de pagos
//     condicionRepo := db.NewCondicionPagoRepo(database)
//     api.RegisterCondicionPagoRoutes(app, condicionRepo)

//     // Agregar el registro de productos
//     productoRepo := db.NewProductoRepo(database)
//     productoHandler := api.NewProductoHandler(productoRepo)
//     productoHandler.RegisterRoutes(app)

//     //  Facturación
//     api.RegisterFacturaRoutes(app, database)
//     api.RegisterDocContableRoutes(app, database)

//     //  Timbrado
//     api.RegisterTimbradoRoutes(app, database)

//     //Impuestos
//     impuestoRepo := db.NewImpuestoRepo(database)
//     impuestoHandler := api.NewImpuestosHandler(impuestoRepo)
//     impuestoHandler.RegisterRoutes(app)
//     //Categoria
//     categoriaRepo := db.NewCategoriaRepo(database)
//     categoriaHandler := api.NewCategoriaHandler(categoriaRepo)
//     categoriaHandler.RegisterRoutes(app)

//         // Marcas
//     marcaRepo := db.NewMarcaRepo(database)
//     api.RegisterMarcaRoutes(app, marcaRepo)

//         // Historial de órdenes
//     ordenesRepo := db.NewOrdenesRepo(database)
//     ordenesHandler := api.NewOrdenesHandler(ordenesRepo)
//     ordenesHandler.RegisterRoutes(app)    

//     return app
// }
