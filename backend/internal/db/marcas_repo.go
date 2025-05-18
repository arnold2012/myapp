package db

import (
	"database/sql"
	"fmt"
)

type Marca struct {
	IDMarca         int64  `json:"id_marcas"`
	DescribeMarca   string `json:"descripcion_marcas"`
}

type MarcaRepo struct {
	DB *sql.DB
}

func NewMarcaRepo(db *sql.DB) *MarcaRepo {
	return &MarcaRepo{DB: db}
}

func (r *MarcaRepo) GetAll() ([]Marca, error) {
	rows, err := r.DB.Query(`SELECT id_marcas, descripcion_marcas FROM item_marcas ORDER BY id_marcas`)
	if err != nil {
		return nil, fmt.Errorf("error consultando marcas: %w", err)
	}
	defer rows.Close()

	var result []Marca
	for rows.Next() {
		var m Marca
		if err := rows.Scan(&m.IDMarca, &m.DescribeMarca); err != nil {
			return nil, err
		}
		result = append(result, m)
	}
	return result, nil
}

func (r *MarcaRepo) GetByID(id int64) (*Marca, error) {
	var m Marca
	err := r.DB.QueryRow(`SELECT id_marcas, descripcion_marcas FROM item_marcas WHERE id_marcas = $1`, id).
		Scan(&m.IDMarca, &m.DescribeMarca)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *MarcaRepo) Create(m Marca) (int64, error) {
	var id int64
	err := r.DB.QueryRow(
		`INSERT INTO item_marcas (descripcion_marcas) VALUES ($1) RETURNING id_marcas`,
		m.DescribeMarca,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *MarcaRepo) Update(m Marca) error {
	res, err := r.DB.Exec(`UPDATE item_marcas SET descripcion_marcas = $1 WHERE id_marcas = $2`, m.DescribeMarca, m.IDMarca)
	if err != nil {
		return err
	}
	count, _ := res.RowsAffected()
	if count == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (r *MarcaRepo) Delete(id int64) error {
	res, err := r.DB.Exec(`DELETE FROM item_marcas WHERE id_marcas = $1`, id)
	if err != nil {
		return err
	}
	count, _ := res.RowsAffected()
	if count == 0 {
		return sql.ErrNoRows
	}
	return nil
}
