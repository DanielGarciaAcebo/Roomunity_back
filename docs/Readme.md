## Structure


```bash
home-manager/
├── cmd/
│   └── api/                     # Punto de entrada del servicio.
│       └── main.go
├── internal/
│   ├── calendar/                # Módulo: Gestión del calendario.
│   │   ├── app/                 # Casos de uso.
│   │   │   ├── create_event.go
│   │   │   ├── update_event.go
│   │   │   └── list_events.go
│   │   ├── domains/             # Lógica de negocio.
│   │   │   └── event.go
│   │   ├── interfaces/          # Adaptadores (HTTP, DB, etc.).
│   │   │   ├── http/
│   │   │   └── database/
│   │   └── pkg/                 # Librerías específicas del módulo.
│   ├── tasks/                   # Módulo: Gestión de tareas comunes.
│   │   ├── app/
│   │   ├── domains/
│   │   ├── interfaces/
│   │   └── pkg/
│   ├── shopping/                # Módulo: Lista de compras.
│   │   ├── app/
│   │   ├── domains/
│   │   ├── interfaces/
│   │   └── pkg/
│   ├── health/                  # Módulo: Salud y alimentación.
│   │   ├── app/
│   │   ├── domains/
│   │   ├── interfaces/
│   │   └── pkg/
│   └── shared/                  # Código compartido entre módulos.
│       ├── logger/              # Logging centralizado.
│       ├── events/              # Publicación/suscripción de eventos internos.
│       └── validation/          # Validaciones comunes.
├── config/                      # Configuración del proyecto.
│   ├── config.yaml
│   └── config.go
├── docs/                        # Documentación del proyecto.
├── pkg/                         # Librerías reutilizables (opcional).
├── test/                        # Pruebas de integración y e2e.
└── build/                       # Infraestructura para despliegue (Docker, Makefiles).

```