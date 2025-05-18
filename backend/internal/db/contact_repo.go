package db

import (
    "database/sql"
    "errors"
)

// Errores personalizados
var (
    ErrContactNotFound = errors.New("contacto no encontrado")
    ErrInvalidData     = errors.New("datos de contacto no válidos")
)

type ContactoCLI struct {
    IDContacto   int    `json:"id_contacto"`
    RazonSocial  string `json:"razon_social"`
    RUCCI        string `json:"ruc_ci"`
    TelefonoCTO  string `json:"telefono_cto"`
    DireccionCTO string `json:"direccion_cto"`
}

type ContactRepo interface {
    Insert(contacto ContactoCLI) (int, error)
    Update(contacto ContactoCLI) error
    Delete(id int) error
    GetByID(id int) (ContactoCLI, error)
    GetAll() ([]ContactoCLI, error)
    Search(razonSocial string) ([]ContactoCLI, error) // Método renombrado a Search
}

type contactRepo struct {
    db *sql.DB
}

func NewContactRepo(db *sql.DB) ContactRepo {
    return &contactRepo{db: db}
}

func (repo *contactRepo) Insert(contacto ContactoCLI) (int, error) {
    if contacto.RazonSocial == "" || contacto.RUCCI == "" {
        return 0, ErrInvalidData
    }

    query := `INSERT INTO contacto_cli (razon_social, ruc_ci, telefono_cto, direccion_cto)
              VALUES ($1, $2, $3, $4) RETURNING id_contacto`
    var id int
    err := repo.db.QueryRow(query, contacto.RazonSocial, contacto.RUCCI, contacto.TelefonoCTO, contacto.DireccionCTO).Scan(&id)
    return id, err
}

func (repo *contactRepo) Update(contacto ContactoCLI) error {
    if contacto.IDContacto == 0 || contacto.RazonSocial == "" || contacto.RUCCI == "" {
        return ErrInvalidData
    }

    query := `UPDATE contacto_cli
              SET razon_social = $1, ruc_ci = $2, telefono_cto = $3, direccion_cto = $4
              WHERE id_contacto = $5`
    result, err := repo.db.Exec(query, contacto.RazonSocial, contacto.RUCCI, contacto.TelefonoCTO, contacto.DireccionCTO, contacto.IDContacto)
    if err != nil {
        return err
    }
    rowsAffected, _ := result.RowsAffected()
    if rowsAffected == 0 {
        return ErrContactNotFound
    }
    return nil
}

func (repo *contactRepo) Delete(id int) error {
    query := `DELETE FROM contacto_cli WHERE id_contacto = $1`
    result, err := repo.db.Exec(query, id)
    if err != nil {
        return err
    }
    rowsAffected, _ := result.RowsAffected()
    if rowsAffected == 0 {
        return ErrContactNotFound
    }
    return nil
}

func (repo *contactRepo) GetByID(id int) (ContactoCLI, error) {
    query := `SELECT id_contacto, razon_social, ruc_ci, telefono_cto, direccion_cto
              FROM  contacto_cli WHERE id_contacto = $1`
    var contacto ContactoCLI
    err := repo.db.QueryRow(query, id).Scan(&contacto.IDContacto, &contacto.RazonSocial, &contacto.RUCCI, &contacto.TelefonoCTO, &contacto.DireccionCTO)
    if err == sql.ErrNoRows {
        return contacto, ErrContactNotFound
    }
    return contacto, err
}

func (repo *contactRepo) GetAll() ([]ContactoCLI, error) {
    query := `SELECT id_contacto, razon_social, ruc_ci, telefono_cto, direccion_cto
              FROM contacto_cli`
    rows, err := repo.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var contactos []ContactoCLI
    for rows.Next() {
        var contacto ContactoCLI
        if err := rows.Scan(&contacto.IDContacto, &contacto.RazonSocial, &contacto.RUCCI, &contacto.TelefonoCTO, &contacto.DireccionCTO); err != nil {
            return nil, err
        }
        contactos = append(contactos, contacto)
    }
    return contactos, nil
}

func (repo *contactRepo) Search(termino string) ([]ContactoCLI, error) {
    query := `SELECT id_contacto, razon_social, ruc_ci, telefono_cto, direccion_cto
              FROM contacto_cli
              WHERE LOWER(razon_social) LIKE LOWER($1)
              OR LOWER(ruc_ci) LIKE LOWER($1)`
    rows, err := repo.db.Query(query, "%"+termino+"%")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var contactos []ContactoCLI
    for rows.Next() {
        var contacto ContactoCLI
        if err := rows.Scan(&contacto.IDContacto, &contacto.RazonSocial, &contacto.RUCCI, &contacto.TelefonoCTO, &contacto.DireccionCTO); err != nil {
            return nil, err
        }
        contactos = append(contactos, contacto)
    }
    return contactos, nil
}
