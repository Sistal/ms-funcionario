# Swagger UI - Documentación Interactiva de la API

## 📖 Acceso a Swagger UI

Una vez que el servicio esté ejecutándose, puedes acceder a la documentación interactiva de Swagger en:

**URL:** `http://localhost:8080/swagger/index.html`

## 🚀 Cómo Usar Swagger UI

### 1. Iniciar el Servicio

```powershell
# Ejecutar el servicio
go run cmd/api/main.go
```

### 2. Abrir Swagger UI

Abre tu navegador y visita:
```
http://localhost:8080/swagger/index.html
```

### 3. Explorar los Endpoints

En Swagger UI podrás:

- ✅ Ver todos los endpoints disponibles organizados por tags
- ✅ Ver los parámetros requeridos para cada endpoint
- ✅ Ver ejemplos de request y response
- ✅ **Probar los endpoints directamente** desde el navegador
- ✅ Ver los códigos de respuesta HTTP
- ✅ Ver los esquemas de datos (DTOs)

### 4. Probar un Endpoint

1. Haz clic en un endpoint (ej: `POST /api/v1/funcionarios`)
2. Haz clic en el botón **"Try it out"**
3. Edita el JSON de ejemplo con tus datos
4. Haz clic en **"Execute"**
5. Verás la respuesta del servidor en tiempo real

## 📋 Endpoints Disponibles

### Funcionarios
- `POST /api/v1/funcionarios` - Crear funcionario
- `GET /api/v1/funcionarios` - Listar todos
- `GET /api/v1/funcionarios/{id}` - Obtener por ID
- `PUT /api/v1/funcionarios/{id}` - Actualizar
- `DELETE /api/v1/funcionarios/{id}` - Eliminar
- `GET /api/v1/funcionarios/filter` - Filtrar con paginación
- `GET /api/v1/funcionarios/rut/{rut}` - Buscar por RUT
- `GET /api/v1/funcionarios/empresa/{id_empresa}` - Por empresa
- `GET /api/v1/funcionarios/sucursal/{id_sucursal}` - Por sucursal
- `GET /api/v1/funcionarios/segmento/{id_segmento}` - Por segmento
- `PATCH /api/v1/funcionarios/{id}/activate` - Activar
- `PATCH /api/v1/funcionarios/{id}/deactivate` - Desactivar

### Medidas
- `POST /api/v1/funcionarios/{id}/medidas` - Registrar medidas
- `GET /api/v1/funcionarios/{id}/medidas` - Obtener activas
- `PUT /api/v1/funcionarios/{id}/medidas` - Actualizar
- `GET /api/v1/funcionarios/{id}/medidas/historial` - Historial

### Health
- `GET /health` - Health check

## 🔄 Regenerar Documentación

Si modificas los comentarios de documentación en los handlers, regenera Swagger:

```powershell
swag init -g cmd/api/main.go -o docs
```

O si swag no está en el PATH:

```powershell
$env:PATH += ";$env:USERPROFILE\go\bin"
swag init -g cmd/api/main.go -o docs
```

## 📝 Formato de Anotaciones Swagger

Las anotaciones en el código siguen este formato:

```go
// @Summary Breve descripción
// @Description Descripción detallada
// @Tags nombre-del-tag
// @Accept json
// @Produce json
// @Param nombre path/query/body tipo requerido "descripción"
// @Success 200 {object} TipoRespuesta
// @Failure 400 {object} map[string]string
// @Router /ruta [metodo]
```

## 🎨 Ejemplo de Uso en Swagger UI

### Crear un Funcionario

1. Busca `POST /api/v1/funcionarios`
2. Click en "Try it out"
3. Edita el JSON:

```json
{
  "rut_funcionario": "12345678-9",
  "nombres": "Juan Pablo",
  "apellido_paterno": "González",
  "apellido_materno": "Pérez",
  "email": "juan.gonzalez@empresa.cl",
  "celular": "+56912345678",
  "id_empresa_cliente": 1,
  "id_genero": 1
}
```

4. Click en "Execute"
5. Verás la respuesta con el funcionario creado

## 📦 Archivos Generados

Swagger genera automáticamente estos archivos en el directorio `docs/`:

- `docs.go` - Código Go con la documentación embebida
- `swagger.json` - Especificación OpenAPI en JSON
- `swagger.yaml` - Especificación OpenAPI en YAML

**Nota:** Estos archivos se generan automáticamente, no los edites manualmente.

## 🌐 Acceso desde Otros Dispositivos

Si quieres acceder a Swagger desde otro dispositivo en tu red local:

1. Obtén tu IP local:
```powershell
ipconfig
```

2. Accede desde otro dispositivo:
```
http://TU_IP:8080/swagger/index.html
```

## ⚙️ Configuración Avanzada

### Cambiar el Host

En `cmd/api/main.go`, modifica la anotación:

```go
// @host localhost:8080  // Cambiar a tu host
```

### Agregar Autenticación (Futuro)

Cuando implementes autenticación, agrega:

```go
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
```

---

**Documentación generada automáticamente con Swag**  
**Última actualización:** 5 de enero de 2026
