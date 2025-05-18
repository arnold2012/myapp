package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Timbrado struct {
	IDTimbrado        int64  `json:"id_timbrado"`
	NumeroAutorizacion string `json:"numero_autorizacion"`
}

// ObtenerTimbradosVigentes recupera los timbrados activos (estado_timbrado = 'VIGENTE')
func ObtenerTimbradosVigentes(ctx context.Context, db *sql.DB) ([]Timbrado, error) {
	rows, err := db.QueryContext(ctx, `
		SELECT id_timbrado, numero_autorizacion
		FROM timbrado
		WHERE estado_timbrado = TRUE
		ORDER BY fecha_inicio_vigencia DESC
	`)
	if err != nil {
		return nil, fmt.Errorf("error consultando timbrados vigentes: %w", err)
	}
	defer rows.Close()

	var lista []Timbrado
	for rows.Next() {
		var t Timbrado
		if err := rows.Scan(&t.IDTimbrado, &t.NumeroAutorizacion); err != nil {
			return nil, fmt.Errorf("error escaneando timbrado: %w", err)
		}
		lista = append(lista, t)
	}
	return lista, nil
}

// ObtenerTimbradoPorEstablecimientoYExpedicion retorna el id_timbrado activo segun establecimiento y expedición
func ObtenerTimbradoPorEstablecimientoYExpedicion(ctx context.Context, db *sql.DB, idEst, idExp int64) (int64, error) {
	var idTimbrado int64
	err := db.QueryRowContext(ctx, `
		SELECT id_timbrado
		FROM timbrado
		WHERE estado_timbrado = TRUE
		  AND id_establecimiento = $1
		  AND id_punto_expedicion = $2
		ORDER BY fecha_inicio_vigencia DESC
		LIMIT 1
	`, idEst, idExp).Scan(&idTimbrado)

	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("no se encontró timbrado vigente para ese punto de expedición")
		}
		return 0, fmt.Errorf("error consultando timbrado: %w", err)
	}

	return idTimbrado, nil
}
