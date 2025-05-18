package db

import (
  "database/sql"
  "fmt"
)

type PuntoExpedicion struct {
  IDPuntoExpedicion int64  `json:"id_punto_expedicion"`
  NumeroExpedicion  string `json:"numero_expedicion"`
  //Descripcion       string `json:"descripcion_puntoex"`
}

type PuntoRepo struct {
  db *sql.DB
}

func NewPuntoRepo(db *sql.DB) *PuntoRepo {
  return &PuntoRepo{db: db}
}

func (r *PuntoRepo) GetAll() ([]PuntoExpedicion, error) {
  rows, err := r.db.Query(`
    SELECT id_punto_expedicion, numero_expedicion
      FROM punto_expedicion
  `)
  if err != nil {
    return nil, fmt.Errorf("GetAll puntos_expedicion: %w", err)
  }
  defer rows.Close()

  var list []PuntoExpedicion
  for rows.Next() {
    var p PuntoExpedicion
    if err := rows.Scan(&p.IDPuntoExpedicion, &p.NumeroExpedicion); err != nil {
      return nil, err
    }
    list = append(list, p)
  }
  return list, nil
}
