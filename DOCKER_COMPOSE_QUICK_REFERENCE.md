# Docker Compose Quick Reference

## One-Liner Start

```bash
cp .env.example .env && nano .env && docker-compose up -d
```

## Common Commands

| Task | Command |
|------|---------|
| Start all services | `docker-compose up -d` |
| Stop all services | `docker-compose down` |
| View logs | `docker-compose logs -f` |
| Service status | `docker-compose ps` |
| Access app | http://localhost:3000 |
| Access Swagger UI | http://localhost:3000/swagger/index.html |
| MongoDB shell | `docker-compose exec mongodb mongosh -u admin -p --authenticationDatabase admin` |
| Rebuild app | `docker-compose build --no-cache` |
| Restart app | `docker-compose restart portfolio` |
| Full reset | `docker-compose down -v && docker-compose up -d` |

## Environment Variables Required

```bash
# Absolute minimum to start
APP_AUTH=admin:password
MONGO_URI=mongodb://admin:password123@mongodb:27017/portfolio?authSource=admin
MONGO_PASSWORD=password123

# For email (optional)
MS_TENANT_ID=your-tenant-id
MS_CLIENT_ID=your-app-id
MS_EMAIL_KEY=your-client-secret
MS_GRAPH_SELF_USER_ID=your-user-id
```

## Service Status

```
Service      Container            Port    Status
------------ -------------------- ------- --------
MongoDB      portfolio-mongodb    27017   Internal
Portfolio    portfolio-app        3000    Public
```

## Docker Compose File Structure

```yaml
Services:
├── mongodb          # Database (required)
│   ├── Port: 27017 (internal)
│   ├── Health: mongosh ping
│   └── Volumes: mongodb_data, mongodb_config
│
└── portfolio        # Application
    ├── Port: 3000 (configurable)
    ├── Depends on: mongodb (healthy)
    ├── Health: GET /health
    └── Build: From local Dockerfile

Networks:
└── portfolio        # Internal bridge network

Volumes:
├── mongodb_data     # Database persistence
└── mongodb_config   # MongoDB config persistence
```

## Typical Startup Sequence

1. `docker-compose up -d` called
2. MongoDB image pulled (if not cached)
3. Portfolio image built (first time ~5 minutes)
4. MongoDB container started
5. MongoDB health check runs (waits for ready)
6. Portfolio container started (after MongoDB healthy)
7. Portfolio app starts and connects to MongoDB
8. Services are ready to use

**First startup**: ~5-10 minutes
**Subsequent startups**: ~10-30 seconds

## .env File Template

```bash
# Required
APP_AUTH=admin:MyPassword
MONGO_PASSWORD=MyMongoPassword
MONGO_URI=mongodb://admin:MyMongoPassword@mongodb:27017/portfolio?authSource=admin

# Optional - Microsoft 365 Email
MS_TENANT_ID=
MS_CLIENT_ID=
MS_EMAIL_KEY=
MS_GRAPH_SELF_USER_ID=
SELF_EMAIL=noreply@gaurav.dev

# Optional - Server Config
PORT=3000
GIN_MODE=release
ENV=release
KEEP_ALIVE_CRON=false
KEEP_ALIVE_BASE_URL=http://localhost:3000

# Optional - TLS (leave blank to disable)
TLS_ENABLED=false
TLS_CERT_PATH=
TLS_KEY_PATH=
```

## Debugging

```bash
# Check if containers are running
docker-compose ps

# View real-time logs
docker-compose logs -f

# Check specific container
docker-compose logs portfolio

# Execute command in container
docker-compose exec portfolio sh

# Check MongoDB connection
docker-compose exec portfolio echo $MONGO_URI

# Test API endpoint
curl http://localhost:3000/health

# Restart single service
docker-compose restart portfolio
```

## File Structure in Repository

```
portfolio-website/
├── docker-compose.yml          # Docker Compose definition
├── .env.example                # Example environment variables
├── DOCKER_COMPOSE_GUIDE.md     # Full documentation
├── Dockerfile                  # Multi-stage build
├── server.go                   # Main entry point
├── handlers/                   # API handlers
├── model/                      # Data models
├── db/                         # Database setup
├── pkg/                        # Packages
└── go.mod                      # Go dependencies
```

## Network Diagram

```
┌─────────────────────────────────────────────┐
│  Host Machine (Your Computer)               │
├─────────────────────────────────────────────┤
│                                             │
│  Port 3000 ──────► Portfolio Container      │
│                   ├─ Port 3000 (internal)   │
│                   └─ http://localhost:3000  │
│                                             │
│  Port 27017 ──────► MongoDB Container       │
│                    ├─ Port 27017 (internal) │
│                    └─ mongodb://localhost   │
│                                             │
│  Docker Bridge Network (portfolio)          │
│  ├─ portfolio (10.0.x.x)                    │
│  ├─ mongodb (10.0.x.x)                      │
│  └─ Can communicate freely                  │
│                                             │
└─────────────────────────────────────────────┘
```

## Scaling & Load Balancing

- **Not needed**: Portfolio is single-instance application
- **MongoDB**: Single instance (can upgrade to replica set later)
- **For high traffic**: Use reverse proxy (nginx) in front of single Portfolio container

## Monitoring

```bash
# CPU and Memory usage
docker stats

# Container inspect
docker inspect portfolio_portfolio_1

# Check resource limits
docker inspect portfolio_portfolio_1 | grep -i '"memory"'
```

## Integration with CI/CD

```bash
# Build image (called by CI/CD)
docker-compose build

# Push to registry
docker tag portfolio_portfolio:latest registry.io/portfolio:latest
docker push registry.io/portfolio:latest

# Pull and run in production
docker pull registry.io/portfolio:latest
docker-compose up -d
```
