# Arcflow

Arcflow is a modern full-stack SaaS platform built with:

- Go (Chi)
- PostgreSQL
- Next.js (TypeScript + Tailwind)
- Docker & Docker Compose
- GitHub Actions (CI/CD)
- AWS (Phase 2 deployment)

---

## Architecture

Frontend → Next.js  
Backend → Go (Chi REST API)  
Database → PostgreSQL (Dockerized locally)

---

## Current Backend Structure

backend/
├── cmd/server        → Application entrypoint  
├── internal/config   → Environment configuration  
├── internal/db       → PostgreSQL connection layer  
└── go.mod            → Module definition  

---

## Development Status

### Phase 1 – Core Backend Infrastructure
- [x] Go module initialization
- [x] Chi router with health endpoint
- [x] Environment-based configuration
- [x] PostgreSQL connection pool (pgx)
- [x] Dockerized PostgreSQL via Compose
- [ ] Database migrations system
- [ ] Domain models & handlers
- [ ] Authentication (JWT)

### Phase 2 – Production Readiness
- [ ] Containerized backend service
- [ ] CI pipeline (GitHub Actions)
- [ ] AWS deployment (ECS + RDS)
- [ ] Structured logging & observability
