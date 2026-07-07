# Mimic — Project Checklist

## Core CLI Commands

- [ ] `mimic new version` — create a new migration script file with timestamp/ordering prefix
- [ ] `mimic status` — show migration state (pending, running, applied, failed, rolled_back)
- [ ] `mimic run` — run the next unran migration in the stack
- [ ] `mimic run-all` — run all pending migrations in order
- [ ] `mimic rollback` — rollback the latest applied migration
- [x] `mimic setup` — initial DB setup (creates tracking tables, advisory lock support)
- [ ] `mimic run-script` — run a tracked one-off script
- [ ] `mimic dry-run` — preview what would run without executing

---

## Migration State Model

- [x] Define migration states: `running`, `applied`, `failed`, `rolled_back` (`pending` is derived from disk vs DB diff)
- [x] Use `running` state to detect mid-flight crashes vs never-started migrations
- [x] Each state transition is a new INSERT — no UPDATEs, no DELETEs (append-only event log)
- [x] `occurred_at` timestamp on every event row; duration derived from first and last event per version
- [x] Store error message on `failed` rows

---

## Tracking Tables (Design Before Writing Code)

- [x] Design `mimic_migrations` table schema
- [x] columns: id, version, name, state, occurred_at, error (append-only event log — one row per state transition)
- [x] Design `mimic_scripts` table schema
- [x] columns: id, name, state, occurred_at, error (same append-only model)
- [x] Ensure tracking tables themselves are idempotent to create (safe to re-run `setup`)

---

## Concurrency & Locking

- [ ] Use `pg_try_advisory_lock` before starting any migration run
- [ ] If lock cannot be acquired, report clearly and exit

---

## Progress & Output

- [ ] Per-migration verbose output available via `--verbose` flag
- [ ] Structured JSON output mode via `--json` flag (for CI/automation)

---

## Migration Ordering

- [ ] Define and document the ordering scheme (timestamp prefix, sequential int, etc.)
- [ ] Guard against two migrations with identical ordering keys
- [ ] Consider merge conflict behavior on branches (document the resolution approach)
- [ ] Enforce linear history in CI if desired (optional strict mode)

---

## Rollback

- [ ] Require explicit rollback SQL — no auto-generation
- [ ] Block or warn on `new version` if no rollback file is provided
- [ ] Rollback only targets the latest `applied` migration
- [ ] Rollback updates state to `rolled_back`, does not delete the record

---

## Auth & Connection

- [x] Connection string auth (standard `postgres://...` URL)
- [ ] IAM auth for RDS (fetch short-lived token via AWS SDK before connecting)
- [ ] Pluggable auth interface so additional providers can be added (Vault, etc.)
- [x] Connection config via env vars and CLI flags (precedence documented)

---

## Type Generation

- [ ] First-party plugin: Postgres enums → TypeScript enums
- [ ] Pluggable codegen interface: `mimic codegen --runner="<command>"` with defined contract
- [ ] Contract: receive connection string + output directory, produce files
- [ ] Document the plugin contract so community can build integrations (pg-to-ts, Zapatos, Drizzle, etc.)
- [ ] Codegen is opt-in, not tied to migration runs

---

## Script Running

- [ ] Scripts are tracked the same way migrations are (same state model)
- [ ] Re-running a script that already applied is blocked by default (`--force` to override)
- [ ] Scripts have their own directory separate from migrations

---

## Dry Run

- [ ] `--dry-run` flag works on `run`, `run-all`, `rollback`, and `run-script`
- [ ] Dry run prints exactly what SQL would execute
- [ ] Dry run never writes to tracking tables
- [ ] Dry run exit code reflects whether there is pending work (useful for CI checks)

---

## Developer Experience

- [ ] `mimic setup` generates config file template
- [ ] Clear error messages with actionable next steps (no raw Postgres errors surfaced naked)
- [ ] `--help` on every command
- [ ] README covers: installation and quickstart
- [ ] Document all env vars and config file options in one place

---

## Open Source Readiness

- [ ] Works against any Postgres-compatible DB: vanilla Postgres, Supabase, Neon, RDS, Cloud SQL
- [ ] No vendor-specific dependencies in the core
- [ ] LICENSE file (decide: MIT, Apache 2.0)
- [ ] CONTRIBUTING guide
- [ ] GitHub Actions CI: lint, test, build on PR
- [ ] Semantic versioning from day one

---

## Future / Backlog

- [ ] Support schema-per-tenant multi-tenant deployments
- [ ] ORM integration via plugin contract (same pattern as codegen)
- [ ] Web UI / dashboard for migration status
- [ ] Migration diff viewer (what changed between two versions)
- [ ] Slack/webhook notifications on failure
- [ ] Audit log export
