# Ejemplos de API - Microservicio Funcionarios

Esta guía contiene ejemplos prácticos de todas las operaciones disponibles en la API.

## Configuración Base

- **URL Base**: `http://localhost:8080`
- **Content-Type**: `application/json`

---

## 1. GESTIÓN DE FUNCIONARIOS

### 1.1 Crear Funcionario

**Request:**
```bash
curl -X POST http://localhost:8080/api/v1/funcionarios \
  -H "Content-Type: application/json" \
  -d '{
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
  }'
```

**Response (201 Created):**
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

### 1.2 Obtener Funcionario por ID

**Request:**
```bash
curl http://localhost:8080/api/v1/funcionarios/1
```

**Response (200 OK):**
```json
{
  "id_funcionario": 1,
  "rut_funcionario": "12345678-9",
  "nombres": "Juan Pablo",
  "apellido_paterno": "González",
  "apellido_materno": "Pérez",
  "email": "juan.gonzalez@empresa.cl",
  "tallas_registradas": true,
  "id_estado": 1
}
```

### 1.3 Listar Todos los Funcionarios

**Request:**
```bash
curl http://localhost:8080/api/v1/funcionarios
```

**Response (200 OK):**
```json
[
  {
    "id_funcionario": 1,
    "rut_funcionario": "12345678-9",
    "nombres": "Juan Pablo",
    "apellido_paterno": "González",
    ...
  },
  {
    "id_funcionario": 2,
    "rut_funcionario": "98765432-1",
    "nombres": "María José",
    "apellido_paterno": "Silva",
    ...
  }
]
```

### 1.4 Actualizar Funcionario

**Request:**
```bash
curl -X PUT http://localhost:8080/api/v1/funcionarios/1 \
  -H "Content-Type: application/json" \
  -d '{
    "rut_funcionario": "12345678-9",
    "nombres": "Juan Pablo Andrés",
    "apellido_paterno": "González",
    "apellido_materno": "Pérez",
    "celular": "+56987654321",
    "email": "juan.gonzalez@empresa.cl",
    "direccion": "Nueva Dirección 456",
    "id_estado": 1
  }'
```

**Response (200 OK):**
```json
{
  "id_funcionario": 1,
  "nombres": "Juan Pablo Andrés",
  "celular": "+56987654321",
  "direccion": "Nueva Dirección 456",
  "fecha_modificacion": "2026-01-05T12:30:00Z",
  ...
}
```

### 1.5 Eliminar Funcionario

**Request:**
```bash
curl -X DELETE http://localhost:8080/api/v1/funcionarios/1
```

**Response (204 No Content)**

---

## 2. BÚSQUEDAS Y FILTROS

### 2.1 Filtrar Funcionarios con Paginación

**Request:**
```bash
curl "http://localhost:8080/api/v1/funcionarios/filter?id_empresa_cliente=5&id_estado=1&tallas_registradas=true&limit=20&offset=0"
```

**Response (200 OK):**
```json
{
  "data": [
    {
      "id_funcionario": 1,
      "nombres": "Juan Pablo",
      ...
    },
    {
      "id_funcionario": 3,
      "nombres": "María José",
      ...
    }
  ],
  "total": 45,
  "limit": 20,
  "offset": 0,
  "total_pages": 3
}
```

### 2.2 Buscar por RUT

**Request:**
```bash
curl http://localhost:8080/api/v1/funcionarios/rut/12345678-9
```

**Response (200 OK):**
```json
{
  "id_funcionario": 1,
  "rut_funcionario": "12345678-9",
  "nombres": "Juan Pablo",
  ...
}
```

### 2.3 Listar Funcionarios por Empresa

**Request:**
```bash
curl http://localhost:8080/api/v1/funcionarios/empresa/5
```

**Response (200 OK):**
```json
[
  {
    "id_funcionario": 1,
    "id_empresa_cliente": 5,
    ...
  },
  {
    "id_funcionario": 2,
    "id_empresa_cliente": 5,
    ...
  }
]
```

### 2.4 Listar Funcionarios por Sucursal

**Request:**
```bash
curl http://localhost:8080/api/v1/funcionarios/sucursal/2
```

**Response (200 OK):** Array de funcionarios

### 2.5 Listar Funcionarios por Segmento

**Request:**
```bash
curl http://localhost:8080/api/v1/funcionarios/segmento/3
```

**Response (200 OK):** Array de funcionarios

---

## 3. GESTIÓN DE ESTADO

### 3.1 Activar Funcionario

**Request:**
```bash
curl -X PATCH http://localhost:8080/api/v1/funcionarios/1/activate
```

**Response (200 OK):**
```json
{
  "message": "funcionario activated successfully"
}
```

### 3.2 Desactivar Funcionario

**Request:**
```bash
curl -X PATCH http://localhost:8080/api/v1/funcionarios/1/deactivate
```

**Response (200 OK):**
```json
{
  "message": "funcionario deactivated successfully"
}
```

---

## 4. GESTIÓN DE MEDIDAS CORPORALES

### 4.1 Registrar Medidas

**Request:**
```bash
curl -X POST http://localhost:8080/api/v1/funcionarios/1/medidas \
  -H "Content-Type: application/json" \
  -d '{
    "estatura_m": 1.75,
    "pecho_cm": 95.5,
    "cintura_cm": 85.0,
    "cadera_cm": 98.0,
    "manga_cm": 62.5,
    "fecha_inicio": "2026-01-05T00:00:00Z"
  }'
```

**Response (201 Created):**
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

**Nota:** Si el funcionario ya tiene medidas activas, estas se cerran automáticamente (se les asigna `fecha_fin`).

### 4.2 Obtener Medidas Activas

**Request:**
```bash
curl http://localhost:8080/api/v1/funcionarios/1/medidas
```

**Response (200 OK):**
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

### 4.3 Actualizar Medidas Activas

**Request:**
```bash
curl -X PUT http://localhost:8080/api/v1/funcionarios/1/medidas \
  -H "Content-Type: application/json" \
  -d '{
    "pecho_cm": 96.0,
    "cintura_cm": 84.5
  }'
```

**Response (200 OK):**
```json
{
  "id_medidas": 1,
  "estatura_m": 1.75,
  "pecho_cm": 96.0,
  "cintura_cm": 84.5,
  "cadera_cm": 98.0,
  "manga_cm": 62.5,
  "fecha_inicio": "2026-01-05T00:00:00Z",
  "fecha_fin": null
}
```

**Nota:** Solo se actualizan los campos proporcionados.

### 4.4 Obtener Historial de Medidas

**Request:**
```bash
curl http://localhost:8080/api/v1/funcionarios/1/medidas/historial
```

**Response (200 OK):**
```json
[
  {
    "id_medidas": 2,
    "estatura_m": 1.75,
    "pecho_cm": 96.0,
    "cintura_cm": 84.5,
    "cadera_cm": 98.0,
    "manga_cm": 62.5,
    "fecha_inicio": "2026-01-05T00:00:00Z",
    "fecha_fin": null
  },
  {
    "id_medidas": 1,
    "estatura_m": 1.75,
    "pecho_cm": 95.5,
    "cintura_cm": 85.0,
    "cadera_cm": 98.0,
    "manga_cm": 62.5,
    "fecha_inicio": "2025-06-01T00:00:00Z",
    "fecha_fin": "2026-01-05T00:00:00Z"
  }
]
```

**Nota:** Ordenado por fecha de inicio descendente (más recientes primero).

---

## 5. HEALTH CHECK

### 5.1 Verificar Estado del Servicio

**Request:**
```bash
curl http://localhost:8080/health
```

**Response (200 OK):**
```json
{
  "status": "ok",
  "service": "ms-funcionario"
}
```

---

## 6. EJEMPLOS DE ERRORES

### 6.1 Funcionario No Encontrado (404)

**Request:**
```bash
curl http://localhost:8080/api/v1/funcionarios/999
```

**Response (404 Not Found):**
```json
{
  "error": "funcionario not found"
}
```

### 6.2 Datos Inválidos (400)

**Request:**
```bash
curl -X POST http://localhost:8080/api/v1/funcionarios \
  -H "Content-Type: application/json" \
  -d '{
    "nombres": "Juan"
  }'
```

**Response (400 Bad Request):**
```json
{
  "error": "Key: 'CreateFuncionarioRequest.RutFuncionario' Error:Field validation for 'RutFuncionario' failed on the 'required' tag"
}
```

### 6.3 RUT Duplicado (409)

**Request:**
```bash
curl -X POST http://localhost:8080/api/v1/funcionarios \
  -H "Content-Type: application/json" \
  -d '{
    "rut_funcionario": "12345678-9",
    "nombres": "Pedro",
    "apellido_paterno": "López",
    "email": "pedro@empresa.cl",
    "id_empresa_cliente": 5
  }'
```

**Response (409 Conflict):**
```json
{
  "error": "funcionario already exists: rut 12345678-9 already exists"
}
```

### 6.4 Medidas No Encontradas (404)

**Request:**
```bash
curl http://localhost:8080/api/v1/funcionarios/1/medidas
```

**Response (404 Not Found):**
```json
{
  "error": "medidas not found"
}
```

---

## 7. CASOS DE USO COMPLETOS

### 7.1 Flujo Completo: Crear Funcionario y Registrar Medidas

**Paso 1: Crear Funcionario**
```bash
curl -X POST http://localhost:8080/api/v1/funcionarios \
  -H "Content-Type: application/json" \
  -d '{
    "rut_funcionario": "11111111-1",
    "nombres": "Ana María",
    "apellido_paterno": "Rodríguez",
    "apellido_materno": "Soto",
    "email": "ana.rodriguez@empresa.cl",
    "celular": "+56911111111",
    "id_empresa_cliente": 5,
    "id_genero": 2
  }'
```

**Paso 2: Registrar Medidas**
```bash
curl -X POST http://localhost:8080/api/v1/funcionarios/1/medidas \
  -H "Content-Type: application/json" \
  -d '{
    "estatura_m": 1.65,
    "pecho_cm": 88.0,
    "cintura_cm": 70.0,
    "cadera_cm": 95.0,
    "manga_cm": 58.0
  }'
```

**Paso 3: Verificar Funcionario con Tallas Registradas**
```bash
curl http://localhost:8080/api/v1/funcionarios/1
```

### 7.2 Flujo: Actualizar Medidas de un Funcionario

**Paso 1: Obtener Medidas Actuales**
```bash
curl http://localhost:8080/api/v1/funcionarios/1/medidas
```

**Paso 2: Crear Nuevas Medidas (cierra las anteriores automáticamente)**
```bash
curl -X POST http://localhost:8080/api/v1/funcionarios/1/medidas \
  -H "Content-Type: application/json" \
  -d '{
    "estatura_m": 1.65,
    "pecho_cm": 90.0,
    "cintura_cm": 72.0,
    "cadera_cm": 96.0,
    "manga_cm": 58.5
  }'
```

**Paso 3: Ver Historial**
```bash
curl http://localhost:8080/api/v1/funcionarios/1/medidas/historial
```

### 7.3 Flujo: Búsqueda y Filtrado Avanzado

**Buscar funcionarios activos de una empresa con tallas registradas:**
```bash
curl "http://localhost:8080/api/v1/funcionarios/filter?id_empresa_cliente=5&id_estado=1&tallas_registradas=true&limit=50&offset=0"
```

**Buscar funcionarios de una sucursal específica:**
```bash
curl "http://localhost:8080/api/v1/funcionarios/filter?id_sucursal=2&limit=100"
```

**Buscar funcionarios por segmento y cargo:**
```bash
curl "http://localhost:8080/api/v1/funcionarios/filter?id_segmento=3&id_cargo=4"
```

---

## 8. TESTING CON POSTMAN

### Importar Colección

Puedes crear una colección de Postman con estos endpoints. Ejemplo de estructura:

```
Funcionarios API
├── Funcionarios
│   ├── Crear Funcionario (POST)
│   ├── Listar Funcionarios (GET)
│   ├── Obtener por ID (GET)
│   ├── Actualizar (PUT)
│   ├── Eliminar (DELETE)
│   ├── Buscar por RUT (GET)
│   ├── Filtrar (GET)
│   ├── Por Empresa (GET)
│   ├── Por Sucursal (GET)
│   ├── Por Segmento (GET)
│   ├── Activar (PATCH)
│   └── Desactivar (PATCH)
├── Medidas
│   ├── Registrar Medidas (POST)
│   ├── Obtener Activas (GET)
│   ├── Actualizar (PUT)
│   └── Historial (GET)
└── Health
    └── Health Check (GET)
```

---

## 9. NOTAS IMPORTANTES

1. **IDs Autoincrementales**: Los IDs de funcionarios y medidas son generados automáticamente por la base de datos.

2. **Fechas**: Las fechas se manejan automáticamente:
   - `fecha_creacion`: Se establece al crear
   - `fecha_modificacion`: Se actualiza en cada modificación
   - `fecha_inicio` (medidas): Por defecto es la fecha actual
   - `fecha_fin` (medidas): NULL indica medidas activas

3. **Estados**: 
   - Por defecto, los funcionarios se crean con `id_estado = 1` (Activo)
   - Usar endpoints de activar/desactivar para cambiar estado

4. **Validaciones**:
   - RUT y email deben ser únicos
   - Email debe tener formato válido
   - Nombres, apellido paterno, RUT y email son obligatorios

5. **Medidas**:
   - Solo puede haber un conjunto de medidas activas por funcionario
   - Al crear nuevas medidas, las anteriores se cierran automáticamente
   - El campo `tallas_registradas` se actualiza automáticamente

---

**Última Actualización:** 5 de enero de 2026
