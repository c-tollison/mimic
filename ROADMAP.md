# Mimic — Project Checklist

## Core CLI Commands

- [ ] `mimic new version` — create a new migration script file with timestamp/ordering prefix
- [ ] `mimic status` — show migration state per schema (pending, running, applied, failed, rolled_back)
- [ ] `mimic run` — run the next unran migration in the stack
- [ ] `mimic run-all` — run all pending migrations in order
- [ ] `mimic rollback` — rollback the latest applied migration
- [ ] `mimic new-network` — provision a new schema for multi-tenant setups
- [ ] `mimic setup` — initial DB setup (creates tracking tables, advisory lock support)
- [ ] `mimic reset` — nuke and reseed (local only, guarded hard)
- [ ] `mimic create-test-data` — scaffold a new test data script
- [ ] `mimic fill-test-data` — run a test data script against a target
- [ ] `mimic run-script` — run a tracked one-off script (single schema or all schemas)
- [ ] `mimic dry-run` — preview what would run without executing

---

## Migration State Model

- [ ] Define migration states: `pending`, `running`, `applied`, `failed`, `rolled_back`
- [ ] Use `running` state to detect mid-flight crashes vs never-started migrations
- [ ] Record start time and end time on every migration run
- [ ] Store error message/stack on `failed` records
- [ ] All state transitions are append-only (no in-place updates to history)

---

## Tracking Tables (Design Before Writing Code)

- [ ] Design `mimic_migrations` table schema
- [ ] migration id, version, name, schema, state, started_at, finished_at, error
- [ ] Design `mimic_scripts` table schema
- [ ] script id, name, schema, state, started_at, finished_at, error
- [ ] Decide: per-schema tracking tables vs global registry in `public` (or both)
- [ ] If global registry: design the shared table and its write path
- [ ] Ensure tracking tables themselves are idempotent to create (safe to re-run `setup`)

---

## Multi-Tenant Schema Support

- [ ] Support both single-schema (public or named) and multi-schema orgs
- [ ] Define schema discovery strategy at runtime:
- [ ] Config file list
- [ ] Naming convention regex (e.g. `^network_.*`)
- [ ] Registry table in shared schema
- [ ] `new-network` creates schema and runs `setup` within it
- [ ] All commands accept `--schema` flag to target a specific schema
- [ ] All commands accept `--all` flag to target every managed schema

---

## Partial Failure Handling (1000-Schema Scale)

- [ ] Per-schema isolation: one schema's failure never blocks others
- [ ] `run-all` continues to remaining schemas after a failure, collects all errors
- [ ] Failed schemas are clearly reported at the end of a run
- [ ] Re-running after failure only targets schemas that have not yet applied the migration
- [ ] Exit code reflects partial failure (non-zero if any schema failed)

---

## Concurrency & Locking

- [ ] Use `pg_try_advisory_lock` per schema before starting any migration run
- [ ] If lock cannot be acquired, report clearly and skip (do not fail silently)
- [ ] `--concurrency N` flag for parallel schema processing (default: sane low number like 5)
- [ ] Connection pool sizing accounts for `--concurrency` value
- [ ] Long-running jobs handle connection timeout/retry gracefully

---

## Progress & Output

- [ ] Live status output during `run-all` (not a firehose — summary view)
- [ ] Format: `X/1000 applied | Y failed | Z in progress`
- [ ] Per-migration verbose output available via `--verbose` flag
- [ ] Structured JSON output mode via `--json` flag (for CI/automation)
- [ ] Final summary report after any multi-schema run

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
- [ ] Rollback only targets the latest `applied` migration per schema
- [ ] Rollback updates state to `rolled_back`, does not delete the record
- [ ] Rollback against multiple schemas follows same partial-failure model

---

## Auth & Connection

- [ ] Connection string auth (standard `postgres://...` URL)
- [ ] IAM auth for RDS (fetch short-lived token via AWS SDK before connecting)
- [ ] Token refresh handling for long-running multi-schema jobs
- [ ] Pluggable auth interface so additional providers can be added (Vault, etc.)
- [ ] Connection config via config file, env vars, and CLI flags (precedence documented)

---

## Reset Guard

- [ ] Check `NODE_ENV` (or `MIMIC_ENV`) before allowing `reset`
- [ ] Require explicit `--confirm` flag
- [ ] Require `--i-know-what-im-doing` or equivalent second confirmation
- [ ] Optionally validate connection string against a known-safe local pattern
- [ ] Never allow `reset` if `MIMIC_ENV=production` is set, regardless of other flags

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
- [ ] Scripts can target a single schema or all managed schemas
- [ ] Re-running a script that already applied is blocked by default (`--force` to override)
- [ ] Scripts have their own directory separate from migrations

---

## Dry Run

- [ ] `--dry-run` flag works on `run`, `run-all`, `rollback`, and `run-script`
- [ ] Dry run prints exactly what SQL would execute, against which schemas
- [ ] Dry run never writes to tracking tables
- [ ] Dry run exit code reflects whether there is pending work (useful for CI checks)

---

## Developer Experience

- [ ] `mimic init` or `setup` generates config file template
- [ ] Clear error messages with actionable next steps (no raw Postgres errors surfaced naked)
- [ ] `--help` on every command
- [ ] README covers: installation, quickstart for single-schema, quickstart for multi-tenant
- [ ] Document all env vars and config file options in one place

---

## Open Source Readiness

- [ ] Works against any Postgres-compatible DB: vanilla Postgres, Supabase, Neon, RDS, Cloud SQL
- [ ] No vendor-specific dependencies in the core
- [ ] Config file is database-agnostic
- [ ] LICENSE file (decide: MIT, Apache 2.0)
- [ ] CONTRIBUTING guide
- [ ] GitHub Actions CI: lint, test, build on PR
- [ ] Semantic versioning from day one

---

## Future / Backlog

- [ ] ORM integration via plugin contract (same pattern as codegen)
- [ ] Web UI / dashboard for migration status across all schemas
- [ ] Migration diff viewer (what changed between two versions)
- [ ] Slack/webhook notifications on failure
- [ ] Audit log export
