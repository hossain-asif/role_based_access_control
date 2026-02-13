- Domain-Driven Design:
In Domain-Driven Design (DDD), the application is divided into domains or bounded contexts, where each domain owns its own layers, including models, repositories, and services. This approach helps to isolate logic and keep the code domain-specific, promoting better organization and clarity.
```
project
├── cmd                    # Command-related files
│   └── app                # Application entry point
│       └── main.go        # Main application logic
├── internal               # Internal codebase
│   ├── user               # Domain 'user'
│   │   ├── handler.go     # User-specific handler
│   │   ├── service.go     # User-specific service
│   │   ├── repository.go  # User-specific repository
│   │   └── user.go        # User model
│   └── product            # Domain 'product'
│       ├── handler.go     # Product-specific handler
│       ├── service.go     # Product-specific service
│       └── repository.go  # Product-specific repository
├── pkg                    # Shared utilities
├── configs                # Configuration files
├── go.mod                 # Go module definition
└── go.sum                 # Go module checksum file
```

- used for
    - Complex Projects
    - Decoupling and Scalability
    - Modularity