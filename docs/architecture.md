## Project Structure (Clean Architecture)


project-root/
├── cmd/
│   └── ewallet/
│       └── main.go                # App entry point
│
├── internal/
│   ├── app/                       # Application layer (use cases, services)
│   │   ├── user/
│   │   │   └── service.go
│   │   ├── wallet/
│   │   │   └── service.go
│   │   ├── transaction/
│   │   │   └── service.go
│   │   ├── withdrawal/
│   │   │   └── service.go
│   │   └── topup/
│           └── service.go
│
│   ├── domain/                    # Domain layer (entities and interfaces)
│   │   ├── user/
│   │   │   ├── model.go           # User entity
│   │   │   └── repository.go      # UserRepository interface
│   │   ├── wallet/
│   │   │   ├── model.go
│   │   │   └── repository.go
│   │   ├── transaction/
│   │   │   ├── model.go
│   │   │   └── repository.go
│   │   ├── walletlog/
│   │   │   ├── model.go
│   │   │   └── repository.go
│   │   ├── withdrawal/
│   │   │   ├── model.go
│   │   │   └── repository.go
│   │   └── topup/
│           ├── model.go
│           └── repository.go
│
│   ├── infrastructure/           # Infra layer (DB, external APIs)
│   │   ├── db/
│   │   │   ├── mysql/
│   │   │   │   ├── user_repo.go
│   │   │   │   ├── wallet_repo.go
│   │   │   │   ├── transaction_repo.go
│   │   │   │   ├── walletlog_repo.go
│   │   │   │   ├── withdrawal_repo.go
│   │   │   │   └── topup_repo.go
│   │   └── migration/
│   │       └── schema.sql        # SQL files for migrations
│
│   ├── interfaces/               # Delivery layer (HTTP, gRPC, etc.)
│   │   └── http/
│   │       ├── handlers/
│   │       │   ├── user_handler.go
│   │       │   ├── wallet_handler.go
│   │       │   ├── transaction_handler.go
│   │       │   ├── walletlog_handler.go
│   │       │   ├── withdrawal_handler.go
│   │       │   └── topup_handler.go
│   │       └── router.go         # HTTP routing
│
│   └── config/
│       └── config.go             # App configuration
│
├── pkg/                          # Shared packages
│   ├── utils/
│   │   ├── uuid.go
│   │   ├── validator.go
│   │   └── logger.go
│   └── errors/
│       └── error.go
│
├── api/                          # OpenAPI/Swagger specs
│   └── openapi.yaml
│
├── web/                          # Optional frontend (if any)
│
├── scripts/                      # Helper scripts (e.g., for migrations, seeding)
│   └── seed_users.go
│
├── tests/                        # Integration or end-to-end tests
│   └── e2e_transaction_test.go
│
└── docs/                         # Documentation
└── architecture.md

-----

├── cmd/
│   ├── your-app/
│   │   ├── main.go
├── internal/
│   ├── app/
│   │   ├── handler.go
│   │   ├── service.go
│   ├── domain/
│   │   ├── model.go
│   │   ├── repository.go
├── pkg/
│   ├── utility/
│   │   ├── ...
├── api/
│   ├── ...
├── web/
│   ├── ...
├── scripts/
├── configs/
├── tests/
└── docs/
