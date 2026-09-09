# CGTA — Conversation to Code Compiler

Production monorepo foundation for the Conversation → Parser → Planner → Code Generator → Filesystem → Deployment → Observability → AI/Agent pipeline.

## Engineering rules

- Production implementation only; no demo/simulation/prototype paths.
- Go backend with Fiber-compatible HTTP boundary.
- Flutter application layer is kept as a first-class client surface.
- Security, validation, auditability, rollback and approval gates are first-class concerns.
- Keep dependencies minimal and isolated by responsibility.
- Use `cat > file` / `eco` in bootstrap automation; do not depend on `nano`.

## Current repository state

This repository was empty when initialization began. The repository is therefore being initialized as a clean production monorepo foundation rather than pretending that previous generated modules already existed here.

The architecture follows the project's existing module plan, including the parser/planner/code-generator chain, AI layer, API/security, database, workflow, source control, testing, documentation, network/ISP, telemetry, autonomous operations and Flutter/NOC surfaces.

## Layout

```text
cmd/cgta/                 application entrypoint
internal/core/            compiler orchestration primitives
internal/config/          environment/config loading
internal/http/            HTTP health/readiness boundary
api/                      API contracts
configs/                  deployment configuration
migrations/               database migrations
deploy/                   Docker/Kubernetes deployment assets
web/                      web application boundary
mobile/                   Flutter application boundary
scripts/                  repository/bootstrap/verification scripts
tests/                    integration and acceptance tests
```

## Verification

```bash
go test ./...
go vet ./...
go build ./cmd/cgta
```
