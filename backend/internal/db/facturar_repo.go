package db

import (
	"context"
	"database/sql"
	"fmt"

	"backend/internal/utils"
)

type FacturaItem struct {
	IDItem    int64   `json:"id_item"`
	Cantidad  float64 `json:"cantidad"`
	UnitPrice float64 `json:"unit_price"`
}

type FacturaParams struct {
	IDContacto        int64         `json:"id_contacto"`
	IDCondicionPago   int64         `json:"idcondicion_pago"`
	IDPuntoExpedicion int64         `json:"id_punto_expedicion"`
	IDTimbrado        int64         `json:"id_timbrado"`
	IDEstablecimiento int64         `json:"id_establecimiento"`
	IDOrder           int64         `json:"id_order"`
	Items             []FacturaItem `json:"items"`
}

type OrderDetalle struct {
	IDOrder         int64   `json:"id_order"`
	IDContacto      int64   `json:"id_contacto"`
	IDCondicionPago int64   `json:"idcondicion_pago"`
	IDItem          int64   `json:"id_item"`
	NumeroOrder     int64   `json:"numero_order"`
	Fecha           string  `json:"fecha"`
	RazonSocial     string  `json:"razon_social"`
	RucCi           string  `json:"ruc_ci"`
	CodigoItem      string  `json:"codigo_item"`
	DescripcionItem string  `json:"descripcion_item"`
	Cantidad        float64 `json:"cantidad"`
	PrecioUnitario  float64 `json:"precio_unitario"`
	Subtotal        float64 `json:"subtotal"`
	PorcentajeIVA   float64 `json:"porcentaje_iva"`
}

func InsertFactura(ctx context.Context, db *sql.DB, params FacturaParams) error {
	var nextID int64
	err := db.QueryRow(`SELECT COALESCE(MAX(id_factura_enca), 0) + 1 FROM factura_encabezado`).Scan(&nextID)
	if err != nil {
		return fmt.Errorf("error obteniendo siguiente ID de factura: %w", err)
	}

	numeroFactura, err := utils.GenerateNumeroFactura(db, params.IDEstablecimiento, params.IDPuntoExpedicion, nextID)
	if err != nil {
		return fmt.Errorf("error generando número de factura: %w", err)
	}

	_, err = db.ExecContext(ctx, `
		INSERT INTO factura_encabezado (
			id_factura_enca, numero_factura, fecha_emision,
			id_contacto, idcondicion_pago, id_punto_expedicion,
			id_timbrado, id_establecimiento, id_order
		) VALUES ($1, $2, CURRENT_DATE, $3, $4, $5, $6, $7, $8)
	`, nextID, numeroFactura, params.IDContacto, params.IDCondicionPago,
		params.IDPuntoExpedicion, params.IDTimbrado, params.IDEstablecimiento, params.IDOrder)

	if err != nil {
		return fmt.Errorf("error insertando factura_encabezado: %w", err)
	}

	for _, item := range params.Items {
		_, err := db.ExecContext(ctx, `
			INSERT INTO factura_detalle (
				id_factura_detail, id_factura_enca,
				cantidad, unit_price, id_item
			) VALUES (DEFAULT, $1, $2, $3, $4)
		`, nextID, item.Cantidad, item.UnitPrice, item.IDItem)
		if err != nil {
			return fmt.Errorf("error insertando factura_detalle: %w", err)
		}
	}

	return nil
}

func SearchOrderByNumero(ctx context.Context, db *sql.DB, numeroOrder int64) ([]OrderDetalle, error) {
	rows, err := db.QueryContext(ctx, `
		SELECT 
			o.id_order,
			o.numero_order,
			TO_CHAR(o.fecha_expedicion, 'DD/MM/YYYY') AS fecha,
			c.id_contacto,
			o.idcondicion_pago,
			i.id_item,
			c.razon_social,
			c.ruc_ci,
			i.cod_item,
			i.descripcion_item,
			od.cantidad,
			i.precio_unitario,
			(od.cantidad * i.precio_unitario) AS subtotal,
			ii.porcentaje_iva
		FROM public."order" o
		INNER JOIN contacto_cli c ON o.id_contacto = c.id_contacto
		INNER JOIN order_detail od ON o.id_order = od.id_order
		INNER JOIN item i ON od.id_item = i.id_item
		INNER JOIN item_impuesto ii ON i.id_impuesto = ii.id_impuesto
		WHERE o.numero_order = $1
		ORDER BY o.fecha_expedicion DESC, i.cod_item
	`, numeroOrder)
	if err != nil {
		return nil, fmt.Errorf("error ejecutando consulta: %w", err)
	}
	defer rows.Close()

	var detalles []OrderDetalle
	for rows.Next() {
		var d OrderDetalle
		if err := rows.Scan(
			&d.IDOrder,
			&d.NumeroOrder,
			&d.Fecha,
			&d.IDContacto,
			&d.IDCondicionPago,
			&d.IDItem,
			&d.RazonSocial,
			&d.RucCi,
			&d.CodigoItem,
			&d.DescripcionItem,
			&d.Cantidad,
			&d.PrecioUnitario,
			&d.Subtotal,
			&d.PorcentajeIVA,
		); err != nil {
			return nil, fmt.Errorf("error escaneando fila: %w", err)
		}
		detalles = append(detalles, d)
	}

	return detalles, nil
}


// package db

// import (
// 	"context"
// 	"database/sql"
// 	"fmt"

// 	"backend/internal/utils"
// )

// type FacturaItem struct {
// 	IDItem    int64   `json:"id_item"`
// 	Cantidad  float64 `json:"cantidad"`
// 	UnitPrice float64 `json:"unit_price"`
// }

// type FacturaParams struct {
// 	IDContacto        int64          `json:"id_contacto"`
// 	IDCondicionPago   int64          `json:"idcondicion_pago"`
// 	IDPuntoExpedicion int64          `json:"id_punto_expedicion"`
// 	IDTimbrado        int64          `json:"id_timbrado"`
// 	IDEstablecimiento int64          `json:"id_establecimiento"`
// 	IDOrder           int64          `json:"id_order"`
// 	Items             []FacturaItem  `json:"items"`
// }

// func InsertFactura(ctx context.Context, db *sql.DB, params FacturaParams) error {
// 	// Obtener el siguiente ID de factura
// 	var nextID int64
// 	err := db.QueryRow(`SELECT COALESCE(MAX(id_factura_enca), 0) + 1 FROM factura_encabezado`).Scan(&nextID)
// 	if err != nil {
// 		return fmt.Errorf("error obteniendo siguiente ID de factura: %w", err)
// 	}

// 	// Generar número de factura
// 	numeroFactura, err := utils.GenerateNumeroFactura(db, params.IDEstablecimiento, params.IDPuntoExpedicion, nextID)
// 	if err != nil {
// 		return fmt.Errorf("error generando número de factura: %w", err)
// 	}

// 	// Insertar en factura_encabezado
// 	_, err = db.ExecContext(ctx, `
// 		INSERT INTO factura_encabezado (
// 			id_factura_enca, numero_factura, fecha_emision,
// 			id_contacto, idcondicion_pago, id_punto_expedicion,
// 			id_timbrado, id_establecimiento, id_order
// 		) VALUES ($1, $2, CURRENT_DATE, $3, $4, $5, $6, $7, $8)
// 	`, nextID, numeroFactura, params.IDContacto, params.IDCondicionPago,
// 		params.IDPuntoExpedicion, params.IDTimbrado, params.IDEstablecimiento, params.IDOrder)

// 	if err != nil {
// 		return fmt.Errorf("error insertando factura_encabezado: %w", err)
// 	}

// 	// Insertar detalles
// 	for _, item := range params.Items {
// 		_, err := db.ExecContext(ctx, `
// 			INSERT INTO factura_detalle (
// 				id_factura_detail, id_factura_enca,
// 				cantidad, unit_price, id_item
// 			) VALUES (DEFAULT, $1, $2, $3, $4)
// 		`, nextID, item.Cantidad, item.UnitPrice, item.IDItem)
// 		if err != nil {
// 			return fmt.Errorf("error insertando factura_detalle: %w", err)
// 		}
// 	}

// 	return nil
// }
