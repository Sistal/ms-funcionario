# ms-funcionario

Microservicio de gestión de funcionarios desarrollado con Golang, utilizando el framework Gin, base de datos PostgreSQL con GORM, y siguiendo una arquitectura hexagonal.

## Características

- Framework web: **Gin**
- ORM: **GORM**
- Base de datos: **PostgreSQL**
- Arquitectura: **Hexagonal (Ports & Adapters)**
- Contenedores: **Docker & Docker Compose**

## Arquitectura Hexagonal

El proyecto está organizado siguiendo los principios de la arquitectura hexagonal:

```
ms-funcionario/
├── cmd/
│   └── api/              # Punto de entrada de la aplicación
├── config/               # Configuración de la aplicación
├── internal/
│   ├── domain/           # Capa de dominio (entidades, interfaces)
│   │   └── funcionario/
│   ├── application/      # Capa de aplicación (casos de uso, servicios)
│   │   └── services/
│   ├── infrastructure/   # Capa de infraestructura (implementaciones)
│   │   ├── database/
│   │   └── repository/
│   └── interfaces/       # Capa de interfaces (HTTP, DTOs)
│       ├── dto/
│       └── http/
│           ├── handler/
│           └── router/
├── docker-compose.yml
├── Dockerfile
└── README.md
```

## Requisitos Previos

- Go 1.21 o superior
- Docker y Docker Compose
- PostgreSQL 15 (si ejecutas localmente sin Docker)

## Instalación y Ejecución

### Opción 1: Con Docker Compose (Recomendado)

1. Clonar el repositorio:
```bash
git clone https://github.com/Sistal/ms-funcionario.git
cd ms-funcionario
```

2. Construir y ejecutar con Docker Compose:
```bash
docker-compose up --build
```

La aplicación estará disponible en `http://localhost:8080`

### Opción 2: Ejecución Local

1. Instalar dependencias:
```bash
go mod download
```

2. Configurar variables de entorno (copiar .env.example a .env y ajustar):
```bash
cp .env.example .env
```

3. Asegurarse de que PostgreSQL está ejecutándose

4. Ejecutar la aplicación:
```bash
go run cmd/api/main.go
```

## API Endpoints

### Health Check
- `GET /health` - Verificar el estado del servicio

### Funcionarios

#### Crear Funcionario
```bash
POST /api/v1/funcionarios
Content-Type: application/json

{
  "nombre": "Juan",
  "apellido": "Pérez",
  "email": "juan.perez@example.com",
  "cargo": "Desarrollador"
}
```

#### Obtener Todos los Funcionarios
```bash
GET /api/v1/funcionarios
```

#### Obtener Funcionario por ID
```bash
GET /api/v1/funcionarios/:id
```

#### Actualizar Funcionario
```bash
PUT /api/v1/funcionarios/:id
Content-Type: application/json

{
  "nombre": "Juan",
  "apellido": "Pérez",
  "email": "juan.perez@example.com",
  "cargo": "Senior Developer",
  "activo": true
}
```

#### Eliminar Funcionario
```bash
DELETE /api/v1/funcionarios/:id
```

## Ejemplos de Uso con cURL

### Crear un funcionario:
```bash
curl -X POST http://localhost:8080/api/v1/funcionarios \
  -H "Content-Type: application/json" \
  -d '{
    "nombre": "María",
    "apellido": "García",
    "email": "maria.garcia@example.com",
    "cargo": "Gerente"
  }'
```

### Listar todos los funcionarios:
```bash
curl http://localhost:8080/api/v1/funcionarios
```

### Obtener un funcionario específico:
```bash
curl http://localhost:8080/api/v1/funcionarios/{id}
```

### Actualizar un funcionario:
```bash
curl -X PUT http://localhost:8080/api/v1/funcionarios/{id} \
  -H "Content-Type: application/json" \
  -d '{
    "nombre": "María",
    "apellido": "García",
    "email": "maria.garcia@example.com",
    "cargo": "Gerente Senior",
    "activo": true
  }'
```

### Eliminar un funcionario:
```bash
curl -X DELETE http://localhost:8080/api/v1/funcionarios/{id}
```

## Variables de Entorno

| Variable | Descripción | Valor por Defecto |
|----------|-------------|-------------------|
| SERVER_PORT | Puerto del servidor | 8080 |
| DB_HOST | Host de PostgreSQL | localhost |
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