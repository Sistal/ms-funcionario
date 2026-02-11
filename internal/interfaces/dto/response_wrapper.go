package dto

// SuccessResponse representa una respuesta exitosa estándar
type SuccessResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message,omitempty"`
}

// ErrorResponse representa una respuesta de error estándar
type ErrorResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Errors  []FieldError `json:"errors,omitempty"`
}

// FieldError representa un error específico de un campo
type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// PaginatedResponse representa una respuesta paginada estándar
type PaginatedResponse struct {
	Success bool           `json:"success"`
	Data    interface{}    `json:"data"`
	Meta    PaginationMeta `json:"meta"`
}

// PaginationMeta contiene metadata de paginación
type PaginationMeta struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
}

// NewSuccessResponse crea una respuesta exitosa
func NewSuccessResponse(data interface{}) SuccessResponse {
	return SuccessResponse{
		Success: true,
		Data:    data,
	}
}

// NewSuccessResponseWithMessage crea una respuesta exitosa con mensaje
func NewSuccessResponseWithMessage(data interface{}, message string) SuccessResponse {
	return SuccessResponse{
		Success: true,
		Data:    data,
		Message: message,
	}
}

// NewErrorResponse crea una respuesta de error
func NewErrorResponse(message string, errors ...FieldError) ErrorResponse {
	return ErrorResponse{
		Success: false,
		Message: message,
		Errors:  errors,
	}
}

// NewPaginatedResponse crea una respuesta paginada
func NewPaginatedResponse(data interface{}, page, limit int, total int64) PaginatedResponse {
	totalPages := int(total) / limit
	if int(total)%limit != 0 {
		totalPages++
	}

	return PaginatedResponse{
		Success: true,
		Data:    data,
		Meta: PaginationMeta{
			Page:       page,
			Limit:      limit,
			Total:      total,
			TotalPages: totalPages,
		},
	}
}
