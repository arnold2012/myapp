#!/bin/bash

# Función para ejecutar el backend
run_backend() {
  echo "Iniciando el servidor backend..."
  cd backend
  go run ./cmd/server &
  cd ..
}

# Función para ejecutar el frontend
run_frontend() {
  echo "Iniciando el servidor frontend..."
  cd frontend
  npm run dev &
  cd ..
}

# Ejecutar ambos servidores
run_backend
run_frontend

# Esperar a que ambos procesos terminen
wait
