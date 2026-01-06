package dto

import (
	"time"

	"github.com/Sistal/ms-funcionario/internal/domain/funcionario"
)

// CreateFuncionarioRequest es el DTO para crear un funcionario
type CreateFuncionarioRequest struct {
	RutFuncionario   string  `json:"rut_funcionario" binding:"required"`
	Nombres          string  `json:"nombres" binding:"required"`
	ApellidoPaterno  string  `json:"apellido_paterno" binding:"required"`
	ApellidoMaterno  string  `json:"apellido_materno"`
	Celular          string  `json:"celular"`
	Telefono         string  `json:"telefono"`
	Email            string  `json:"email" binding:"required,email"`
	Direccion        string  `json:"direccion"`
	IDGenero         *int    `json:"id_genero"`
	IDEmpresaCliente *int    `json:"id_empresa_cliente" binding:"required"`
	IDSucursal       *int    `json:"id_sucursal"`
	IDSegmento       *int    `json:"id_segmento"`
	IDCargo          *int    `json:"id_cargo"`
	IDUsuario        *int    `json:"id_usuario"`
}

// UpdateFuncionarioRequest es el DTO para actualizar un funcionario
type UpdateFuncionarioRequest struct {
	RutFuncionario   string  `json:"rut_funcionario"`
	Nombres          string  `json:"nombres"`
	ApellidoPaterno  string  `json:"apellido_paterno"`
	ApellidoMaterno  string  `json:"apellido_materno"`
	Celular          string  `json:"celular"`
	Telefono         string  `json:"telefono"`
	Email            string  `json:"email"`
	Direccion        string  `json:"direccion"`
	IDGenero         *int    `json:"id_genero"`
	IDEmpresaCliente *int    `json:"id_empresa_cliente"`
	IDSucursal       *int    `json:"id_sucursal"`
	IDSegmento       *int    `json:"id_segmento"`
	IDCargo          *int    `json:"id_cargo"`
	IDEstado         *int    `json:"id_estado"`
}

// FuncionarioResponse es el DTO de respuesta para un funcionario
type FuncionarioResponse struct {
	IDFuncionario     int        `json:"id_funcionario"`
	RutFuncionario    string     `json:"rut_funcionario"`
	Nombres           string     `json:"nombres"`
	ApellidoPaterno   string     `json:"apellido_paterno"`
	ApellidoMaterno   string     `json:"apellido_materno"`
	Celular           string     `json:"celular"`
	Telefono          string     `json:"telefono"`
	Email             string     `json:"email"`
	TallasRegistradas bool       `json:"tallas_registradas"`
	Direccion         string     `json:"direccion"`
	FechaCreacion     *time.Time `json:"fecha_creacion"`
	FechaModificacion *time.Time `json:"fecha_modificacion"`
	IDGenero          *int       `json:"id_genero"`
	IDMedidas         *int       `json:"id_medidas"`
	IDUsuario         *int       `json:"id_usuario"`
	IDEstado          *int       `json:"id_estado"`
	IDSucursal        *int       `json:"id_sucursal"`
	IDEmpresaCliente  *int       `json:"id_empresa_cliente"`
	IDSegmento        *int       `json:"id_segmento"`
	IDCargo           *int       `json:"id_cargo"`
}

// CreateMedidasRequest es el DTO para crear/actualizar medidas de un funcionario
type CreateMedidasRequest struct {
	EstaturaM   *float64   `json:"estatura_m" binding:"required"`
	PechoCm     *float64   `json:"pecho_cm"`
	CinturaCm   *float64   `json:"cintura_cm"`
	CaderaCm    *float64   `json:"cadera_cm"`
	MangaCm     *float64   `json:"manga_cm"`
	FechaInicio *time.Time `json:"fecha_inicio"`
}

// UpdateMedidasRequest es el DTO para actualizar medidas
type UpdateMedidasRequest struct {
	EstaturaM *float64   `json:"estatura_m"`
	PechoCm   *float64   `json:"pecho_cm"`
	CinturaCm *float64   `json:"cintura_cm"`
	CaderaCm  *float64   `json:"cadera_cm"`
	MangaCm   *float64   `json:"manga_cm"`
	FechaFin  *time.Time `json:"fecha_fin"`
}

// MedidasResponse es el DTO de respuesta para medidas
type MedidasResponse struct {
	IDMedidas   int        `json:"id_medidas"`
	EstaturaM   *float64   `json:"estatura_m"`
	PechoCm     *float64   `json:"pecho_cm"`
	CinturaCm   *float64   `json:"cintura_cm"`
	CaderaCm    *float64   `json:"cadera_cm"`
	MangaCm     *float64   `json:"manga_cm"`
	FechaInicio *time.Time `json:"fecha_inicio"`
	FechaFin    *time.Time `json:"fecha_fin"`
}

// FuncionarioFilterRequest es el DTO para filtrar funcionarios
type FuncionarioFilterRequest struct {
	RutFuncionario    string `form:"rut_funcionario"`
	Email             string `form:"email"`
	IDEmpresaCliente  *int   `form:"id_empresa_cliente"`
	IDSucursal        *int   `form:"id_sucursal"`
	IDSegmento        *int   `form:"id_segmento"`
	IDEstado          *int   `form:"id_estado"`
	IDCargo           *int   `form:"id_cargo"`
	TallasRegistradas *bool  `form:"tallas_registradas"`
	Limit             int    `form:"limit"`
	Offset            int    `form:"offset"`
}

// PaginatedFuncionariosResponse es el DTO de respuesta paginada
type PaginatedFuncionariosResponse struct {
	Data       []*FuncionarioResponse `json:"data"`
	Total      int64                  `json:"total"`
	Limit      int                    `json:"limit"`
	Offset     int                    `json:"offset"`
	TotalPages int                    `json:"total_pages"`
}

// Convertidores de dominio a DTO

func ToFuncionarioResponse(f *funcionario.Funcionario) *FuncionarioResponse {
	if f == nil {
		return nil
	}
	return &FuncionarioResponse{
		IDFuncionario:     f.IDFuncionario,
		RutFuncionario:    f.RutFuncionario,
		Nombres:           f.Nombres,
		ApellidoPaterno:   f.ApellidoPaterno,
		ApellidoMaterno:   f.ApellidoMaterno,
		Celular:           f.Celular,
		Telefono:          f.Telefono,
		Email:             f.Email,
		TallasRegistradas: f.TallasRegistradas,
		Direccion:         f.Direccion,
		FechaCreacion:     f.FechaCreacion,
		FechaModificacion: f.FechaModificacion,
		IDGenero:          f.IDGenero,
		IDMedidas:         f.IDMedidas,
		IDUsuario:         f.IDUsuario,
		IDEstado:          f.IDEstado,
		IDSucursal:        f.IDSucursal,
		IDEmpresaCliente:  f.IDEmpresaCliente,
		IDSegmento:        f.IDSegmento,
		IDCargo:           f.IDCargo,
	}
}

func ToFuncionarioResponseList(funcionarios []*funcionario.Funcionario) []*FuncionarioResponse {
	responses := make([]*FuncionarioResponse, len(funcionarios))
	for i, f := range funcionarios {
		responses[i] = ToFuncionarioResponse(f)
	}
	return responses
}

func ToMedidasResponse(m *funcionario.MedidasFuncionario) *MedidasResponse {
	if m == nil {
		return nil
	}
	return &MedidasResponse{
		IDMedidas:   m.IDMedidas,
		EstaturaM:   m.EstaturaM,
		PechoCm:     m.PechoCm,
		CinturaCm:   m.CinturaCm,
		CaderaCm:    m.CaderaCm,
		MangaCm:     m.MangaCm,
		FechaInicio: m.FechaInicio,
		FechaFin:    m.FechaFin,
	}
}

func ToMedidasResponseList(medidas []*funcionario.MedidasFuncionario) []*MedidasResponse {
	responses := make([]*MedidasResponse, len(medidas))
	for i, m := range medidas {
		responses[i] = ToMedidasResponse(m)
	}
	return responses
}

// Convertidores de DTO a dominio

func (req *CreateFuncionarioRequest) ToFuncionario() *funcionario.Funcionario {
	return &funcionario.Funcionario{
		RutFuncionario:   req.RutFuncionario,
		Nombres:          req.Nombres,
		ApellidoPaterno:  req.ApellidoPaterno,
		ApellidoMaterno:  req.ApellidoMaterno,
		Celular:          req.Celular,
		Telefono:         req.Telefono,
		Email:            req.Email,
		Direccion:        req.Direccion,
		IDGenero:         req.IDGenero,
		IDEmpresaCliente: req.IDEmpresaCliente,
		IDSucursal:       req.IDSucursal,
		IDSegmento:       req.IDSegmento,
		IDCargo:          req.IDCargo,
		IDUsuario:        req.IDUsuario,
	}
}

func (req *UpdateFuncionarioRequest) ToFuncionario(id int) *funcionario.Funcionario {
	return &funcionario.Funcionario{
		IDFuncionario:    id,
		RutFuncionario:   req.RutFuncionario,
		Nombres:          req.Nombres,
		ApellidoPaterno:  req.ApellidoPaterno,
		ApellidoMaterno:  req.ApellidoMaterno,
		Celular:          req.Celular,
		Telefono:         req.Telefono,
		Email:            req.Email,
		Direccion:        req.Direccion,
		IDGenero:         req.IDGenero,
		IDEmpresaCliente: req.IDEmpresaCliente,
		IDSucursal:       req.IDSucursal,
		IDSegmento:       req.IDSegmento,
		IDCargo:          req.IDCargo,
		IDEstado:         req.IDEstado,
	}
}

func (req *CreateMedidasRequest) ToMedidas() *funcionario.MedidasFuncionario {
	return &funcionario.MedidasFuncionario{
		EstaturaM:   req.EstaturaM,
		PechoCm:     req.PechoCm,
		CinturaCm:   req.CinturaCm,
		CaderaCm:    req.CaderaCm,
		MangaCm:     req.MangaCm,
		FechaInicio: req.FechaInicio,
	}
}

func (req *FuncionarioFilterRequest) ToFilter() funcionario.FuncionarioFilter {
	return funcionario.FuncionarioFilter{
		RutFuncionario:    req.RutFuncionario,
		Email:             req.Email,
		IDEmpresaCliente:  req.IDEmpresaCliente,
		IDSucursal:        req.IDSucursal,
		IDSegmento:        req.IDSegmento,
		IDEstado:          req.IDEstado,
		IDCargo:           req.IDCargo,
		TallasRegistradas: req.TallasRegistradas,
		Limit:             req.Limit,
		Offset:            req.Offset,
	}
}
