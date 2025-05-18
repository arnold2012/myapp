package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"backend/internal/config"
	"backend/internal/db"
	"backend/internal/router"
)

func main() {
	// 1. Carga configuración
	cfg := config.LoadConfig()

	// 2. Conecta a PostgreSQL
	postgresDB, err := db.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("Error conectando a Postgres: %v", err)
	}
	defer func() {
		if err := postgresDB.Close(); err != nil {
			log.Printf("Error cerrando conexión a DB: %v", err)
		}
	}()

	// 3. Configura rutas con Fiber
	app := router.SetupFiberRoutes(postgresDB)

	// 4. Configurar shutdown graceful
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// 5. Verificar disponibilidad del puerto
	port := ":8080"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Puerto %s en uso: %v", port, err)
	}
	listener.Close()

	// 6. Arrancar servidor en goroutine para manejar shutdown
	serverShutdown := make(chan error, 1)
	go func() {
		log.Printf("Servidor corriendo en http://localhost%s", port)
		if err := app.Listen(port); err != nil {
			serverShutdown <- err
		}
	}()

	// Esperar señal de shutdown
	select {
	case <-ctx.Done():
		log.Println("Recibida señal de apagado, cerrando servidor...")
	case err := <-serverShutdown:
		log.Printf("Error en el servidor: %v", err)
	}

	// Apagado controlado
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := app.ShutdownWithContext(shutdownCtx); err != nil {
		log.Printf("Error durante el apagado: %v", err)
	} else {
		log.Println("Servidor apagado correctamente")
	}
}