package db

import (
	"database/sql"
	"fmt"
)
// Impuestos representa la entidad del impuesto de un ítem
type Impuestos struct {
	ID            int64   `json:"id_impuesto"`
	PorcentajeIVA float64 `json:"porcentaje_iva"`
}

// ImpuestosRepo proporciona acceso a los datos de la tabla item_impuesto
type ImpuestosRepo struct {
	DB *sql.DB
}

// NewImpuestoRepo crea una nueva instancia del repositorio
func NewImpuestoRepo(db *sql.DB) *ImpuestosRepo {
	return &ImpuestosRepo{DB: db}
}

// GetAll obtiene todos los registros de impuestos
func (r *ImpuestosRepo) GetAll() ([]Impuestos, error) {
	rows, err := r.DB.Query(`SELECT id_impuesto, porcentaje_iva FROM item_impuesto ORDER BY id_impuesto`)
	if err != nil {
		return nil, fmt.Errorf("error consultando item_impuesto: %w", err)
	}
	defer rows.Close()

	var result []Impuestos
	for rows.Next() {
		var imp Impuestos
		if err := rows.Scan(&imp.ID, &imp.PorcentajeIVA); err != nil {
			return nil, fmt.Errorf("error escaneando item_impuesto: %w", err)
		}
		result = append(result, imp)
	}

	return result, nil
}

// GetByID obtiene un impuesto por su ID
func (r *ImpuestosRepo) GetByID(id int64) (*Impuestos, error) {
	var imp Impuestos
	err := r.DB.QueryRow(`SELECT id_impuesto, porcentaje_iva FROM item_impuesto WHERE id_impuesto = $1`, id).
		Scan(&imp.ID, &imp.PorcentajeIVA)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error buscando item_impuesto: %w", err)
	}
	return &imp, nil
}

// Insert agrega un nuevo registro de impuesto
func (r *ImpuestosRepo) Insert(imp Impuestos) (int64, error) {
	var id int64
	err := r.DB.QueryRow(`INSERT INTO item_impuesto (porcentaje_iva) VALUES ($1) RETURNING id_impuesto`,
		imp.PorcentajeIVA).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("error insertando item_impuesto: %w", err)
	}
	return id, nil
}

// Update actualiza un impuesto existente
func (r *ImpuestosRepo) Update(imp Impuestos) error {
	result, err := r.DB.Exec(`UPDATE item_impuesto SET porcentaje_iva = $1 WHERE id_impuesto = $2`,
		imp.PorcentajeIVA, imp.ID)
	if err != nil {
		return fmt.Errorf("error actualizando item_impuesto: %w", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error verificando actualización: %w", err)
	}
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}

// Delete elimina un impuesto por ID
func (r *ImpuestosRepo) Delete(id int64) error {
	result, err := r.DB.Exec(`DELETE FROM item_impuesto WHERE id_impuesto = $1`, id)
	if err != nil {
		return fmt.Errorf("error eliminando item_impuesto: %w", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error verificando eliminación: %w", err)
	}
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}
