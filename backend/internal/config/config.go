package config

import (
    "os"
)

// Config representa la configuración de la aplicación
type Config struct {
    DBHost     string
    DBPort     string
    DBUser     string
    DBName     string
    DBPassword string
}

// LoadConfig carga la configuración desde las variables de entorno
func LoadConfig() *Config {
    // Define las variables de entorno directamente en el código
    os.Setenv("DB_HOST", "127.0.0.1")
    os.Setenv("DB_PORT", "5432")
    os.Setenv("DB_USER", "postgres")
    os.Setenv("DB_NAME", "datafake")
    os.Setenv("DB_PASSWORD", "123456")

    return &Config{
        DBHost:     os.Getenv("DB_HOST"),
        DBPort:     os.Getenv("DB_PORT"),
        DBUser:     os.Getenv("DB_USER"),
        DBName:     os.Getenv("DB_NAME"),
        DBPassword: os.Getenv("DB_PASSWORD"),
    }
}
