package db

import (
	"database/sql"
	"fmt"
	"strconv"
)

type Producto struct {
	IDItem              int64   `json:"id_item"`
	CodItem             string  `json:"cod_item"`
	DescripcionItem     string  `json:"descripcion_item"`
	PrecioUnitario      float64 `json:"precio_unitario"`
	ItemActivo          bool    `json:"item_activo"`
	IDMarcas            int64   `json:"id_marcas"`
	DescripcionMarcas   string  `json:"descripcion_marcas"`   // ← Nuevo campo
	IDImpuesto          int64   `json:"id_impuesto"`
	IDCategoria         int64   `json:"id_categoria"`
	DescribeCategoria   string  `json:"describe_categoria"`   // ← Nuevo campo
	CantidadInicial     int     `json:"cantidad_inicial"`
	PorcentajeIVA       float64 `json:"porcentaje_iva"`
}

type ProductoRepo struct {
	DB *sql.DB
}

func NewProductoRepo(db *sql.DB) *ProductoRepo {
	return &ProductoRepo{DB: db}
}

func (r *ProductoRepo) Create(p Producto) error {
	query := `INSERT INTO item (cod_item, descripcion_item, precio_unitario, item_activo, id_marcas, id_impuesto, id_categoria, cantidad_inicial)
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := r.DB.Exec(query, p.CodItem, p.DescripcionItem, p.PrecioUnitario, p.ItemActivo, p.IDMarcas, p.IDImpuesto, p.IDCategoria, p.CantidadInicial)
	return err
}

func (r *ProductoRepo) GetNextCodItem() (string, error) {
	var ultimoCodigo string
	err := r.DB.QueryRow(`SELECT cod_item FROM item ORDER BY id_item DESC LIMIT 1`).Scan(&ultimoCodigo)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}

	next := "001" // valor por defecto
	if ultimoCodigo != "" {
		num, convErr := strconv.Atoi(ultimoCodigo)
		if convErr == nil {
			next = fmt.Sprintf("%03d", num+1)
		}
	}

	return next, nil
}


func (r *ProductoRepo) GetAll() ([]Producto, error) {
	rows, err := r.DB.Query(`
		SELECT
			i.id_item,
			i.cod_item,
			i.descripcion_item,
			i.precio_unitario,
			i.item_activo,
			im.id_marcas,
			im.descripcion_marcas,
			i.id_impuesto,
			ic.id_categoria,
			ic.describe_categoria,
			i.cantidad_inicial,
			ii.porcentaje_iva
		FROM
			item i
		JOIN item_impuesto ii ON i.id_impuesto = ii.id_impuesto
		JOIN item_marcas im ON i.id_marcas = im.id_marcas
		JOIN item_categoria ic ON i.id_categoria = ic.id_categoria
		ORDER BY i.id_item
	`)
	if err != nil {
		return nil, fmt.Errorf("error consultando productos: %w", err)
	}
	defer rows.Close()

	var productos []Producto
	for rows.Next() {
		var p Producto
		err := rows.Scan(
			&p.IDItem, &p.CodItem, &p.DescripcionItem, &p.PrecioUnitario,
			&p.ItemActivo, &p.IDMarcas, &p.DescripcionMarcas, &p.IDImpuesto,
			&p.IDCategoria, &p.DescribeCategoria, &p.CantidadInicial, &p.PorcentajeIVA,
		)
		if err != nil {
			return nil, fmt.Errorf("error escaneando producto: %w", err)
		}
		productos = append(productos, p)
	}
	return productos, nil
}

func (r *ProductoRepo) GetByID(id int64) (*Producto, error) {
	var p Producto
	err := r.DB.QueryRow(`
		SELECT
			i.id_item,
			i.cod_item,
			i.descripcion_item,
			i.precio_unitario,
			i.item_activo,
			im.id_marcas,
			im.descripcion_marcas,
			i.id_impuesto,
			ic.id_categoria,
			ic.describe_categoria,
			i.cantidad_inicial,
			ii.porcentaje_iva
		FROM
			item i
		JOIN item_impuesto ii ON i.id_impuesto = ii.id_impuesto
		JOIN item_marcas im ON i.id_marcas = im.id_marcas
		JOIN item_categoria ic ON i.id_categoria = ic.id_categoria
		WHERE i.id_item = $1
	`, id).Scan(
		&p.IDItem, &p.CodItem, &p.DescripcionItem, &p.PrecioUnitario,
		&p.ItemActivo, &p.IDMarcas, &p.DescripcionMarcas, &p.IDImpuesto,
		&p.IDCategoria, &p.DescribeCategoria, &p.CantidadInicial, &p.PorcentajeIVA,
	)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *ProductoRepo) Update(id int64, p Producto) error {
	query := `UPDATE item SET cod_item=$1, descripcion_item=$2, precio_unitario=$3, item_activo=$4, id_marcas=$5, id_impuesto=$6, id_categoria=$7, cantidad_inicial=$8 WHERE id_item=$9`
	_, err := r.DB.Exec(query, p.CodItem, p.DescripcionItem, p.PrecioUnitario, p.ItemActivo, p.IDMarcas, p.IDImpuesto, p.IDCategoria, p.CantidadInicial, id)
	return err
}

func (r *ProductoRepo) Delete(id int64) error {
	_, err := r.DB.Exec("DELETE FROM item WHERE id_item = $1", id)
	return err
}

func (r *ProductoRepo) Search(query string) ([]Producto, error) {
	rows, err := r.DB.Query(`
		SELECT
			i.id_item,
			i.cod_item,
			i.descripcion_item,
			i.precio_unitario,
			i.item_activo,
			im.id_marcas,
			im.descripcion_marcas,
			i.id_impuesto,
			ic.id_categoria,
			ic.describe_categoria,
			i.cantidad_inicial,
			ii.porcentaje_iva
		FROM
			item i
		JOIN item_impuesto ii ON i.id_impuesto = ii.id_impuesto
		JOIN item_marcas im ON i.id_marcas = im.id_marcas
		JOIN item_categoria ic ON i.id_categoria = ic.id_categoria
		WHERE LOWER(i.cod_item) LIKE LOWER($1)
		   OR LOWER(i.descripcion_item) LIKE LOWER($1)
	`, "%"+query+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var productos []Producto
	for rows.Next() {
		var p Producto
		err := rows.Scan(
			&p.IDItem, &p.CodItem, &p.DescripcionItem, &p.PrecioUnitario,
			&p.ItemActivo, &p.IDMarcas, &p.DescripcionMarcas, &p.IDImpuesto,
			&p.IDCategoria, &p.DescribeCategoria, &p.CantidadInicial, &p.PorcentajeIVA,
		)
		if err != nil {
			return nil, err
		}
		productos = append(productos, p)
	}
	return productos, nil
}




// // backend/internal/db/producto_repo.go
// package db

// import (
// 	"database/sql"
// )

// type Producto struct {
// 	IDItem          int64    `json:"id_item"`
// 	CodItem         string   `json:"cod_item"`
// 	DescripcionItem string   `json:"descripcion_item"`
// 	PrecioUnitario  float64  `json:"precio_unitario"`
// 	ItemActivo      bool     `json:"item_activo"`
// 	IDMarcas        int64    `json:"id_marcas"`
// 	IDImpuesto      int64    `json:"id_impuesto"`
// 	IDCategoria     int64    `json:"id_categoria"`
// 	CantidadInicial int      `json:"cantidad_inicial"`
// 	PorcentajeIVA   float64  `json:"porcentaje_iva"` // Nuevo campo para el porcentaje de IVA
// }

// type ProductoRepo struct {
// 	DB *sql.DB
// }

// func NewProductoRepo(db *sql.DB) *ProductoRepo {
// 	return &ProductoRepo{DB: db}
// }

// func (r *ProductoRepo) Create(p Producto) error {
// 	query := `INSERT INTO item (cod_item, descripcion_item, precio_unitario, item_activo, id_marcas, id_impuesto, id_categoria, cantidad_inicial)
// 	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
// 	_, err := r.DB.Exec(query, p.CodItem, p.DescripcionItem, p.PrecioUnitario, p.ItemActivo, p.IDMarcas, p.IDImpuesto, p.IDCategoria, p.CantidadInicial)
// 	return err
// }

// func (r *ProductoRepo) GetAll() ([]Producto, error) {
// 	rows, err := r.DB.Query("SELECT id_item, cod_item, descripcion_item, precio_unitario, item_activo, id_marcas, id_impuesto, id_categoria, cantidad_inicial FROM item")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var productos []Producto
// 	for rows.Next() {
// 		var p Producto
// 		err := rows.Scan(&p.IDItem, &p.CodItem, &p.DescripcionItem, &p.PrecioUnitario, &p.ItemActivo, &p.IDMarcas, &p.IDImpuesto, &p.IDCategoria, &p.CantidadInicial)
// 		if err != nil {
// 			return nil, err
// 		}
// 		productos = append(productos, p)
// 	}
// 	return productos, nil
// }

// func (r *ProductoRepo) GetByID(id int64) (*Producto, error) {
// 	var p Producto
// 	err := r.DB.QueryRow("SELECT id_item, cod_item, descripcion_item, precio_unitario, item_activo, id_marcas, id_impuesto, id_categoria, cantidad_inicial FROM item WHERE id_item = $1", id).
// 		Scan(&p.IDItem, &p.CodItem, &p.DescripcionItem, &p.PrecioUnitario, &p.ItemActivo, &p.IDMarcas, &p.IDImpuesto, &p.IDCategoria, &p.CantidadInicial)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &p, nil
// }

// func (r *ProductoRepo) Update(id int64, p Producto) error {
// 	query := `UPDATE item SET cod_item=$1, descripcion_item=$2, precio_unitario=$3, item_activo=$4, id_marcas=$5, id_impuesto=$6, id_categoria=$7, cantidad_inicial=$8 WHERE id_item=$9`
// 	_, err := r.DB.Exec(query, p.CodItem, p.DescripcionItem, p.PrecioUnitario, p.ItemActivo, p.IDMarcas, p.IDImpuesto, p.IDCategoria, p.CantidadInicial, id)
// 	return err
// }

// func (r *ProductoRepo) Delete(id int64) error {
// 	_, err := r.DB.Exec("DELETE FROM item WHERE id_item = $1", id)
// 	return err
// }

// func (r *ProductoRepo) Search(query string) ([]Producto, error) {
// 	rows, err := r.DB.Query(`
// 		SELECT
// 			i.id_item,
// 			i.cod_item,
// 			i.descripcion_item,
// 			i.precio_unitario,
// 			i.item_activo,
// 			i.id_marcas,
// 			i.id_impuesto,
// 			i.id_categoria,
// 			i.cantidad_inicial,
// 			ii.porcentaje_iva
// 		FROM
// 			item i
// 		JOIN
// 			item_impuesto ii ON i.id_impuesto = ii.id_impuesto
// 		WHERE
// 			LOWER(i.cod_item) LIKE LOWER($1)
// 			OR LOWER(i.descripcion_item) LIKE LOWER($1)`, "%"+query+"%")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var productos []Producto
// 	for rows.Next() {
// 		var p Producto
// 		err := rows.Scan(&p.IDItem, &p.CodItem, &p.DescripcionItem, &p.PrecioUnitario, &p.ItemActivo, &p.IDMarcas, &p.IDImpuesto, &p.IDCategoria, &p.CantidadInicial, &p.PorcentajeIVA)
// 		if err != nil {
// 			return nil, err
// 		}
// 		productos = append(productos, p)
// 	}
// 	return productos, nil
// }
