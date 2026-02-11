package dto

import (
	"time"

	"github.com/Sistal/ms-funcionario/internal/domain/funcionario"
)

// CreateFuncionarioRequest es el DTO para crear un funcionario
type CreateFuncionarioRequest struct {
	RutFuncionario   string `json:"rut_funcionario" binding:"required"`
	Nombres          string `json:"nombres" binding:"required"`
	ApellidoPaterno  string `json:"apellido_paterno" binding:"required"`
	ApellidoMaterno  string `json:"apellido_materno"`
	Celular          string `json:"celular"`
	Telefono         string `json:"telefono"`
	Email            string `json:"email" binding:"required,email"`
	Direccion        string `json:"direccion"`
	IDGenero         *int   `json:"id_genero" binding:"required"`
	IDEstado         *int   `json:"id_estado"` // default: 10 (Activo)
	IDSucursal       *int   `json:"id_sucursal" binding:"required"`
	IDCargo          *int   `json:"id_cargo" binding:"required"`
	IDUsuario        *int   `json:"id_usuario"`
	IDEmpresaCliente *int   `json:"id_empresa_cliente"`
	IDSegmento       *int   `json:"id_segmento"`
}

// UpdateFuncionarioRequest es el DTO para actualizar un funcionario
type UpdateFuncionarioRequest struct {
	Nombres         string `json:"nombres"`
	ApellidoPaterno string `json:"apellido_paterno"`
	ApellidoMaterno string `json:"apellido_materno"`
	Celular         string `json:"celular"`
	Email           string `json:"email"`
	IDGenero        *int   `json:"id_genero"`
	IDEstado        *int   `json:"id_estado"`
	IDSucursal      *int   `json:"id_sucursal"`
	IDCargo         *int   `json:"id_cargo"`
	// Campos adicionales (no en contrato pero presentes en BD)
	Telefono         string `json:"telefono"`
	Direccion        string `json:"direccion"`
	IDEmpresaCliente *int   `json:"id_empresa_cliente"`
	IDSegmento       *int   `json:"id_segmento"`
}

// GeneroDTO representa la información del género
type GeneroDTO struct {
	IDGenero     int    `json:"id_genero"`
	NombreGenero string `json:"nombre_genero"`
}

// EstadoDTO representa la información del estado
type EstadoDTO struct {
	IDEstado     int    `json:"id_estado"`
	NombreEstado string `json:"nombre_estado"`
	TablaEstado  string `json:"tabla_estado,omitempty"`
}

// CargoDTO representa la información del cargo
type CargoDTO struct {
	IDCargo     int    `json:"id_cargo"`
	NombreCargo string `json:"nombre_cargo"`
}

// SucursalDTO representa la información de la sucursal
type SucursalDTO struct {
	IDSucursal        int        `json:"id_sucursal"`
	NombreSucursal    string     `json:"nombre_sucursal"`
	Direccion         string     `json:"direccion,omitempty"`
	EstadoSucursal    *int       `json:"estado_sucursal,omitempty"`
	Estado            *EstadoDTO `json:"estado,omitempty"`
	TotalFuncionarios *int       `json:"total_funcionarios,omitempty"`
}

// RolDTO representa el rol de un usuario
type RolDTO struct {
	IDRol     int    `json:"id_rol"`
	NombreRol string `json:"nombre_rol"`
}

// UsuarioDTO representa la información del usuario asociado
type UsuarioDTO struct {
	IDUsuario      int     `json:"id_usuario"`
	NombreUsuario  string  `json:"nombre_usuario"`
	NombreCompleto string  `json:"nombre_completo,omitempty"`
	Rol            *RolDTO `json:"rol,omitempty"`
}

// FuncionarioResponse es el DTO de respuesta para un funcionario
type FuncionarioResponse struct {
	IDFuncionario     int              `json:"id_funcionario"`
	RutFuncionario    string           `json:"rut_funcionario"`
	Nombres           string           `json:"nombres"`
	ApellidoPaterno   string           `json:"apellido_paterno"`
	ApellidoMaterno   string           `json:"apellido_materno,omitempty"`
	NombreCompleto    string           `json:"nombre_completo"`
	Celular           string           `json:"celular,omitempty"`
	Email             string           `json:"email"`
	TallasRegistradas bool             `json:"tallas_registradas"`
	Genero            *GeneroDTO       `json:"genero,omitempty"`
	Medidas           *MedidasResponse `json:"medidas,omitempty"`
	Usuario           *UsuarioDTO      `json:"usuario,omitempty"`
	Estado            *EstadoDTO       `json:"estado,omitempty"`
	Sucursal          *SucursalDTO     `json:"sucursal,omitempty"`
	Cargo             *CargoDTO        `json:"cargo,omitempty"`
	FechaCreacion     *time.Time       `json:"fecha_creacion,omitempty"`
	FechaModificacion *time.Time       `json:"fecha_modificacion,omitempty"`
	// Campos adicionales (no en contrato pero presentes en BD)
	Telefono         string `json:"telefono,omitempty"`
	Direccion        string `json:"direccion,omitempty"`
	IDEmpresaCliente *int   `json:"id_empresa_cliente,omitempty"`
	IDSegmento       *int   `json:"id_segmento,omitempty"`
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
	IDMedidas     int        `json:"id_medidas"`
	IDFuncionario int        `json:"id_funcionario,omitempty"`
	EstaturaM     *float64   `json:"estatura_m"`
	PechoCm       *float64   `json:"pecho_cm"`
	CinturaCm     *float64   `json:"cintura_cm"`
	CaderaCm      *float64   `json:"cadera_cm"`
	MangaCm       *float64   `json:"manga_cm"`
	FechaInicio   *time.Time `json:"fecha_inicio,omitempty"`
	FechaFin      *time.Time `json:"fecha_fin,omitempty"`
}

// FuncionarioFilterRequest es el DTO para filtrar funcionarios
type FuncionarioFilterRequest struct {
	Page              int    `form:"page"`
	Limit             int    `form:"limit"`
	IDSucursal        *int   `form:"id_sucursal"`
	IDEstado          *int   `form:"id_estado"`
	IDGenero          *int   `form:"id_genero"`
	IDCargo           *int   `form:"id_cargo"`
	TallasRegistradas *bool  `form:"tallas_registradas"`
	Search            string `form:"search"`
	SortBy            string `form:"sort_by"`
	Order             string `form:"order"`
	// Campos adicionales (no en contrato)
	RutFuncionario   string `form:"rut_funcionario"`
	Email            string `form:"email"`
	IDEmpresaCliente *int   `form:"id_empresa_cliente"`
	IDSegmento       *int   `form:"id_segmento"`
	Offset           int    `form:"offset"` // legacy
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

	// Calcular nombre completo
	nombreCompleto := f.Nombres + " " + f.ApellidoPaterno
	if f.ApellidoMaterno != "" {
		nombreCompleto += " " + f.ApellidoMaterno
	}

	// Mapear Genero
	var generoDTO *GeneroDTO
	if f.Genero != nil {
		generoDTO = &GeneroDTO{
			IDGenero:     f.Genero.IDGenero,
			NombreGenero: f.Genero.NombreGenero,
		}
	}

	// Mapear Estado
	var estadoDTO *EstadoDTO
	if f.Estado != nil {
		estadoDTO = &EstadoDTO{
			IDEstado:     f.Estado.IDEstado,
			NombreEstado: f.Estado.NombreEstado,
			TablaEstado:  f.Estado.TablaEstado,
		}
	}

	// Mapear Cargo
	var cargoDTO *CargoDTO
	if f.Cargo != nil {
		cargoDTO = &CargoDTO{
			IDCargo:     f.Cargo.IDCargo,
			NombreCargo: f.Cargo.NombreCargo,
		}
	}

	// Mapear Sucursal
	var sucursalDTO *SucursalDTO
	if f.Sucursal != nil {
		sucursalDTO = &SucursalDTO{
			IDSucursal:     f.Sucursal.IDSucursal,
			NombreSucursal: f.Sucursal.NombreSucursal,
			Direccion:      f.Sucursal.Direccion,
			EstadoSucursal: &f.Sucursal.EstadoSucursal,
		}
	}

	// Mapear Medidas
	var medidasDTO *MedidasResponse
	if f.Medidas != nil {
		medidasDTO = ToMedidasResponse(f.Medidas, f.IDFuncionario)
	}

	// TODO: Mapear Usuario cuando esté disponible desde MS-Authentication
	var usuarioDTO *UsuarioDTO
	if f.IDUsuario != nil {
		// Por ahora solo con ID, en el futuro se puede consultar a MS-Auth
		usuarioDTO = &UsuarioDTO{
			IDUsuario: *f.IDUsuario,
		}
	}

	return &FuncionarioResponse{
		IDFuncionario:     f.IDFuncionario,
		RutFuncionario:    f.RutFuncionario,
		Nombres:           f.Nombres,
		ApellidoPaterno:   f.ApellidoPaterno,
		ApellidoMaterno:   f.ApellidoMaterno,
		NombreCompleto:    nombreCompleto,
		Celular:           f.Celular,
		Telefono:          f.Telefono,
		Email:             f.Email,
		TallasRegistradas: f.TallasRegistradas,
		Direccion:         f.Direccion,
		FechaCreacion:     f.FechaCreacion,
		FechaModificacion: f.FechaModificacion,
		Genero:            generoDTO,
		Medidas:           medidasDTO,
		Usuario:           usuarioDTO,
		Estado:            estadoDTO,
		Sucursal:          sucursalDTO,
		Cargo:             cargoDTO,
		IDEmpresaCliente:  f.IDEmpresaCliente,
		IDSegmento:        f.IDSegmento,
	}
}

func ToFuncionarioResponseList(funcionarios []*funcionario.Funcionario) []*FuncionarioResponse {
	responses := make([]*FuncionarioResponse, len(funcionarios))
	for i, f := range funcionarios {
		responses[i] = ToFuncionarioResponse(f)
	}
	return responses
}

func ToMedidasResponse(m *funcionario.MedidasFuncionario, idFuncionario ...int) *MedidasResponse {
	if m == nil {
		return nil
	}

	response := &MedidasResponse{
		IDMedidas:   m.IDMedidas,
		EstaturaM:   m.EstaturaM,
		PechoCm:     m.PechoCm,
		CinturaCm:   m.CinturaCm,
		CaderaCm:    m.CaderaCm,
		MangaCm:     m.MangaCm,
		FechaInicio: m.FechaInicio,
		FechaFin:    m.FechaFin,
	}

	// Incluir ID de funcionario si se proporciona
	if len(idFuncionario) > 0 {
		response.IDFuncionario = idFuncionario[0]
	}

	return response
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
	// Default estado to 10 (Activo) if not provided
	idEstado := req.IDEstado
	if idEstado == nil {
		defaultEstado := 10
		idEstado = &defaultEstado
	}

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
		IDEstado:         idEstado,
		IDSucursal:       req.IDSucursal,
		IDCargo:          req.IDCargo,
		IDUsuario:        req.IDUsuario,
		IDEmpresaCliente: req.IDEmpresaCliente,
		IDSegmento:       req.IDSegmento,
	}
}

func (req *UpdateFuncionarioRequest) ToFuncionario(id int) *funcionario.Funcionario {
	return &funcionario.Funcionario{
		IDFuncionario:    id,
		Nombres:          req.Nombres,
		ApellidoPaterno:  req.ApellidoPaterno,
		ApellidoMaterno:  req.ApellidoMaterno,
		Celular:          req.Celular,
		Telefono:         req.Telefono,
		Email:            req.Email,
		Direccion:        req.Direccion,
		IDGenero:         req.IDGenero,
		IDEstado:         req.IDEstado,
		IDSucursal:       req.IDSucursal,
		IDCargo:          req.IDCargo,
		IDEmpresaCliente: req.IDEmpresaCliente,
		IDSegmento:       req.IDSegmento,
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
		IDGenero:          req.IDGenero,
		TallasRegistradas: req.TallasRegistradas,
		Search:            req.Search,
		SortBy:            req.SortBy,
		Order:             req.Order,
		Limit:             req.Limit,
		Offset:            req.Offset,
	}
}

// ProfileResponse es el DTO detallado para el perfil del empleado (BFF contract)
type ProfileResponse struct {
	ID                int              `json:"id"`
	Rut               string           `json:"rut"`
	Nombres           string           `json:"nombres"`
	ApellidoPaterno   string           `json:"apellido_paterno"`
	ApellidoMaterno   string           `json:"apellido_materno"`
	Cargo             *CargoDTO        `json:"cargo,omitempty"`
	Sucursal          *SucursalDTO     `json:"sucursal,omitempty"`
	Email             string           `json:"email"`
	Celular           string           `json:"celular"`
	TallasRegistradas bool             `json:"tallas_registradas"`
	Medidas           *MedidasResponse `json:"medidas,omitempty"`
}

// ContactUpdateRequest es el DTO para actualizar el contacto
type ContactUpdateRequest struct {
	Celular string `json:"celular" binding:"required"`
	Email   string `json:"email" binding:"required,email"`
}

// TransferRequest es el DTO para solicitar traslado
type TransferRequest struct {
	TargetBranchID int    `json:"target_branch_id" binding:"required"`
	Reason         string `json:"reason"`
}

func ToProfileResponse(f *funcionario.Funcionario) *ProfileResponse {
	if f == nil {
		return nil
	}

	var cargoDTO *CargoDTO
	if f.Cargo != nil {
		cargoDTO = &CargoDTO{
			IDCargo:     f.Cargo.IDCargo,
			NombreCargo: f.Cargo.NombreCargo,
		}
	}

	var sucursalDTO *SucursalDTO
	if f.Sucursal != nil {
		sucursalDTO = &SucursalDTO{
			IDSucursal:     f.Sucursal.IDSucursal,
			NombreSucursal: f.Sucursal.NombreSucursal,
			Direccion:      f.Sucursal.Direccion,
		}
	}

	var medidasResp *MedidasResponse
	if f.Medidas != nil {
		medidasResp = ToMedidasResponse(f.Medidas, f.IDFuncionario)
	}

	return &ProfileResponse{
		ID:                f.IDFuncionario,
		Rut:               f.RutFuncionario,
		Nombres:           f.Nombres,
		ApellidoPaterno:   f.ApellidoPaterno,
		ApellidoMaterno:   f.ApellidoMaterno,
		Cargo:             cargoDTO,
		Sucursal:          sucursalDTO,
		Email:             f.Email,
		Celular:           f.Celular,
		TallasRegistradas: f.TallasRegistradas,
		Medidas:           medidasResp,
	}
}
