package validation

import "regexp"

// IsValidProductCode acepta:
// - Códigos EAN/UPC válidos (8, 12, 13 dígitos)
// - Códigos personalizados alfanuméricos de 2+ caracteres
func IsValidProductCode(code string) bool {
	// EAN/UPC: solo dígitos con longitud específica
	if matched, _ := regexp.MatchString(`^\d{8}$|^\d{12}$|^\d{13}$`, code); matched {
		return true
	}
	// Códigos alfanuméricos personalizados: mínimo 2 caracteres
	if matched, _ := regexp.MatchString(`^[a-zA-Z0-9_-]{2,}$`, code); matched {
		return true
	}
	return false
}
