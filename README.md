# Realtime Chat Server
## Table of Contents
- [About](#about)
- [Features](#features)
- [Development](#development)
- [Running the Server](#running-the-server)
- [Contact](#contact)
- [License](#license)

## About
Realtime Chat Server is a side project aimed at honing my Golang and software design skills after spending a significant amount of time in a management role. This project challenges me to create a system that excels in concurrency, real-time operations, scalability, and maintainability.

The service comprises four main components:
1. **Server** - Contains business logic and all features.
2. **Database** - PostgreSQL database for persistent storage.
3. **Migrate** - Responsible for data migration and seeding.
4. **Redis** - Used for token management.

## Features
**Current**
- [x] Authentication: Signup / Signin / Signout
- [ ] User profile management
- [ ] Chatroom:
  - [ ] One-to-one chat
  - [ ] Group chat
  - [ ] Mark as read
  - [ ] Typing indicator
  - [ ] Join and leave alerts
- [ ] Notification

**Upcoming**
- [ ] Client UI
- [ ] Multimedia messaging

## Development
This service is implemented using Go as the base language.

**Service Architecture**
Following Clean Architecture principles, the project is divided into services, each containing layers aligned with Clean Architecture. The folder structure is as follows:

```
├── internal
│   ├── auth (service)
│   │   ├── entity // -> Entity layer
│   │   ├── usecase // -> Usecase layer
│   │   ├── repository -> Interface Adapters layer
│   │   └── handler -> Interface Adapters layer
│   │
│   └── another service
│       └── ...
```

1. **Entity** - The innermost layer, containing only business entities, independent of external service tools or dependencies.
2. **Usecase** - Contains business logic interacting with entities and the connector layer. Designed to minimize external dependencies using dependency injection techniques.
3. **Interface Adapters** - Connects use cases with external systems, strictly adhering to interfaces. Implementations align with dependencies/external tools. Contains no business logic, ensuring cost-effective external dependency changes.
4. **External** - Includes all external services such as databases, client UIs, or other third-party services.

**Dependencies**
Key dependencies used:
- `gorilla/mux` - HTTP server framework
- `golang-jwt/jwt/v5` - Token management
- `viper` - ENV configuration and Docker ENV management
- `pq` - Database connection and management
- `DATA-DOG/go-sqlmock` - Database mocking for unit tests

**Technology**
- **Docker and Docker Compose** - Facilitate easy running and deployment of the service, providing containerization capabilities.

**Upcoming Technology**
- Load balancer
- Database replication
- Database partitioning (vertical and horizontal) --> for fun

## Running the Server
Available commands:
- `make test` - Run test cases
- `make coverage` - Run tests and generate a coverage report in HTML
- `make clean-test` - Clean up test reports
- `make build` - Build the project
- `make run` - Build and run the project
- `make docker-run` - Run all containers
- `make docker-run-dependency` - Run database and Redis (for local server runs or skipping DB migration)
- `make run-migrate` - Called by migrate
- `make migrate` - Run migration (accepts `ACTION` (up/down) and `VERSION` (number of versions to move forward/backward))
- `make clean` - Clean up the project, remove build files, and stop all containers

Detailed commands can be found in the Makefile.

## Contact
Interested in discussing anything or having a chat? Reach out to me at:
- Email: thanapat.jutha@gmail.com
- GitHub: [https://github.com/patja60](https://github.com/patja60)
- LinkedIn: [https://www.linkedin.com/in/thanapat-juthavantana/](https://www.linkedin.com/in/thanapat-juthavantana/)

## License