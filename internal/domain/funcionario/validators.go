package funcionario

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	ErrInvalidRUTFormat     = errors.New("formato de RUT inválido, debe ser 12345678-9 o 12345678-K")
	ErrInvalidRUTCheckDigit = errors.New("dígito verificador del RUT es inválido")
	ErrInvalidCelularFormat = errors.New("formato de celular inválido, debe ser +56912345678")
	ErrEmailTooLong         = errors.New("el email no puede superar los 255 caracteres")
	ErrInvalidMeasurement   = errors.New("medida fuera del rango válido")
)

// ValidateRUT valida el formato y dígito verificador de un RUT chileno
func ValidateRUT(rut string) error {
	if rut == "" {
		return nil // RUT vacío puede ser válido dependiendo del contexto
	}

	// Formato esperado: 12345678-9 o 12345678-K
	rutRegex := regexp.MustCompile(`^(\d{1,8})-([0-9Kk])$`)
	matches := rutRegex.FindStringSubmatch(rut)

	if matches == nil {
		return ErrInvalidRUTFormat
	}

	number := matches[1]
	dv := strings.ToUpper(matches[2])

	// Calcular dígito verificador
	calculatedDV := calculateRUTCheckDigit(number)

	if calculatedDV != dv {
		return ErrInvalidRUTCheckDigit
	}

	return nil
}

// calculateRUTCheckDigit calcula el dígito verificador de un RUT
func calculateRUTCheckDigit(rut string) string {
	var sum int
	multiplier := 2

	// Recorrer el RUT de derecha a izquierda
	for i := len(rut) - 1; i >= 0; i-- {
		digit, _ := strconv.Atoi(string(rut[i]))
		sum += digit * multiplier
		multiplier++
		if multiplier > 7 {
			multiplier = 2
		}
	}

	remainder := sum % 11
	dv := 11 - remainder

	switch dv {
	case 11:
		return "0"
	case 10:
		return "K"
	default:
		return strconv.Itoa(dv)
	}
}

// ValidateCelular valida el formato de un celular chileno
func ValidateCelular(celular string) error {
	if celular == "" {
		return nil // Celular vacío puede ser válido (campo opcional)
	}

	// Formato esperado: +56912345678 (12-13 caracteres)
	celularRegex := regexp.MustCompile(`^\+56\d{9}$`)

	if !celularRegex.MatchString(celular) {
		return ErrInvalidCelularFormat
	}

	if len(celular) < 12 || len(celular) > 13 {
		return ErrInvalidCelularFormat
	}

	return nil
}

// ValidateEmail valida el email
func ValidateEmail(email string) error {
	if len(email) > 255 {
		return ErrEmailTooLong
	}
	return nil
}

// ValidateMedidas valida las medidas corporales según los rangos del contrato
func ValidateMedidas(m *MedidasFuncionario) error {
	if m == nil {
		return nil
	}

	// Validar estatura: 0 < x < 3 metros
	if m.EstaturaM != nil {
		if *m.EstaturaM <= 0 || *m.EstaturaM >= 3 {
			return fmt.Errorf("%w: estatura_m debe estar entre 0 y 3 metros", ErrInvalidMeasurement)
		}
	}

	// Validar pecho: 0 < x < 200 cm
	if m.PechoCm != nil {
		if *m.PechoCm <= 0 || *m.PechoCm >= 200 {
			return fmt.Errorf("%w: pecho_cm debe estar entre 0 y 200 cm", ErrInvalidMeasurement)
		}
	}

	// Validar cintura: 0 < x < 200 cm
	if m.CinturaCm != nil {
		if *m.CinturaCm <= 0 || *m.CinturaCm >= 200 {
			return fmt.Errorf("%w: cintura_cm debe estar entre 0 y 200 cm", ErrInvalidMeasurement)
		}
	}

	// Validar cadera: 0 < x < 200 cm
	if m.CaderaCm != nil {
		if *m.CaderaCm <= 0 || *m.CaderaCm >= 200 {
			return fmt.Errorf("%w: cadera_cm debe estar entre 0 y 200 cm", ErrInvalidMeasurement)
		}
	}

	// Validar manga: 0 < x < 100 cm
	if m.MangaCm != nil {
		if *m.MangaCm <= 0 || *m.MangaCm >= 100 {
			return fmt.Errorf("%w: manga_cm debe estar entre 0 y 100 cm", ErrInvalidMeasurement)
		}
	}

	return nil
}

// Validate valida los campos del funcionario
func (f *Funcionario) Validate() error {
	if f.Nombres == "" {
		return errors.New("nombres es requerido")
	}
	if f.ApellidoPaterno == "" {
		return errors.New("apellido_paterno es requerido")
	}
	if f.Email == "" {
		return errors.New("email es requerido")
	}
	if f.RutFuncionario == "" {
		return errors.New("rut_funcionario es requerido")
	}

	// Validar RUT
	if err := ValidateRUT(f.RutFuncionario); err != nil {
		return err
	}

	// Validar celular
	if err := ValidateCelular(f.Celular); err != nil {
		return err
	}

	// Validar email
	if err := ValidateEmail(f.Email); err != nil {
		return err
	}

	return nil
}
