package db

import (
    "database/sql"
    "log"
)

// PurgeStaleOrders marcará como 'CADUCADA' todas las órdenes PENDIENTE
// cuya fecha_expedicion sea mayor a 1 minuto atrás.
func PurgeStaleOrders(db *sql.DB) {
    res, err := db.Exec(`
        UPDATE public."order"
        SET estado_order = 'CADUCADA'
        WHERE estado_order = 'PENDIENTE'
          AND fecha_expedicion < now() - INTERVAL '1 minute'
    `)
    if err != nil {
        log.Printf("Error purgando órdenes antiguas: %v", err)
        return
    }
    n, _ := res.RowsAffected()
    log.Printf("Órdenes caducadas: %d", n)
}
