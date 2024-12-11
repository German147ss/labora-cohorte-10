# API REST con Go y DDD (Domain-Driven Design)

Este proyecto implementa una API REST siguiendo los principios de Domain-Driven Design (DDD) en Go, demostrando una arquitectura limpia y mantenible para gestionar usuarios.

## Introducción a DDD

Domain-Driven Design es un enfoque de desarrollo de software que:
- Prioriza el dominio del negocio y su lógica
- Basa el diseño en un modelo del dominio
- Establece límites claros entre diferentes partes del software

## Estructura del Proyecto

```bash
├── cmd
│   └── api
│       └── main.go           # Punto de entrada de la aplicación
├── internal
│   ├── domain               # Capa de dominio
│   │   ├── user
│   │   │   ├── entity.go    # Entidades y reglas de negocio
│   │   │   ├── repository.go # Interfaces del repositorio
│   │   │   └── service.go    # Lógica de negocio
│   ├── infrastructure       # Capa de infraestructura
│   │   ├── http
│   │   │   ├── handler      # Manejadores HTTP
│   │   │   │   └── user_handler.go
│   │   │   └── router       # Configuración de rutas
│   │   │       └── router.go
│   │   └── persistence      # Implementaciones de persistencia
│   │       └── memory
│   │           └── user_repository.go
```

### Explicación de las Capas

1. **Domain Layer** (`internal/domain`)
   - Contiene las entidades centrales del negocio
   - Define interfaces y contratos
   - Implementa la lógica de negocio core
   - Es independiente de frameworks y tecnologías externas

2. **Infrastructure Layer** (`internal/infrastructure`)
   - Implementa las interfaces definidas en el dominio
   - Maneja la comunicación HTTP
   - Gestiona la persistencia de datos
   - Integra con servicios externos

3. **Entry Point** (`cmd/api`)
   - Inicializa y configura la aplicación
   - Maneja la inyección de dependencias
   - Arranca el servidor HTTP

## Endpoints de la API

| Método | Ruta | Descripción |
|--------|------|-------------|
| GET | `/users` | Obtiene todos los usuarios |
| GET | `/users/{id}` | Obtiene un usuario por ID |
| POST | `/users` | Crea un nuevo usuario |
| PUT | `/users/{id}` | Actualiza un usuario existente |
| DELETE | `/users/{id}` | Elimina un usuario |

## Beneficios de esta Arquitectura

1. **Separación de Responsabilidades**
   - Cada capa tiene un propósito específico
   - Facilita el mantenimiento y las pruebas
   - Reduce el acoplamiento entre componentes

2. **Escalabilidad**
   - Fácil de agregar nuevas características
   - Simple de modificar implementaciones existentes
   - Preparado para crecer con el negocio

3. **Testabilidad**
   - Arquitectura orientada a pruebas
   - Fácil mock de dependencias
   - Pruebas unitarias más limpias

4. **Mantenibilidad**
   - Código organizado y predecible
   - Fácil de entender y modificar
   - Documentación implícita en la estructura


### Modulos en Go

```bash
go mod init api-rest-postgresql
go mod tidy
```

- Que es go mod?

    - Es un sistema de gestión de dependencias para Go
    - Permite a los desarrolladores especificar y gestionar las dependencias de sus proyectos

- Que es go mod init?

    - Inicializa un nuevo módulo Go
    - Crea un archivo go.mod en el directorio actual
    - Agrega el módulo al sistema de módulos de Go

- Que hace go mod tidy?

    - Agrega los modulos necesarios para el proyecto
    - Elimina los modulos que no se usan
    - Actualiza el archivo go.mod con las dependencias necesarias


### Instalar PostgreSQL

```bash
brew install postgresql
```

### Instalar pq en go

```bash
go get github.com/lib/pq
```

- Que es pq?

    - Es una librería para interactuar con PostgreSQL desde Go
    - Proporciona funciones para ejecutar consultas SQL y manejar errores
    - Facilita la manipulación de datos en la base de datos
