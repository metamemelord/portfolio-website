[![Build and publish](https://github.com/metamemelord/portfolio-website/actions/workflows/build.yml/badge.svg)](https://github.com/metamemelord/portfolio-website/actions/workflows/build.yml)

# Gaurav Saini Portfolio Website 

**Live at**: [grv.app](https://grv.app) | [gaurav.dev](https://gaurav.dev) | [gauravsaini.dev](https://gauravsaini.dev) | [metamemelord.com](https://metamemelord.com) | [metamemelord.dev](https://metamemelord.dev)

## Features

### 📊 Content Management
- **Blog Posts**: Create, read, update, and delete blog posts with visibility control
- **Work Experience**: Manage professional experience entries with dates, companies, and descriptions
- **Technologies**: Curate and showcase technology stack with metadata and ordering
- **Social Media Links**: Manage social profiles with custom handles and metadata
- **User Profile**: Dynamic portfolio profile with occupation, location, and contact information

### 🔄 Real-time Data Integration
- **GitHub Repositories**: Live sync of GitHub repositories with automatic periodic refresh
- **WordPress Blogs**: Live sync of WordPress blog posts with scheduled updates
- **Automatic Scheduler**: Periodic data refresh every 2 hours to keep content current
- **Data Caching**: Intelligent caching with configurable TTL for performance

### 🔗 URL Management
- **Built-in URL Shortener**: Create short, memorable URLs backed by MongoDB
- **Query Parameter Support**: Forward query parameters through short links
- **Path Forwarding**: Support for path-based routing in redirections
- **Expiry Management**: Set expiration dates for temporary short links
- **Hit Tracking**: Track click counts on all redirections
- **Subdomain Redirection**: Route requests through subdomains without relying on DNS changes

### 🔐 Authentication & Security
- **Basic Auth**: Secure endpoints with basic authentication
- **Credential Verification**: Admin-only operations for content management
- **Role-based Access**: Separate public GET endpoints from authenticated POST/PUT/DELETE operations

### 📧 Communication
- **Email Integration**: Send emails via Microsoft 365 Graph API
- **Contact Form**: Public endpoint for portfolio visitors to send messages
- **Dynamic Email Routing**: Support for custom email templates and recipients

### 📚 API Documentation
- **Swagger UI**: Complete interactive API documentation at `/swagger/index.html`
- **Auto-generated Docs**: Documentation automatically generated and embedded in Docker image
- **CI/CD Integration**: Swagger docs regenerated on every deployment

### 🐳 DevOps & Infrastructure
- **Docker Support**: Multi-stage Docker build for optimized production images
- **Docker Compose**: Complete local development environment with MongoDB
- **MongoDB Integration**: Persistent data storage with health checks
- **Environment Configuration**: 18 configurable environment variables
- **TLS/HTTPS**: Optional TLS encryption with custom certificates
- **GitHub Actions**: Automated CI/CD pipeline with multi-registry publishing
- **Keep-Alive Service**: Automatic uptime monitoring and keep-alive pings
- **Container Health Checks**: Built-in health endpoints for orchestration

### 🎨 Frontend
- **Vue.js Framework**: Modern SPA frontend
- **Static Asset Serving**: Optimized CSS, JavaScript, and image delivery
- **Sitemap & Robots.txt**: SEO optimization support
- **CORS Support**: Cross-origin requests properly configured

### 🛠️ Developer Experience
- **Hot Reload Support**: Development mode with rebuild capabilities
- **Code Generation**: Enum types generated from code
- **Dependency Management**: Go modules for backend, npm for frontend
- **Database Persistence**: MongoDB volumes for data retention across restarts
- **Validation Scripts**: Configuration validation for both Unix and Windows
