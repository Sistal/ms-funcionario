# Microservicio de Funcionarios

Microservicio RESTful para la gestión integral de funcionarios en un sistema de control de uniformes empresariales. Desarrollado en Go con Gin Framework y PostgreSQL.

## 🚀 Características Principales

- ✅ **CRUD completo** de funcionarios
- 📏 **Gestión de medidas corporales** con historial
- 🔍 **Búsquedas avanzadas** con filtros múltiples
- 📊 **Paginación** en listados
- 🏢 **Consultas por empresa, sucursal y segmento**
- 🔄 **Activación/desactivación** de funcionarios
- 🏗️ **Arquitectura hexagonal** (Ports & Adapters)
- 🐳 **Dockerizado** con Docker Compose
- 📖 **Swagger UI** integrado para documentación interactiva

## 📋 Tecnologías

- **Go** 1.23.0+
- **Gin Web Framework** 1.11.0
- **GORM** 1.31.1 (ORM)
- **PostgreSQL** con driver pgx/v5
- **Docker & Docker Compose**
- **Swagger/OpenAPI** 3.0 para documentación

## 📁 Estructura del Proyecto

```
ms-funcionario/
├── cmd/api/main.go                      # Punto de entrada
├── config/config.go                     # Configuración
├── internal/
│   ├── domain/funcionario/              # Entidades y contratos
│   │   ├── funcionario.go               # Modelos de dominio
│   │   └── repository.go                # Interfaces
│   ├── application/services/            # Lógica de negocio
│   │   └── funcionario_service.go
│   ├── infrastructure/                  # Implementaciones
│   │   ├── database/postgres.go
│   │   └── repository/funcionario_repository.go
│   └── interfaces/                      # Adaptadores HTTP
│       ├── dto/funcionario_dto.go
│       └── http/
│           ├── handler/funcionario_handler.go
│           └── router/router.go
├── docker-compose.yml
├── Dockerfile
└── project_framework.md                 # Documentación completa
```

## 🎯 API Endpoints

### Funcionarios

| Método | Endpoint | Descripción |
|--------|----------|-------------|
| POST | `/api/v1/funcionarios` | Crear funcionario |
| GET | `/api/v1/funcionarios` | Listar todos |
| GET | `/api/v1/funcionarios/{id}` | Obtener por ID |
| PUT | `/api/v1/funcionarios/{id}` | Actualizar |
| DELETE | `/api/v1/funcionarios/{id}` | Eliminar |
| GET | `/api/v1/funcionarios/filter` | Filtrar con paginación |
| GET | `/api/v1/funcionarios/rut/{rut}` | Buscar por RUT |
| GET | `/api/v1/funcionarios/empresa/{id}` | Listar por empresa |
| GET | `/api/v1/funcionarios/sucursal/{id}` | Listar por sucursal |
| GET | `/api/v1/funcionarios/segmento/{id}` | Listar por segmento |
| PATCH | `/api/v1/funcionarios/{id}/activate` | Activar |
| PATCH | `/api/v1/funcionarios/{id}/deactivate` | Desactivar |

### Medidas Corporales

| Método | Endpoint | Descripción |
|--------|----------|-------------|
| POST | `/api/v1/funcionarios/{id}/medidas` | Registrar medidas |
| GET | `/api/v1/funcionarios/{id}/medidas` | Obtener medidas activas |
| PUT | `/api/v1/funcionarios/{id}/medidas` | Actualizar medidas |
| GET | `/api/v1/funcionarios/{id}/medidas/historial` | Historial de medidas |

### Health Check

| Método | Endpoint | Descripción |
|--------|----------|-------------|
| GET | `/health` | Estado del servicio |
| GET | `/swagger/index.html` | Documentación Swagger UI |

## 📖 Documentación Interactiva

**Swagger UI:** Una vez que el servicio esté ejecutándose, accede a la documentación interactiva en:

🔗 **http://localhost:8080/swagger/index.html**

Desde Swagger UI puedes:
- Ver todos los endpoints disponibles
- Probar cada endpoint directamente desde el navegador
- Ver ejemplos de request y response
- Consultar los esquemas de datos

Para más detalles, consulta [SWAGGER.md](SWAGGER.md)

## 🚀 Inicio Rápido

### Con Docker Compose (Recomendado)

```bash
# Levantar servicios
docker-compose up -d

# Ver logs
docker-compose logs -f ms-funcionario

# Detener servicios
docker-compose down
```

La aplicación estará disponible en `http://localhost:8080`

### Sin Docker (Desarrollo Local)

1. **Instalar dependencias:**
```bash
go mod download
```

2. **Configurar PostgreSQL** (debe estar ejecutándose)

3. **Configurar variables de entorno:**
```bash
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=postgres
export DB_NAME=funcionarios_db
export SERVER_PORT=8080
```

4. **Ejecutar la aplicación:**
```bash
go run cmd/api/main.go
```

## 💡 Ejemplos de Uso

### Crear Funcionario

```bash
curl -X POST http://localhost:8080/api/v1/funcionarios \
  -H "Content-Type: application/json" \
  -d '{
    "rut_funcionario": "12345678-9",
    "nombres": "Juan Pablo",
    "apellido_paterno": "González",
    "apellido_materno": "Pérez",
    "email": "juan.gonzalez@empresa.cl",
    "celular": "+56912345678",
    "direccion": "Av. Principal 123",
    "id_empresa_cliente": 5,
    "id_genero": 1
  }'
```

### Listar con Filtros y Paginación

```bash
curl "http://localhost:8080/api/v1/funcionarios/filter?id_empresa_cliente=5&id_estado=1&limit=20&offset=0"
```

### Buscar por RUT

```bash
curl http://localhost:8080/api/v1/funcionarios/rut/12345678-9
```

### Registrar Medidas

```bash
curl -X POST http://localhost:8080/api/v1/funcionarios/1/medidas \
  -H "Content-Type: application/json" \
  -d '{
    "estatura_m": 1.75,
    "pecho_cm": 95.5,
    "cintura_cm": 85.0,
    "cadera_cm": 98.0,
    "manga_cm": 62.5
  }'
```

### Obtener Medidas Activas

```bash
curl http://localhost:8080/api/v1/funcionarios/1/medidas
```

### Activar/Desactivar Funcionario

```bash
# Activar
curl -X PATCH http://localhost:8080/api/v1/funcionarios/1/activate

# Desactivar
curl -X PATCH http://localhost:8080/api/v1/funcionarios/1/deactivate
```

## ⚙️ Variables de Entorno

| Variable | Descripción | Valor por Defecto |
|----------|-------------|-------------------|
| SERVER_PORT | Puerto del servidor | 8080 |
| DB_HOST | Host de PostgreSQL | localhost |
| DB_PORT | Puerto de PostgreSQL | 5432 |
| DB_USER | Usuario de base de datos | postgres |
| DB_PASSWORD | Contraseña de base de datos | - |
| DB_NAME | Nombre de la base de datos | funcionarios_db |

## 🗄️ Modelo de Datos

### Funcionario

- **id_funcionario**: ID único (autoincremental)
- **rut_funcionario**: RUT único del funcionario
- **nombres**: Nombres completos
- **apellido_paterno**: Apellido paterno
- **apellido_materno**: Apellido materno
- **celular**: Número de celular
- **telefono**: Teléfono fijo
- **email**: Correo electrónico único
- **tallas_registradas**: Indica si tiene medidas
- **direccion**: Dirección completa
- **Relaciones**: empresa, sucursal, segmento, cargo, género, estado, medidas

### Medidas Funcionario

- **id_medidas**: ID único (autoincremental)
- **estatura_m**: Estatura en metros
- **pecho_cm**: Medida de pecho en cm
- **cintura_cm**: Medida de cintura en cm
- **cadera_cm**: Medida de cadera en cm
- **manga_cm**: Medida de manga en cm
- **fecha_inicio**: Fecha de inicio de vigencia
- **fecha_fin**: Fecha de fin (NULL si está activa)

## 🧪 Testing

```bash
# Ejecutar todos los tests
go test ./...

# Con cobertura
go test -cover ./...

# Tests específicos
go test ./internal/application/services/...
```

## 📦 Comandos Make

```bash
make build        # Compilar la aplicación
make run          # Ejecutar la aplicación
make test         # Ejecutar tests
make docker-up    # Levantar con Docker Compose
make docker-down  # Detener Docker Compose
make clean        # Limpiar binarios
```

## 🏗️ Arquitectura

El proyecto sigue **Arquitectura Hexagonal** con separación en capas:

- **Domain**: Entidades y reglas de negocio
- **Application**: Casos de uso y servicios
- **Infrastructure**: Implementaciones (DB, repositorios)
- **Interfaces**: Adaptadores HTTP (handlers, DTOs, router)

## 📚 Documentación Completa

Para documentación técnica detallada, consulta [project_framework.md](project_framework.md)

## 🔒 Seguridad

- ✅ Validación de entrada con binding tags
- ✅ Protección contra SQL injection (GORM)
- ⚠️ Autenticación/Autorización: Pendiente de implementar
- ⚠️ CORS: Configurar según necesidades

## 🚦 Health Check

```bash
curl http://localhost:8080/health

# Respuesta:
{
  "status": "ok",
  "service": "ms-funcionario"
}
```

## 📝 Licencia

Este proyecto es parte del sistema Sistal de gestión de uniformes empresariales.

## 👥 Contribución

Para contribuir al proyecto, por favor seguir las guías de estilo de Go y crear pull requests detallados.

---

**Versión:** 1.0.0  
**Última Actualización:** 5 de enero de 2026
| DB_PORT | Puerto de PostgreSQL | 5432 |
| DB_USER | Usuario de la base de datos | postgres |
| DB_PASSWORD | Contraseña de la base de datos | postgres |
| DB_NAME | Nombre de la base de datos | funcionarios_db |
| DB_SSLMODE | Modo SSL de PostgreSQL | disable |

## Estructura de la Base de Datos

### Tabla: funcionarios

| Campo | Tipo | Descripción |
|-------|------|-------------|
| id | UUID | Identificador único (PK) |
| nombre | VARCHAR(100) | Nombre del funcionario |
| apellido | VARCHAR(100) | Apellido del funcionario |
| email | VARCHAR(100) | Email único del funcionario |
| cargo | VARCHAR(100) | Cargo del funcionario |
| activo | BOOLEAN | Estado activo/inactivo |
| created_at | TIMESTAMP | Fecha de creación |
| updated_at | TIMESTAMP | Fecha de última actualización |

## Detener la Aplicación

```bash
docker-compose down
```

Para eliminar también los volúmenes:
```bash
docker-compose down -v
```

## Licencia

MIT