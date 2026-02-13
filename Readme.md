- Flat Folder Structure
```
project/
├── main.go
├── config
|   ├── db
|   ├── env               # Application entry point
├── router
|   |–– router.go          # HTTP handlers
├── controllers
|   |–– handler.go          # HTTP handlers
├── services
|   |–– service.go              # Business logic
├── db
|   |–– repository       # Database repository
|   |–– migrations
├── utils.go            # Utility functions
├── go.mod              # Go module file
└── go.sum              # Go module dependency file
```

- used for
    - Small Projects:
    - Prototypes and MVPs
    - Command-Line Tools
    - Learning and Experimentation: Beneficial for beginners