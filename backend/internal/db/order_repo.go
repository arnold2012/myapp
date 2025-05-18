package db

import (
    "context"
    "database/sql"
    "encoding/json"
    "fmt"
)

type OrderItem struct {
    IDItem   int64 `json:"id_item"`
    Cantidad int   `json:"cantidad"`
}

type InsertOrderParams struct {
    IDContacto      int64       `json:"id_contacto"`
    NumeroOrder     float64     `json:"numero_order"`
    FechaExpedicion string      `json:"fecha_expedicion"`
    Items           []OrderItem `json:"items"`
    IDCondicionPago int64       `json:"idcondicion_pago"`
}

// InsertSalesOrder ahora devuelve el ID de la orden creada
func InsertSalesOrder(ctx context.Context, db *sql.DB, params InsertOrderParams) (int64, error) {
    itemsJSON, err := json.Marshal(params.Items)
    if err != nil {
        return 0, fmt.Errorf("error serializando items: %w", err)
    }

    // Ejecutar el procedimiento almacenado
    _, err = db.ExecContext(ctx, `
        CALL public.insert_sales_order($1, $2, $3, $4::jsonb, $5)
    `,
        params.IDContacto,
        params.NumeroOrder,
        params.FechaExpedicion,
        itemsJSON,
        params.IDCondicionPago,
    )
    if err != nil {
        return 0, fmt.Errorf("error ejecutando procedimiento almacenado: %w", err)
    }

    // Consultar el ID de la orden recién creada
    var orderID int64
    err = db.QueryRowContext(ctx, `
        SELECT id_order FROM "order" 
        WHERE numero_order = $1 AND id_contacto = $2 AND fecha_expedicion = $3
        ORDER BY id_order DESC LIMIT 1
    `, params.NumeroOrder, params.IDContacto, params.FechaExpedicion).Scan(&orderID)
    
    if err != nil {
        return 0, fmt.Errorf("error obteniendo ID de la orden: %w", err)
    }

    return orderID, nil
}



// curl -X POST http://localhost:8080/api/orders   -H "Content-Type: application/json"   -d '{
//     "id_contacto": 1,
//     "numero_order": 1001,
//     "fecha_expedicion": "2025-04-24",
//     "idcondicion_pago": 2,
//     "items": [
//       { "id_item": 1, "cantidad": 3 },
//       { "id_item": 2, "cantidad": 5 }
//     ]
//   }'