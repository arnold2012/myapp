package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"
)

type OrderItem struct {
	IDItem   int64 `json:"id_item"`
	Cantidad int   `json:"cantidad"`
}

type InsertOrderParams struct {
	IDContacto      int64       `json:"id_contacto"`
	NumeroOrder     float64     `json:"numero_order"`
	Items           []OrderItem `json:"items"`
	IDCondicionPago int64       `json:"idcondicion_pago"`
}

// InsertSalesOrder now inserts using CURRENT_TIMESTAMP and returns the new order ID
func InsertSalesOrder(ctx context.Context, db *sql.DB, params InsertOrderParams) (int64, error) {
	itemsJSON, err := json.Marshal(params.Items)
	if err != nil {
		return 0, fmt.Errorf("error serializando items: %w", err)
	}

	// Call stored procedure without fecha_expedicion parameter
	_, err = db.ExecContext(ctx, `
		CALL public.insert_sales_order($1, $2, $3::jsonb, $4)
	`,
		params.IDContacto,
		params.NumeroOrder,
		itemsJSON,
		params.IDCondicionPago,
	)
	if err != nil {
		return 0, fmt.Errorf("error ejecutando procedimiento almacenado: %w", err)
	}

	// Query the newly inserted order ID
	var orderID int64
	err = db.QueryRowContext(ctx, `
		SELECT id_order
		FROM public."order"
		WHERE numero_order = $1
		  AND id_contacto = $2
		ORDER BY id_order DESC
		LIMIT 1
	`, params.NumeroOrder, params.IDContacto).Scan(&orderID)

	if err != nil {
		return 0, fmt.Errorf("error obteniendo ID de la orden: %w", err)
	}

	// Programar la caducidad de la orden después de 1 minuto
	go func(id int64) {
		time.Sleep(1 * time.Minute)
		
		// Crear un nuevo contexto para esta operación
		expireCtx := context.Background()
		
		// Actualizar la orden si sigue en estado PENDIENTE
		db.ExecContext(expireCtx, `
			UPDATE public."order"
			SET estado_order = 'CADUCADA'
			WHERE id_order = $1
			AND estado_order = 'PENDIENTE'
		`, id)
	}(orderID)

	return orderID, nil
}

// UpdateExpiredOrders actualiza todas las órdenes pendientes que han pasado 1 minuto
func UpdateExpiredOrders(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, `
		UPDATE public."order"
		SET estado_order = 'CADUCADA'
		WHERE estado_order = 'PENDIENTE'
		AND fecha_expedicion < (CURRENT_TIMESTAMP - INTERVAL '1 minute')
	`)
	
	return err
}

// package db

// import (
// 	"context"
// 	"database/sql"
// 	"encoding/json"
// 	"fmt"
// )

// type OrderItem struct {
// 	IDItem   int64 `json:"id_item"`
// 	Cantidad int   `json:"cantidad"`
// }

// type InsertOrderParams struct {
// 	IDContacto      int64       `json:"id_contacto"`
// 	NumeroOrder     float64     `json:"numero_order"`
// 	Items           []OrderItem `json:"items"`
// 	IDCondicionPago int64       `json:"idcondicion_pago"`
// }

// // InsertSalesOrder now inserts using CURRENT_TIMESTAMP and returns the new order ID
// func InsertSalesOrder(ctx context.Context, db *sql.DB, params InsertOrderParams) (int64, error) {
// 	itemsJSON, err := json.Marshal(params.Items)
// 	if err != nil {
// 		return 0, fmt.Errorf("error serializando items: %w", err)
// 	}

// 	// Call stored procedure without fecha_expedicion parameter
// 	_, err = db.ExecContext(ctx, `
// 		CALL public.insert_sales_order($1, $2, $3::jsonb, $4)
// 	`,
// 		params.IDContacto,
// 		params.NumeroOrder,
// 		itemsJSON,
// 		params.IDCondicionPago,
// 	)
// 	if err != nil {
// 		return 0, fmt.Errorf("error ejecutando procedimiento almacenado: %w", err)
// 	}

// 	// Query the newly inserted order ID
// 	var orderID int64
// 	err = db.QueryRowContext(ctx, `
// 		SELECT id_order
// 		FROM public."order"
// 		WHERE numero_order = $1
// 		  AND id_contacto = $2
// 		ORDER BY id_order DESC
// 		LIMIT 1
// 	`, params.NumeroOrder, params.IDContacto).Scan(&orderID)

// 	if err != nil {
// 		return 0, fmt.Errorf("error obteniendo ID de la orden: %w", err)
// 	}

// 	return orderID, nil
// }
