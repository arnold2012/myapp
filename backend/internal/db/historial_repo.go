package db

import (
	"database/sql"
	"fmt"
)

type HistorialOrden struct {
	IDOrder         int64   `json:"id_order"`
	NumeroOrder     string  `json:"numero_order"`
	FechaExpedicion string  `json:"fecha_expedicion"`
	IDCondicionPago int64   `json:"idcondicion_pago"`
	IDItem          int64   `json:"id_item"`
	Cantidad        int     `json:"cantidad"`
	IDContacto      int64   `json:"id_contacto"`
	RazonSocial     string  `json:"razon_social"`
	RucCI           string  `json:"ruc_ci"`
	Condicion       string  `json:"describe_condicion"`
	CodItem         string  `json:"cod_item"`
	DescripcionItem string  `json:"descripcion_item"`
	PrecioUnitario  float64 `json:"precio_unitario"`
	PorcentajeIVA   float64 `json:"porcentaje_iva"`
	TotalItem       float64 `json:"total_item"`
	Gravado10       float64 `json:"gravado_10"`
	Gravado5        float64 `json:"gravado_5"`
	Exenta          float64 `json:"exenta"`
	IVACalculado    float64 `json:"iva_calculado"`
}

type HistorialRepo struct {
	DB *sql.DB
}

func NewHistorialRepo(db *sql.DB) *HistorialRepo {
	return &HistorialRepo{DB: db}
}

func (r *HistorialRepo) GetHistorialByOrderID(orderID int64) ([]HistorialOrden, error) {
	query := `
		SELECT
			o.id_order,
			o.numero_order,
			o.fecha_expedicion,
			o.idcondicion_pago,
			d.id_item,
			d.cantidad,
			cl.id_contacto,
			cl.razon_social,
			cl.ruc_ci,
			cp.describe_condicion,
			i.cod_item,
			i.descripcion_item,
			i.precio_unitario,
			ii.porcentaje_iva,
			(i.precio_unitario * d.cantidad) AS total_item,
			CASE WHEN ii.porcentaje_iva = 10 THEN (i.precio_unitario * d.cantidad) ELSE 0 END AS gravado_10,
			CASE WHEN ii.porcentaje_iva = 5 THEN (i.precio_unitario * d.cantidad) ELSE 0 END AS gravado_5,
			CASE WHEN ii.porcentaje_iva = 0 THEN (i.precio_unitario * d.cantidad) ELSE 0 END AS exenta,
			CASE 
				WHEN ii.porcentaje_iva = 10 THEN ROUND((i.precio_unitario * d.cantidad) / 11, 3)
				WHEN ii.porcentaje_iva = 5 THEN ROUND((i.precio_unitario * d.cantidad) / 21, 3)
				ELSE 0
			END AS iva_calculado
		FROM "order" o
		INNER JOIN order_detail d ON o.id_order = d.id_order
		INNER JOIN condicion_pagos cp ON o.idcondicion_pago = cp.idcondicion_pago
		INNER JOIN item i ON d.id_item = i.id_item
		INNER JOIN item_impuesto ii ON i.id_impuesto = ii.id_impuesto
		INNER JOIN contacto_cli cl ON o.id_contacto = cl.id_contacto
		WHERE o.id_order = $1`

	rows, err := r.DB.Query(query, orderID)
	if err != nil {
		return nil, fmt.Errorf("error consultando historial: %w", err)
	}
	defer rows.Close()

	var historial []HistorialOrden
	for rows.Next() {
		var h HistorialOrden
		err := rows.Scan(
			&h.IDOrder, &h.NumeroOrder, &h.FechaExpedicion, &h.IDCondicionPago,
			&h.IDItem, &h.Cantidad, &h.IDContacto, &h.RazonSocial, &h.RucCI,
			&h.Condicion, &h.CodItem, &h.DescripcionItem, &h.PrecioUnitario,
			&h.PorcentajeIVA, &h.TotalItem, &h.Gravado10, &h.Gravado5,
			&h.Exenta, &h.IVACalculado,
		)
		if err != nil {
			return nil, fmt.Errorf("error escaneando historial: %w", err)
		}
		historial = append(historial, h)
	}

	return historial, nil
}
