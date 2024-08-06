# Realtime Chat Server

## About
Realtime chat server is my side project aim to sharpen Golang and software design skill after long time be in management role. I decide to implement this service to challenge myself to create the system that should provide technical in term of concurrency, realtime, scalability,  and maintainability.

The service contains four components:
1. **Server** - Contains business logic and all features.
2. **Database** - PostgreSQL database for persistent storage.
3. **Migrate** - Responsible for data migration and seeding.
4. **Redis** - Used for token management.

## Features
**Current**
- [x] Authentication - Signup / Signin / Signout
- [ ] User profile management
- [ ] Chatroom
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
The service is implemented using Go as the base language.

**Service Architecture**
Applying Clean Architecture, the project is divided by service, with each service divided into layers aligned with Clean Architecture. You can see the folder structure below:

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

1. **Entity** - The innermost layer, containing only business entities, independent of external service tools or any dependencies.
2. **Usecase** - This layer contains business logic that interacts with entities and the third layer (referred to as connectors). It is designed to minimize dependencies on external services using dependency injection techniques.
3. **Interface Adapters** - This layer acts as an interface to connect use cases with external systems. It strictly adheres to interfaces, with implementations aligned to the dependencies/external tools used. It contains no business logic, ensuring changes to external dependencies affect only this layer, making it cost-effective and ensuring consistent business performance.
4. **External** - This layer includes all external services used by the system, such as databases, client UIs, or other third-party services.

**Dependencies**
These are the main dependencies used:
- `gorilla/mux` - HTTP server framework
- `golang-jwt/jwt/v5` - Handles token management
- `viper` - Manages ENV configuration files and Docker ENV
- `pq` - Connects and manages the database
- `DATA-DOG/go-sqlmock` - Facilitates mocking the database for unit tests on repositories

**Technology**
- **Docker and Docker Compose** - Used for easier running and deployment of the service, providing containerization capabilities.

**Upcoming Technology**
- Load balancer
- Database replication
- Database partitioning (both vertical and horizontal) for experimentation

## Running the Server
Available commands:
- `make test` - Runs test cases
- `make coverage` - Runs test cases and outputs coverage report in HTML
- `make clean-test` - Cleans up test reports
- `make build` - Builds the project
- `make run` - Builds and runs the project
- `make docker-run` - Runs all containers
- `make docker-run-dependency` - Runs the database and Redis (used when running the server locally or skipping DB migration)
- `make run-migrate` - Called by migrate
- `make migrate` - Runs migration, accepts two variables: `ACTION` (up or down) and `VERSION` (number of versions to forward or backward)
- `make clean` - Cleans up the project, removes build files, and stops all containers.

You can see detailed commands in the Makefile.

## Contact
Interested in anything or want to have a chat with me? Contact me at:
- Email: thanapat.jutha@gmail.com
- GitHub: [https://github.com/patja60](https://github.com/patja60)
- LinkedIn: [https://www.linkedin.com/in/thanapat-juthavantana/](https://www.linkedin.com/in/thanapat-juthavantana/)


## License