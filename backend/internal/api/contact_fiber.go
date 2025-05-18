package api

import (
    "backend/internal/db"
    "github.com/gofiber/fiber/v2"
    "strconv"
)

type ContactHandler struct {
    repo db.ContactRepo
}

func NewContactHandler(repo db.ContactRepo) *ContactHandler {
    return &ContactHandler{repo: repo}
}
func (h *ContactHandler) RegisterRoutes(app *fiber.App) {
	api := app.Group("/api/contactos")
	api.Get("/search", h.searchContacts)      // debe ir antes
	api.Get("/", h.getAllContacts)
	api.Get("/:id", h.getContactByID)
	api.Post("/", h.createContact)
	api.Put("/:id", h.updateContact)
	api.Delete("/:id", h.deleteContact)
}


func (h *ContactHandler) createContact(c *fiber.Ctx) error {
    var contacto db.ContactoCLI
    if err := c.BodyParser(&contacto); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Formato de solicitud inválido",
        })
    }

    id, err := h.repo.Insert(contacto)
    if err != nil {
        if err == db.ErrInvalidData {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": "Razón social y RUC/CI son obligatorios",
            })
        }
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Error al crear el contacto",
        })
    }

    contacto.IDContacto = id
    return c.Status(fiber.StatusCreated).JSON(contacto)
}

func (h *ContactHandler) updateContact(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "ID de contacto no válido",
        })
    }

    var contacto db.ContactoCLI
    if err := c.BodyParser(&contacto); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Formato de solicitud inválido",
        })
    }

    contacto.IDContacto = id
    if err := h.repo.Update(contacto); err != nil {
        switch err {
        case db.ErrInvalidData:
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": "Datos de contacto no válidos",
            })
        case db.ErrContactNotFound:
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
                "error": "Contacto no encontrado",
            })
        default:
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": "Error al actualizar el contacto",
            })
        }
    }

    return c.Status(fiber.StatusOK).JSON(contacto)
}

func (h *ContactHandler) deleteContact(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "ID de contacto no válido",
        })
    }

    if err := h.repo.Delete(id); err != nil {
        if err == db.ErrContactNotFound {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
                "error": "Contacto no encontrado",
            })
        }
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Error al eliminar el contacto",
        })
    }

    return c.SendStatus(fiber.StatusNoContent)
}

func (h *ContactHandler) getContactByID(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "ID de contacto no válido",
        })
    }

    contacto, err := h.repo.GetByID(id)
    if err != nil {
        if err == db.ErrContactNotFound {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
                "error": "Contacto no encontrado",
            })
        }
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Error al obtener el contacto",
        })
    }

    return c.Status(fiber.StatusOK).JSON(contacto)
}

func (h *ContactHandler) getAllContacts(c *fiber.Ctx) error {
    contactos, err := h.repo.GetAll()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Error al obtener los contactos",
        })
    }

    if contactos == nil {
        contactos = []db.ContactoCLI{} // Devuelve un array vacío en lugar de null
    }

    return c.Status(fiber.StatusOK).JSON(contactos)
}

func (h *ContactHandler) searchContacts(c *fiber.Ctx) error {
    query := c.Query("q")
    if query == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Parámetro 'q' es obligatorio",
        })
    }

    contactos, err := h.repo.Search(query)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Error al buscar los contactos",
        })
    }

    return c.Status(fiber.StatusOK).JSON(contactos)
}
