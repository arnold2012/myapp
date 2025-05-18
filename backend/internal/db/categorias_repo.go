package db

import (
    "database/sql"
    "errors"
    "fmt"
)

type Categoria struct {
    IDCategoria        int64  `json:"id_categoria"`
    DescribeCategoria  string `json:"describe_categoria"`
}

var (
    ErrCategoriaNotFound = errors.New("categoría no encontrada")
    ErrCategoriaInvalida = errors.New("descripción requerida")
)

type CategoriaRepo interface {
    GetAll() ([]Categoria, error)
    GetByID(id int64) (*Categoria, error)
    Create(cat Categoria) (int64, error)
    Update(cat Categoria) error
    Delete(id int64) error
}

type categoriaRepo struct {
    db *sql.DB
}

func NewCategoriaRepo(db *sql.DB) CategoriaRepo {
    return &categoriaRepo{db: db}
}

func (r *categoriaRepo) GetAll() ([]Categoria, error) {
    rows, err := r.db.Query(`SELECT id_categoria, describe_categoria FROM item_categoria ORDER BY id_categoria`)
    if err != nil {
        return nil, fmt.Errorf("error al obtener categorías: %w", err)
    }
    defer rows.Close()

    var cats []Categoria
    for rows.Next() {
        var c Categoria
        if err := rows.Scan(&c.IDCategoria, &c.DescribeCategoria); err != nil {
            return nil, err
        }
        cats = append(cats, c)
    }
    return cats, nil
}

func (r *categoriaRepo) GetByID(id int64) (*Categoria, error) {
    var c Categoria
    err := r.db.QueryRow(`SELECT id_categoria, describe_categoria FROM item_categoria WHERE id_categoria = $1`, id).
        Scan(&c.IDCategoria, &c.DescribeCategoria)
    if err == sql.ErrNoRows {
        return nil, ErrCategoriaNotFound
    } else if err != nil {
        return nil, err
    }
    return &c, nil
}

func (r *categoriaRepo) Create(c Categoria) (int64, error) {
    if c.DescribeCategoria == "" {
        return 0, ErrCategoriaInvalida
    }

    var id int64
    err := r.db.QueryRow(`INSERT INTO item_categoria (describe_categoria) VALUES ($1) RETURNING id_categoria`, c.DescribeCategoria).
        Scan(&id)
    return id, err
}

func (r *categoriaRepo) Update(c Categoria) error {
    if c.IDCategoria == 0 || c.DescribeCategoria == "" {
        return ErrCategoriaInvalida
    }

    res, err := r.db.Exec(`UPDATE item_categoria SET describe_categoria = $1 WHERE id_categoria = $2`, c.DescribeCategoria, c.IDCategoria)
    if err != nil {
        return err
    }
    rows, _ := res.RowsAffected()
    if rows == 0 {
        return ErrCategoriaNotFound
    }
    return nil
}

func (r *categoriaRepo) Delete(id int64) error {
    res, err := r.db.Exec(`DELETE FROM item_categoria WHERE id_categoria = $1`, id)
    if err != nil {
        return err
    }
    rows, _ := res.RowsAffected()
    if rows == 0 {
        return ErrCategoriaNotFound
    }
    return nil
}
