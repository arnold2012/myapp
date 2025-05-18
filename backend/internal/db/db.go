package db

import (
    "database/sql"
    "fmt"
    "log"

    "backend/internal/config"
    _ "github.com/lib/pq" // Importa el controlador de PostgreSQL
)

// NewPostgresDB crea una nueva conexión a la base de datos PostgreSQL
func NewPostgresDB(cfg *config.Config) (*sql.DB, error) {
    dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
        cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBName, cfg.DBPassword)
    log.Printf("Conectando a la base de datos con DSN: %s", dsn)
    db, err := sql.Open("postgres", dsn)
    if err != nil {
        return nil, err
    }
    if err = db.Ping(); err != nil {
        return nil, err
    }
    log.Println("Conexión a la base de datos establecida correctamente")
    return db, nil
}
