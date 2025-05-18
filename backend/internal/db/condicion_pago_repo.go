package db

import (
    "database/sql"
    "fmt"
)

type CondicionPago struct {
    ID                  int64  `json:"idcondicion_pago"`
    DescribeCondicion   string `json:"describe_condicion"`
}

type CondicionPagoRepo struct {
    DB *sql.DB
}

func NewCondicionPagoRepo(db *sql.DB) *CondicionPagoRepo {
    return &CondicionPagoRepo{DB: db}
}

func (r *CondicionPagoRepo) Create(condicion *CondicionPago) error {
    query := `INSERT INTO condicion_pagos (describe_condicion) VALUES ($1) RETURNING idcondicion_pago`
    return r.DB.QueryRow(query, condicion.DescribeCondicion).Scan(&condicion.ID)
}

func (r *CondicionPagoRepo) GetAll() ([]CondicionPago, error) {
    query := `SELECT idcondicion_pago, describe_condicion FROM condicion_pagos`
    rows, err := r.DB.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var condiciones []CondicionPago
    for rows.Next() {
        var c CondicionPago
        if err := rows.Scan(&c.ID, &c.DescribeCondicion); err != nil {
            return nil, err
        }
        condiciones = append(condiciones, c)
    }

    return condiciones, nil
}

func (r *CondicionPagoRepo) GetByID(id int64) (*CondicionPago, error) {
    query := `SELECT idcondicion_pago, describe_condicion FROM condicion_pagos WHERE idcondicion_pago = $1`
    row := r.DB.QueryRow(query, id)

    var c CondicionPago
    err := row.Scan(&c.ID, &c.DescribeCondicion)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        return nil, err
    }
    return &c, nil
}

func (r *CondicionPagoRepo) Update(condicion *CondicionPago) error {
    query := `UPDATE condicion_pagos SET describe_condicion = $1 WHERE idcondicion_pago = $2`
    result, err := r.DB.Exec(query, condicion.DescribeCondicion, condicion.ID)
    if err != nil {
        return err
    }
    rowsAffected, _ := result.RowsAffected()
    if rowsAffected == 0 {
        return fmt.Errorf("no se encontró condición con ID %d", condicion.ID)
    }
    return nil
}

func (r *CondicionPagoRepo) Delete(id int64) error {
    query := `DELETE FROM condicion_pagos WHERE idcondicion_pago = $1`
    result, err := r.DB.Exec(query, id)
    if err != nil {
        return err
    }
    rowsAffected, _ := result.RowsAffected()
    if rowsAffected == 0 {
        return fmt.Errorf("no se encontró condición con ID %d", id)
    }
    return nil
}

