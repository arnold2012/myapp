package utils

import (
	"database/sql"
	"errors"
	"fmt"
)

// GenerateNumeroFactura genera el número de factura completo en base a los códigos
// ya almacenados ('001', '002', etc.) y al secuencial.
func GenerateNumeroFactura(db *sql.DB, idEstablecimiento, idPuntoExpedicion int64, secuencial int64) (string, error) {
	var numEst, numPunto string

	// Nota: la columna se llama numero_establecimiento (con 'm')
	err := db.QueryRow(
		`SELECT numero_establecimiento
		   FROM establecimiento
		  WHERE id_establecimiento = $1`,
		idEstablecimiento,
	).Scan(&numEst)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("establecimiento %d no encontrado", idEstablecimiento)
		}
		return "", fmt.Errorf("error obteniendo número de establecimiento: %w", err)
	}

	err = db.QueryRow(
		`SELECT numero_expedicion
		   FROM punto_expedicion
		  WHERE id_punto_expedicion = $1`,
		idPuntoExpedicion,
	).Scan(&numPunto)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("punto de expedición %d no encontrado", idPuntoExpedicion)
		}
		return "", fmt.Errorf("error obteniendo número de punto de expedición: %w", err)
	}

	// El secuencial sí lo rellenamos a 7 dígitos
	return fmt.Sprintf("%s-%s-%07d", numEst, numPunto, secuencial), nil
}

// ObtenerProximoNumeroFactura genera el próximo número de factura sin necesidad de reformatear
// establecimiento/expedición, pues ya vienen en '001', '002', etc.
func ObtenerProximoNumeroFactura(db *sql.DB, numeroEstablecimiento, numeroExpedicion string) (string, error) {
	var secuencial int64

	err := db.QueryRow(`
		SELECT COALESCE(MAX(id_factura_enca), 0) + 1
		  FROM factura_encabezado fe
		 WHERE fe.id_establecimiento = (
		       SELECT id_establecimiento
		         FROM establecimiento
		        WHERE numero_establecimiento = $1
		     )
		   AND fe.id_punto_expedicion = (
		       SELECT id_punto_expedicion
		         FROM punto_expedicion
		        WHERE numero_expedicion = $2
		     )
	`, numeroEstablecimiento, numeroExpedicion).Scan(&secuencial)
	if err != nil {
		return "", fmt.Errorf("error obteniendo secuencial de factura: %w", err)
	}

	return fmt.Sprintf("%s-%s-%07d", numeroEstablecimiento, numeroExpedicion, secuencial), nil
}


// package utils

// import (
// 	"database/sql"
// 	"fmt"
// )

// // GenerateNumeroFactura genera el número de factura completo en base a los IDs de establecimiento y punto de expedición
// func GenerateNumeroFactura(db *sql.DB, idEstablecimiento, idPuntoExpedicion int64, secuencial int64) (string, error) {
// 	var numeroEstablecimiento, numeroPunto string

// 	// Obtener número de establecimiento
// 	err := db.QueryRow(`SELECT numero_establecimiento FROM establecimiento WHERE id_establecimiento = $1`, idEstablecimiento).Scan(&numeroEstablecimiento)
// 	if err != nil {
// 		return "", fmt.Errorf("error obteniendo número de establecimiento: %w", err)
// 	}

// 	// Obtener número de punto de expedición
// 	err = db.QueryRow(`SELECT numero_expedicion FROM punto_expedicion WHERE id_punto_expedicion = $1`, idPuntoExpedicion).Scan(&numeroPunto)
// 	if err != nil {
// 		return "", fmt.Errorf("error obteniendo número de punto de expedición: %w", err)
// 	}

// 	// Formatear el número de factura
// 	numeroFactura := fmt.Sprintf("%03s-%03s-%07d", numeroEstablecimiento, numeroPunto, secuencial)
// 	return numeroFactura, nil
// }
