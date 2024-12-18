
# Kredit

Service kredit
## Link

- [ERD](https://drawsql.app/teams/team-1756/diagrams/kredit)
- [Postman Api](https://crimson-crater-385688.postman.co/workspace/Open~b2b49bd4-fdb3-472f-ae87-d815a933728a/collection/27541101-9d6c7a2b-bb01-4afd-a537-a78a8d7b6110?action=share&creator=27541101&active-environment=27541101-2db38100-c936-4952-887b-cc478f8892ec)
- [Docker Hub](https://hub.docker.com/repository/docker/nxrfandi/kredit-plus/general)

## folder structur

```plaintext
├── kredit/                   # Contains business logic or modules specific to the "kredit" feature
├── cmd/                      # Contains application entry points or command-line commands
│   └── server.go             # Main server setup and initialization
├── doc/                      # Documentation files (e.g., API specs, design notes)
├── internal/                 # Private application code, organized by domain
│   ├── database/             # Database setup and access (e.g., connection, migrations)
│   ├── dto/                  # Data Transfer Objects for structuring input and output data
│   ├── models/               # Data models or ORM entities for database operations
│   ├── repository/           # Interfaces and implementations for data persistence
│   ├── router/               # HTTP router setup, including API route definitions
│   ├── service/              # Business logic and core application services
├── pkg/                      # Utility and reusable libraries shared across the application
├── tests/                    # Unit and integration tests
├── .env                      # Environment variables configuration file
├── docker-compose.yaml       # Docker Compose file for container orchestration
├── Dockerfile                # Dockerfile for building the application container
├── main.go                   # Application entry point
├── Makefile                  # Automation tasks (e.g., build, test, deploy)
└── README.md                 # Project overview and usage guide
```

## Authors

- [@fandinxr](https://www.github.com/fnxr21)

