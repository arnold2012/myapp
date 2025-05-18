package db

import (
  "database/sql"
  "fmt"
)

type Establecimiento struct {
  IDEstablecimiento      int64  `json:"id_establecimiento"`
  NumeroEstableciminto   string `json:"numero_establecimiento"`
//   RazonSocial            string `json:"razon_social"`
}

type EstablecimientoRepo struct {
  db *sql.DB
}

func NewEstablecimientoRepo(db *sql.DB) *EstablecimientoRepo {
  return &EstablecimientoRepo{db: db}
}

func (r *EstablecimientoRepo) GetAll() ([]Establecimiento, error) {
  rows, err := r.db.Query(`
    SELECT id_establecimiento, numero_establecimiento
      FROM establecimiento
  `)
  if err != nil {
    return nil, fmt.Errorf("GetAll establecimientos: %w", err)
  }
  defer rows.Close()

  var list []Establecimiento
  for rows.Next() {
    var e Establecimiento
    if err := rows.Scan(&e.IDEstablecimiento, &e.NumeroEstableciminto); err != nil {
      return nil, err
    }
    list = append(list, e)
  }
  return list, nil
}
