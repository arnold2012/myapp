
package db

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"
	"time"
)

// Timbrado representa la estructura de un timbrado en la base de datos
type Timbrado struct {
	IDTimbrado           int64     `json:"id_timbrado"`
	NumeroAutorizacion   string    `json:"numero_autorizacion"`
	FechaAutorizacion    time.Time `json:"fecha_autorizacion"`
	FechaInicioVigencia  time.Time `json:"fecha_inicio_vigencia"`
	EstadoTimbrado       bool      `json:"estado_timbrado"`
	IDEstablecimiento    int64     `json:"id_establecimiento"`
	IDPuntoExpedicion    int64     `json:"id_punto_expedicion"`
}

// ObtenerTimbradosVigentes recupera los timbrados activos con todos sus campos
func ObtenerTimbradosVigentes(ctx context.Context, db *sql.DB) ([]Timbrado, error) {
	rows, err := db.QueryContext(ctx, `
		SELECT 
			t.id_timbrado, 
			t.numero_autorizacion, 
			t.fecha_autorizacion, 
			t.fecha_inicio_vigencia, 
			t.estado_timbrado, 
			t.id_establecimiento, 
			t.id_punto_expedicion
		FROM timbrado t
		WHERE t.estado_timbrado = TRUE
		ORDER BY t.fecha_inicio_vigencia DESC
	`)
	if err != nil {
		return nil, fmt.Errorf("error consultando timbrados vigentes: %w", err)
	}
	defer rows.Close()

	var lista []Timbrado
	for rows.Next() {
		var t Timbrado
		if err := rows.Scan(
			&t.IDTimbrado,
			&t.NumeroAutorizacion,
			&t.FechaAutorizacion,
			&t.FechaInicioVigencia,
			&t.EstadoTimbrado,
			&t.IDEstablecimiento,
			&t.IDPuntoExpedicion,
		); err != nil {
			return nil, fmt.Errorf("error escaneando timbrado: %w", err)
		}
		lista = append(lista, t)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterando filas: %w", err)
	}

	return lista, nil
}

// ObtenerTimbradoPorEstablecimientoYExpedicion retorna el id_timbrado activo según establecimiento y expedición
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
			return 0, fmt.Errorf("no se encontró timbrado vigente para establecimiento %d y punto de expedición %d", idEst, idExp)
		}
		return 0, fmt.Errorf("error consultando timbrado: %w", err)
	}

	return idTimbrado, nil
}

// CrearTimbrado inserta un nuevo timbrado con validaciones
func CrearTimbrado(ctx context.Context, db *sql.DB, t Timbrado) (int64, error) {
	// Validar formato de numero_autorizacion (solo números, mínimo 6 dígitos)
	if !regexp.MustCompile(`^[0-9]{6,}$`).MatchString(t.NumeroAutorizacion) {
		return 0, fmt.Errorf("el número de autorización debe contener solo números y tener al menos 6 dígitos")
	}

	// Validar fechas
	if t.FechaAutorizacion.IsZero() || t.FechaInicioVigencia.IsZero() {
		return 0, fmt.Errorf("las fechas de autorización y de inicio de vigencia son obligatorias")
	}

	// Verificar existencia de establecimiento
	var estCount int
	err := db.QueryRowContext(ctx, `
		SELECT COUNT(*) 
		FROM establecimiento 
		WHERE id_establecimiento = $1
	`, t.IDEstablecimiento).Scan(&estCount)
	if err != nil {
		return 0, fmt.Errorf("error verificando establecimiento: %w", err)
	}
	if estCount == 0 {
		return 0, fmt.Errorf("el establecimiento con ID %d no existe", t.IDEstablecimiento)
	}

	// Verificar existencia de punto de expedición
	var expCount int
	err = db.QueryRowContext(ctx, `
		SELECT COUNT(*) 
		FROM punto_expedicion 
		WHERE id_punto_expedicion = $1
	`, t.IDPuntoExpedicion).Scan(&expCount)
	if err != nil {
		return 0, fmt.Errorf("error verificando punto de expedición: %w", err)
	}
	if expCount == 0 {
		return 0, fmt.Errorf("el punto de expedición con ID %d no existe", t.IDPuntoExpedicion)
	}

	// Verificar si ya existe un timbrado vigente para el mismo establecimiento y punto de expedición
	var existingID int64
	err = db.QueryRowContext(ctx, `
		SELECT id_timbrado
		FROM timbrado
		WHERE estado_timbrado = TRUE
		  AND id_establecimiento = $1
		  AND id_punto_expedicion = $2
	`, t.IDEstablecimiento, t.IDPuntoExpedicion).Scan(&existingID)
	if err == nil {
		return 0, fmt.Errorf("ya existe un timbrado vigente con ID %d para este establecimiento y punto de expedición", existingID)
	}
	if err != sql.ErrNoRows {
		return 0, fmt.Errorf("error verificando timbrado existente: %w", err)
	}

	// Iniciar transacción
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("error iniciando transacción: %w", err)
	}

	var id int64
	err = tx.QueryRowContext(ctx, `
		INSERT INTO timbrado (
			numero_autorizacion,
			fecha_autorizacion,
			fecha_inicio_vigencia,
			estado_timbrado,
			id_establecimiento,
			id_punto_expedicion
		)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id_timbrado
	`, t.NumeroAutorizacion, t.FechaAutorizacion, t.FechaInicioVigencia,
		t.EstadoTimbrado, t.IDEstablecimiento, t.IDPuntoExpedicion).Scan(&id)

	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("error creando timbrado: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return 0, fmt.Errorf("error confirmando transacción: %w", err)
	}

	return id, nil
}

// ActualizarTimbrado actualiza un timbrado existente
func ActualizarTimbrado(ctx context.Context, db *sql.DB, t Timbrado) error {
	// Validar formato de numero_autorizacion
	if !regexp.MustCompile(`^[0-9]{6,}$`).MatchString(t.NumeroAutorizacion) {
		return fmt.Errorf("el número de autorización debe contener solo números y tener al menos 6 dígitos")
	}

	// Validar fechas
	if t.FechaAutorizacion.IsZero() || t.FechaInicioVigencia.IsZero() {
		return fmt.Errorf("las fechas de autorización y de inicio de vigencia son obligatorias")
	}

	// Verificar existencia de establecimiento
	var estCount int
	err := db.QueryRowContext(ctx, `
		SELECT COUNT(*) 
		FROM establecimiento 
		WHERE id_establecimiento = $1
	`, t.IDEstablecimiento).Scan(&estCount)
	if err != nil {
		return fmt.Errorf("error verificando establecimiento: %w", err)
	}
	if estCount == 0 {
		return fmt.Errorf("el establecimiento con ID %d no existe", t.IDEstablecimiento)
	}

	// Verificar existencia de punto de expedición
	var expCount int
	err = db.QueryRowContext(ctx, `
		SELECT COUNT(*) 
		FROM punto_expedicion 
		WHERE id_punto_expedicion = $1
	`, t.IDPuntoExpedicion).Scan(&expCount)
	if err != nil {
		return fmt.Errorf("error verificando punto de expedición: %w", err)
	}
	if expCount == 0 {
		return fmt.Errorf("el punto de expedición con ID %d no existe", t.IDPuntoExpedicion)
	}

	// Iniciar transacción
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("error iniciando transacción: %w", err)
	}

	result, err := tx.ExecContext(ctx, `
		UPDATE timbrado
		SET
			numero_autorizacion = $1,
			fecha_autorizacion = $2,
			fecha_inicio_vigencia = $3,
			estado_timbrado = $4,
			id_establecimiento = $5,
			id_punto_expedicion = $6
		WHERE id_timbrado = $7
	`, t.NumeroAutorizacion, t.FechaAutorizacion, t.FechaInicioVigencia,
		t.EstadoTimbrado, t.IDEstablecimiento, t.IDPuntoExpedicion, t.IDTimbrado)

	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error actualizando timbrado: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error verificando filas afectadas: %w", err)
	}
	if rowsAffected == 0 {
		tx.Rollback()
		return fmt.Errorf("no se encontró el timbrado con ID %d", t.IDTimbrado)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error confirmando transacción: %w", err)
	}

	return nil
}

// EliminarTimbrado hace un soft delete (cambia estado a FALSE)
func EliminarTimbrado(ctx context.Context, db *sql.DB, id int64) error {
	result, err := db.ExecContext(ctx, `
		UPDATE timbrado
		SET estado_timbrado = FALSE
		WHERE id_timbrado = $1
	`, id)
	if err != nil {
		return fmt.Errorf("error eliminando timbrado: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error verificando filas afectadas: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no se encontró el timbrado con ID %d", id)
	}

	return nil
}

/*
Recomendaciones de índices para optimizar consultas:
CREATE INDEX idx_timbrado_estado ON timbrado(estado_timbrado);
CREATE INDEX idx_timbrado_est_exp ON timbrado(id_establecimiento, id_punto_expedicion);
CREATE INDEX idx_timbrado_fecha_vigencia ON timbrado(fecha_inicio_vigencia);
*/
// package db

// import (
// 	"context"
// 	"database/sql"
// 	"fmt"
// 	"time"

// )

// type Timbrado struct {
// 	IDTimbrado           int64     `json:"id_timbrado"`
// 	NumeroAutorizacion   string    `json:"numero_autorizacion"`
// 	FechaAutorizacion    time.Time `json:"fecha_autorizacion"`
// 	FechaInicioVigencia  time.Time `json:"fecha_inicio_vigencia"`
// 	EstadoTimbrado       bool      `json:"estado_timbrado"`
// 	IDEstablecimiento    int64     `json:"id_establecimiento"`
// 	IDPuntoExpedicion    int64     `json:"id_punto_expedicion"`
// }



// // ObtenerTimbradosVigentes recupera los timbrados activos (estado_timbrado = 'VIGENTE')
// func ObtenerTimbradosVigentes(ctx context.Context, db *sql.DB) ([]Timbrado, error) {
// 	rows, err := db.QueryContext(ctx, `
// 		SELECT id_timbrado, numero_autorizacion
// 		FROM timbrado
// 		WHERE estado_timbrado = TRUE
// 		ORDER BY fecha_inicio_vigencia DESC
// 	`)
// 	if err != nil {
// 		return nil, fmt.Errorf("error consultando timbrados vigentes: %w", err)
// 	}
// 	defer rows.Close()

// 	var lista []Timbrado
// 	for rows.Next() {
// 		var t Timbrado
// 		if err := rows.Scan(&t.IDTimbrado, &t.NumeroAutorizacion); err != nil {
// 			return nil, fmt.Errorf("error escaneando timbrado: %w", err)
// 		}
// 		lista = append(lista, t)
// 	}
// 	return lista, nil
// }

// // ObtenerTimbradoPorEstablecimientoYExpedicion retorna el id_timbrado activo segun establecimiento y expedición
// func ObtenerTimbradoPorEstablecimientoYExpedicion(ctx context.Context, db *sql.DB, idEst, idExp int64) (int64, error) {
// 	var idTimbrado int64
// 	err := db.QueryRowContext(ctx, `
// 		SELECT id_timbrado
// 		FROM timbrado
// 		WHERE estado_timbrado = TRUE
// 		  AND id_establecimiento = $1
// 		  AND id_punto_expedicion = $2
// 		ORDER BY fecha_inicio_vigencia DESC
// 		LIMIT 1
// 	`, idEst, idExp).Scan(&idTimbrado)

// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return 0, fmt.Errorf("no se encontró timbrado vigente para ese punto de expedición")
// 		}
// 		return 0, fmt.Errorf("error consultando timbrado: %w", err)
// 	}

// 	return idTimbrado, nil
// }
// // CrearTimbrado inserta un nuevo timbrado
// func CrearTimbrado(ctx context.Context, db *sql.DB, t Timbrado) (int64, error) {
// 	var id int64
// 	err := db.QueryRowContext(ctx, `
// 		INSERT INTO timbrado (
// 			numero_autorizacion,
// 			fecha_autorizacion,
// 			fecha_inicio_vigencia,
// 			estado_timbrado,
// 			id_establecimiento,
// 			id_punto_expedicion
// 		)
// 		VALUES ($1, $2, $3, $4, $5, $6)
// 		RETURNING id_timbrado
// 	`, t.NumeroAutorizacion, t.FechaAutorizacion, t.FechaInicioVigencia,
// 		t.EstadoTimbrado, t.IDEstablecimiento, t.IDPuntoExpedicion).Scan(&id)

// 	if err != nil {
// 		return 0, fmt.Errorf("error creando timbrado: %w", err)
// 	}
// 	return id, nil
// }

// // ActualizarTimbrado actualiza un timbrado existente
// func ActualizarTimbrado(ctx context.Context, db *sql.DB, t Timbrado) error {
// 	_, err := db.ExecContext(ctx, `
// 		UPDATE timbrado
// 		SET
// 			numero_autorizacion = $1,
// 			fecha_autorizacion = $2,
// 			fecha_inicio_vigencia = $3,
// 			estado_timbrado = $4,
// 			id_establecimiento = $5,
// 			id_punto_expedicion = $6
// 		WHERE id_timbrado = $7
// 	`, t.NumeroAutorizacion, t.FechaAutorizacion, t.FechaInicioVigencia,
// 		t.EstadoTimbrado, t.IDEstablecimiento, t.IDPuntoExpedicion, t.IDTimbrado)

// 	if err != nil {
// 		return fmt.Errorf("error actualizando timbrado: %w", err)
// 	}
// 	return nil
// }

// // EliminarTimbrado hace un soft delete (cambia estado a FALSE)
// func EliminarTimbrado(ctx context.Context, db *sql.DB, id int64) error {
// 	_, err := db.ExecContext(ctx, `
// 		UPDATE timbrado
// 		SET estado_timbrado = FALSE
// 		WHERE id_timbrado = $1
// 	`, id)
// 	if err != nil {
// 		return fmt.Errorf("error eliminando timbrado: %w", err)
// 	}
// 	return nil
// }
