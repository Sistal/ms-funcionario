# Microservicio de Funcionarios - Documentación Técnica

## Descripción General

Microservicio RESTful para la gestión integral de funcionarios en un sistema de control de uniformes empresariales. Parte de una arquitectura de microservicios que maneja el ciclo completo de información de empleados, incluyendo datos personales, medidas corporales, asignaciones empresariales y gestión de tallas.

## Tecnologías Utilizadas

### Lenguaje y Framework
- **Go (Golang)** 1.23.0+
- **Gin Web Framework** 1.11.0 - Framework HTTP de alto rendimiento

### Base de Datos
- **PostgreSQL** - Base de datos relacional principal
- **GORM** 1.31.1 - ORM (Object-Relational Mapping) para Go
- **pgx/v5** 5.6.0 - Driver PostgreSQL nativo

### Herramientas y Dependencias
- **UUID** (Google) 1.6.0 - Generación de identificadores únicos
- **Docker & Docker Compose** - Contenedorización y orquestación
- **Makefile** - Automatización de tareas de desarrollo

## Arquitectura

El proyecto sigue una arquitectura hexagonal (puertos y adaptadores) con separación clara de responsabilidades:

```
ms-funcionario/
├── cmd/
│   └── api/
│       └── main.go                      # Punto de entrada de la aplicación
├── config/
│   └── config.go                        # Configuración de la aplicación
├── internal/
│   ├── domain/                          # Capa de dominio (entidades y reglas de negocio)
│   │   └── funcionario/
│   │       ├── funcionario.go           # Entidades Funcionario y MedidasFuncionario
│   │       └── repository.go            # Interfaces de repositorio
│   ├── application/                     # Capa de aplicación (casos de uso)
│   │   └── services/
│   │       └── funcionario_service.go   # Lógica de negocio
│   ├── infrastructure/                  # Capa de infraestructura
│   │   ├── database/
│   │   │   └── postgres.go              # Conexión a PostgreSQL
│   │   └── repository/
│   │       └── funcionario_repository.go # Implementación de repositorios
│   └── interfaces/                      # Capa de interfaces (adaptadores)
│       ├── dto/
│       │   └── funcionario_dto.go       # Data Transfer Objects
│       └── http/
│           ├── handler/
│           │   └── funcionario_handler.go # Controladores HTTP
│           └── router/
│               └── router.go             # Configuración de rutas
├── docker-compose.yml                   # Orquestación de contenedores
├── Dockerfile                           # Imagen Docker del servicio
├── go.mod                               # Dependencias del proyecto
├── Makefile                             # Comandos de automatización
└── README.md                            # Documentación básica
```

### Capas de la Arquitectura

#### 1. **Capa de Dominio** (`internal/domain`)
Contiene las entidades de negocio y reglas fundamentales:

**Entidades:**
- `Funcionario`: Representa un empleado con todos sus datos personales, laborales y relaciones
- `MedidasFuncionario`: Medidas corporales para confección de uniformes con historial temporal

**Interfaces de Repositorio:**
- Define contratos que deben cumplir las implementaciones de persistencia
- Desacopla la lógica de negocio de la infraestructura

#### 2. **Capa de Aplicación** (`internal/application`)
Implementa los casos de uso del negocio:

**Servicios:**
- `FuncionarioService`: Orquesta operaciones CRUD y lógica de negocio
- Validaciones de integridad de datos
- Gestión de medidas corporales con versionamiento

#### 3. **Capa de Infraestructura** (`internal/infrastructure`)
Implementaciones concretas de persistencia y conexiones externas:

**Repositorios:**
- `FuncionarioRepository`: Operaciones de base de datos para funcionarios
- `MedidasRepository`: Operaciones de base de datos para medidas

**Base de Datos:**
- Configuración y conexión a PostgreSQL con GORM

#### 4. **Capa de Interfaces** (`internal/interfaces`)
Adaptadores para comunicación externa:

**DTOs (Data Transfer Objects):**
- Estructuras optimizadas para transferencia de datos HTTP
- Validaciones de entrada con binding tags
- Conversores entre DTOs y entidades de dominio

**Handlers HTTP:**
- Controladores REST con manejo de errores
- Serialización/deserialización JSON
- Mapeo de códigos HTTP apropiados

## API REST - Endpoints

### 1. Gestión de Funcionarios

#### Crear Funcionario
```http
POST /api/v1/funcionarios
Content-Type: application/json

{
  "rut_funcionario": "12345678-9",
  "nombres": "Juan Pablo",
  "apellido_paterno": "González",
  "apellido_materno": "Pérez",
  "celular": "+56912345678",
  "telefono": "+56212345678",
  "email": "juan.gonzalez@empresa.cl",
  "direccion": "Av. Principal 123, Santiago",
  "id_genero": 1,
  "id_empresa_cliente": 5,
  "id_sucursal": 2,
  "id_segmento": 3,
  "id_cargo": 4
}
```

**Respuesta:** `201 Created`
```json
{
  "id_funcionario": 1,
  "rut_funcionario": "12345678-9",
  "nombres": "Juan Pablo",
  "apellido_paterno": "González",
  "apellido_materno": "Pérez",
  "celular": "+56912345678",
  "telefono": "+56212345678",
  "email": "juan.gonzalez@empresa.cl",
  "tallas_registradas": false,
  "direccion": "Av. Principal 123, Santiago",
  "fecha_creacion": "2026-01-05T00:00:00Z",
  "fecha_modificacion": null,
  "id_genero": 1,
  "id_medidas": null,
  "id_usuario": null,
  "id_estado": 1,
  "id_sucursal": 2,
  "id_empresa_cliente": 5,
  "id_segmento": 3,
  "id_cargo": 4
}
```

#### Obtener Funcionario por ID
```http
GET /api/v1/funcionarios/{id}
```

**Respuesta:** `200 OK`

#### Listar Todos los Funcionarios
```http
GET /api/v1/funcionarios
```

**Respuesta:** `200 OK` - Array de funcionarios

#### Actualizar Funcionario
```http
PUT /api/v1/funcionarios/{id}
Content-Type: application/json

{
  "nombres": "Juan Pablo",
  "apellido_paterno": "González",
  "apellido_materno": "Pérez",
  "celular": "+56912345678",
  "email": "juan.gonzalez@empresa.cl",
  "id_estado": 1
}
```

**Respuesta:** `200 OK`

#### Eliminar Funcionario
```http
DELETE /api/v1/funcionarios/{id}
```

**Respuesta:** `204 No Content`

### 2. Búsquedas y Filtros

#### Filtrar Funcionarios (con Paginación)
```http
GET /api/v1/funcionarios/filter?id_empresa_cliente=5&id_estado=1&limit=20&offset=0
```

**Parámetros de Query:**
- `rut_funcionario`: RUT del funcionario (exacto)
- `email`: Email del funcionario (exacto)
- `id_empresa_cliente`: ID de la empresa
- `id_sucursal`: ID de la sucursal
- `id_segmento`: ID del segmento
- `id_estado`: ID del estado
- `id_cargo`: ID del cargo
- `tallas_registradas`: Boolean - si tiene tallas registradas
- `limit`: Límite de resultados (default: 20)
- `offset`: Desplazamiento para paginación (default: 0)

**Respuesta:** `200 OK`
```json
{
  "data": [...],
  "total": 150,
  "limit": 20,
  "offset": 0,
  "total_pages": 8
}
```

#### Buscar por RUT
```http
GET /api/v1/funcionarios/rut/{rut}
```

**Ejemplo:** `GET /api/v1/funcionarios/rut/12345678-9`

**Respuesta:** `200 OK`

#### Listar por Empresa
```http
GET /api/v1/funcionarios/empresa/{id_empresa}
```

**Respuesta:** `200 OK` - Array de funcionarios de la empresa

#### Listar por Sucursal
```http
GET /api/v1/funcionarios/sucursal/{id_sucursal}
```

**Respuesta:** `200 OK`

#### Listar por Segmento
```http
GET /api/v1/funcionarios/segmento/{id_segmento}
```

**Respuesta:** `200 OK`

### 3. Gestión de Estado

#### Activar Funcionario
```http
PATCH /api/v1/funcionarios/{id}/activate
```

**Respuesta:** `200 OK`
```json
{
  "message": "funcionario activated successfully"
}
```

#### Desactivar Funcionario
```http
PATCH /api/v1/funcionarios/{id}/deactivate
```

**Respuesta:** `200 OK`
```json
{
  "message": "funcionario deactivated successfully"
}
```

### 4. Gestión de Medidas Corporales

#### Registrar Medidas
```http
POST /api/v1/funcionarios/{id}/medidas
Content-Type: application/json

{
  "estatura_m": 1.75,
  "pecho_cm": 95.5,
  "cintura_cm": 85.0,
  "cadera_cm": 98.0,
  "manga_cm": 62.5,
  "fecha_inicio": "2026-01-05T00:00:00Z"
}
```

**Respuesta:** `201 Created`
```json
{
  "id_medidas": 1,
  "estatura_m": 1.75,
  "pecho_cm": 95.5,
  "cintura_cm": 85.0,
  "cadera_cm": 98.0,
  "manga_cm": 62.5,
  "fecha_inicio": "2026-01-05T00:00:00Z",
  "fecha_fin": null
}
```

**Lógica:**
- Si el funcionario ya tiene medidas activas (sin `fecha_fin`), estas se cierran automáticamente
- El campo `tallas_registradas` del funcionario se actualiza a `true`
- Se crea el historial de medidas

#### Obtener Medidas Activas
```http
GET /api/v1/funcionarios/{id}/medidas
```

**Respuesta:** `200 OK` - Medidas sin `fecha_fin`

#### Actualizar Medidas Activas
```http
PUT /api/v1/funcionarios/{id}/medidas
Content-Type: application/json

{
  "pecho_cm": 96.0,
  "cintura_cm": 84.5
}
```

**Respuesta:** `200 OK`

**Nota:** Solo se actualizan los campos proporcionados

#### Obtener Historial de Medidas
```http
GET /api/v1/funcionarios/{id}/medidas/historial
```

**Respuesta:** `200 OK` - Array de todas las medidas ordenadas por fecha descendente

### 5. Health Check

#### Verificar Estado del Servicio
```http
GET /health
```

**Respuesta:** `200 OK`
```json
{
  "status": "ok",
  "service": "ms-funcionario"
}
```

## Códigos de Respuesta HTTP

| Código | Descripción |
|--------|-------------|
| 200 | OK - Operación exitosa |
| 201 | Created - Recurso creado exitosamente |
| 204 | No Content - Eliminación exitosa |
| 400 | Bad Request - Datos de entrada inválidos |
| 404 | Not Found - Recurso no encontrado |
| 409 | Conflict - Conflicto (ej: RUT o email duplicado) |
| 500 | Internal Server Error - Error del servidor |

## Modelo de Datos

### Tabla: Funcionario

| Campo | Tipo | Descripción |
|-------|------|-------------|
| id_funcionario | int4 | Primary Key, autoincremental |
| rut_funcionario | varchar(20) | RUT único del funcionario |
| nombres | varchar(100) | Nombres del funcionario |
| apellido_paterno | varchar(100) | Apellido paterno |
| apellido_materno | varchar(100) | Apellido materno |
| celular | varchar(20) | Número de celular |
| telefono | varchar(20) | Teléfono fijo |
| email | varchar(100) | Correo electrónico |
| tallas_registradas | bool | Indica si tiene medidas registradas |
| direccion | varchar(255) | Dirección completa |
| fecha_creación | date | Fecha de creación del registro |
| fecha_modificación | date | Fecha de última modificación |
| id_genero | int4 | FK a tabla Genero |
| id_medidas | int4 | FK a tabla Medidas Funcionario |
| id_usuario | int4 | FK a tabla Usuario (creador) |
| id_estado | int4 | FK a tabla Estado |
| id_sucursal | int4 | FK a tabla Sucursal |
| id_empresa_cliente | int4 | FK a tabla Empresa |
| id_segmento | int4 | FK a tabla Segmento |
| id_cargo | int4 | FK a tabla cargo |

### Tabla: Medidas Funcionario

| Campo | Tipo | Descripción |
|-------|------|-------------|
| id_medidas | int4 | Primary Key, autoincremental |
| estatura_m | numeric(5,2) | Estatura en metros |
| pecho_cm | numeric(5,2) | Medida de pecho en centímetros |
| cintura_cm | numeric(5,2) | Medida de cintura en centímetros |
| cadera_cm | numeric(5,2) | Medida de cadera en centímetros |
| manga_cm | numeric(5,2) | Medida de manga en centímetros |
| fecha_inicio | date | Fecha de inicio de vigencia |
| fecha_fin | date | Fecha de fin de vigencia (NULL si está activa) |

## Reglas de Negocio

### Funcionarios

1. **Unicidad:**
   - El RUT debe ser único en el sistema
   - El email debe ser único en el sistema

2. **Validaciones:**
   - RUT, nombres, apellido paterno y email son obligatorios
   - Email debe tener formato válido

3. **Estados:**
   - Estado 1: Activo
   - Estado 2: Inactivo
   - Por defecto, los funcionarios se crean en estado activo

4. **Relaciones:**
   - Un funcionario debe pertenecer a una empresa cliente
   - Puede estar asignado a una sucursal específica
   - Puede estar asociado a un segmento
   - Puede tener un cargo asignado

### Medidas Corporales

1. **Versionamiento:**
   - Se mantiene un historial de medidas
   - Solo puede haber una medida activa (sin `fecha_fin`) por funcionario
   - Al crear nuevas medidas, las anteriores se cierran automáticamente

2. **Validaciones:**
   - La estatura es obligatoria al crear medidas
   - Las demás medidas son opcionales

3. **Actualización:**
   - Solo se pueden actualizar las medidas activas
   - Se actualizan parcialmente (solo campos proporcionados)

## Configuración

### Variables de Entorno

El servicio utiliza las siguientes variables de entorno (definidas en `config/config.go`):

| Variable | Descripción | Default |
|----------|-------------|---------|
| SERVER_PORT | Puerto del servidor | 8080 |
| DB_HOST | Host de PostgreSQL | localhost |
| DB_PORT | Puerto de PostgreSQL | 5432 |
| DB_USER | Usuario de base de datos | postgres |
| DB_PASSWORD | Contraseña de base de datos | - |
| DB_NAME | Nombre de la base de datos | funcionarios_db |

### Ejemplo docker-compose.yml

```yaml
version: '3.8'

services:
  postgres:
    image: postgres:15-alpine
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: postgres
      POSTGRES_DB: funcionarios_db
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data

  ms-funcionario:
    build: .
    ports:
      - "8080:8080"
    environment:
      SERVER_PORT: 8080
      DB_HOST: postgres
      DB_PORT: 5432
      DB_USER: postgres
      DB_PASSWORD: postgres
      DB_NAME: funcionarios_db
    depends_on:
      - postgres

volumes:
  postgres_data:
```

## Comandos de Desarrollo

### Usando Makefile

```bash
# Construir la aplicación
make build

# Ejecutar la aplicación
make run

# Ejecutar tests
make test

# Levantar servicios con Docker Compose
make docker-up

# Detener servicios
make docker-down

# Limpiar binarios
make clean
```

### Comandos Docker

```bash
# Construir imagen
docker build -t ms-funcionario:latest .

# Ejecutar contenedor
docker run -p 8080:8080 \
  -e DB_HOST=postgres \
  -e DB_USER=postgres \
  -e DB_PASSWORD=postgres \
  ms-funcionario:latest

# Usar docker-compose
docker-compose up -d
docker-compose down
```

### Comandos Go

```bash
# Descargar dependencias
go mod download

# Actualizar dependencias
go mod tidy

# Ejecutar aplicación
go run cmd/api/main.go

# Compilar
go build -o bin/ms-funcionario cmd/api/main.go

# Ejecutar tests
go test ./...

# Ver cobertura
go test -cover ./...
```

## Integración con el Ecosistema

Este microservicio se integra con otras entidades del sistema:

### Dependencias de Datos

- **Empresa**: Los funcionarios pertenecen a empresas cliente
- **Sucursal**: Pueden estar asignados a sucursales específicas
- **Segmento**: Clasificación de funcionarios por tipo de uniforme
- **Estado**: Control de estado activo/inactivo
- **Genero**: Clasificación de género para uniformes
- **Cargo**: Puesto de trabajo del funcionario
- **Usuario**: Auditoría de creación/modificación

### Servicios Relacionados

- **ms-uniforme**: Consulta funcionarios para asignación de uniformes
- **ms-peticion**: Utiliza datos de funcionarios y medidas para solicitudes
- **ms-empresa**: Valida existencia de empresas cliente
- **ms-sucursal**: Valida asignaciones de sucursales

## Mejores Prácticas Implementadas

1. **Arquitectura Hexagonal**: Separación clara de responsabilidades
2. **Dependency Injection**: Facilita testing y mantenibilidad
3. **Repository Pattern**: Abstracción de la capa de datos
4. **DTO Pattern**: Separación de modelos de dominio y transferencia
5. **Error Handling**: Manejo consistente de errores con tipos específicos
6. **Paginación**: Endpoints que retornan listas soportan paginación
7. **Versionamiento de API**: Prefijo `/api/v1` para versionamiento
8. **Health Checks**: Endpoint de salud para orquestadores
9. **Logging**: Registro de eventos importantes
10. **Configuración por Entorno**: Variables de entorno para diferentes ambientes

## Consideraciones de Seguridad

- **Validación de Entrada**: Todos los DTOs tienen validaciones
- **SQL Injection**: GORM protege contra inyecciones SQL
- **CORS**: Configurar según necesidades (no implementado por defecto)
- **Rate Limiting**: Considerar implementar para producción
- **Autenticación/Autorización**: No implementada (debe agregarse según arquitectura general)

## Roadmap y Mejoras Futuras

1. **Autenticación y Autorización**: Integración con JWT/OAuth2
2. **Validación de RUT**: Algoritmo de validación chileno
3. **Caché**: Redis para consultas frecuentes
4. **Eventos**: Event sourcing para auditoría completa
5. **Métricas**: Prometheus/Grafana para monitoreo
6. **API Documentation**: Swagger/OpenAPI
7. **Tests**: Cobertura de tests unitarios e integración
8. **CI/CD**: Pipeline automatizado
9. **Versionamiento de Medidas**: Mejora en gestión de historial
10. **Búsqueda Avanzada**: Elasticsearch para búsquedas complejas

## Contacto y Soporte

Para preguntas, problemas o contribuciones, contactar al equipo de desarrollo del proyecto Sistal.

---

**Versión:** 1.0.0  
**Última Actualización:** 5 de enero de 2026  
**Autor:** Sistema de Gestión de Uniformes Sistal
