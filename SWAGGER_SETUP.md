# Swagger UI Setup Guide

This project uses **swaggo** with **gin-swagger** to provide interactive API documentation.

## Documentation

The Swagger UI is available at:
```
http://localhost:3000/swagger/index.html
```

The API documentation is also available as:
- JSON format: `/swagger/doc.json`
- YAML format: `/swagger/swagger.yaml`

## Building Swagger Docs

### Locally in WSL (with Go installed)

Install the swag CLI tool:
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

Generate/regenerate the documentation:
```bash
swag init -g server.go
```

This will create/update:
- `docs/docs.go`
- `docs/swagger.json`
- `docs/swagger.yaml`

### In Docker

The Dockerfile automatically:
1. Installs the swag CLI tool
2. Runs `swag init` during the build process
3. Includes the generated docs in the final image

### In GitHub Actions

The workflow automatically:
1. Checks out the repository
2. Sets up Go
3. Installs the swag CLI
4. Generates Swagger docs before building the Docker image

## Adding API Documentation

### To Add Swagger Comments

1. **For the main API info**, edit `server.go`:
   ```go
   // @title Portfolio API
   // @description API for portfolio website
   // @version 1.0
   ```

2. **For endpoints**, add comments above handler functions:
   ```go
   // @Summary Get user profile
   // @Description Get the user's profile information
   // @Accept json
   // @Produce json
   // @Success 200 {object} UserProfile
   // @Router /profile [get]
   func getProfile(c *gin.Context) {
       // ...
   }
   ```

3. **For endpoints requiring authentication**:
   ```go
   // @Security BasicAuth
   ```

### Swagger Comment Reference

Common annotations:
- `@Summary` - Brief description
- `@Description` - Detailed description
- `@Accept` - Content types accepted (json, x-www-form-urlencoded, etc.)
- `@Produce` - Response content types
- `@Param` - Parameter definition
- `@Success` - Successful response definition
- `@Failure` - Error response definition
- `@Router` - Route path and HTTP method
- `@Security` - Security requirements
- `@Deprecated` - Mark endpoint as deprecated

## File Structure

```
.
├── docs/                      # Generated Swagger documentation
│   ├── docs.go               # Go swagger definitions
│   ├── swagger.json          # Swagger JSON spec
│   └── swagger.yaml          # Swagger YAML spec
├── handlers/                  # Handler files with swagger comments
│   ├── handlers.go
│   ├── blogs.go
│   ├── experience.go
│   ├── technologies.go
│   ├── misc.go
│   └── redirection.go
├── server.go                  # Main swagger info definitions
└── go.mod                     # Updated with swaggo dependencies
```

## Testing Swagger UI

After generating docs and running the server:

1. Navigate to: `http://localhost:3000/swagger/index.html`
2. You should see all API endpoints listed
3. Click on any endpoint to see details
4. For authenticated endpoints, click "Authorize" and enter credentials

## Troubleshooting

### Docs not showing up?
- Ensure `swag init` was run successfully
- Check that `docs/` directory exists with generated files
- Verify the import in handlers.go: `_ "github.com/metamemelord/portfolio-website/docs"`

### Changes not reflected?
- Run `swag init -g server.go` again after making changes
- The docs are generated from code comments, not runtime

### Swagger UI returns 404?
- Ensure the route is registered: `g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))`
- Check that swaggo packages are installed

## Dependencies

The following dependencies were added to support Swagger:
- `github.com/swaggo/swag` - Swagger documentation generator
- `github.com/swaggo/files` - Embedded Swagger files
- `github.com/swaggo/gin-swagger` - Gin integration for Swagger UI

These are managed in `go.mod` and installed automatically via `go mod download`.
