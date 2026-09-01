# Roadmap — Disk Usage Analyzer

My execution plan for this project, broken into phases. I'll check items off as I go.

## Phase 1 — Small Project

- [x] Set up project (`go mod init`) → verify: `main.go` & `go.mod` created
- [x] Read directory path from CLI arg (`os.Args`) → verify: path printed back, clear error if missing
- [ ] List one directory's contents, non-recursive (`os.ReadDir`) → verify: entry names + type (file/dir) printed
- [ ] Sum file sizes in one directory (`entry.Info().Size()`) → verify: total matches File Explorer / `du -sb`
- [ ] Make it recursive with `filepath.WalkDir` → verify: total size of a large folder matches a reference tool
- [ ] Aggregate size per directory (struct + map) → verify: size of each top-level subdirectory printed
- [ ] Sort descending + human-readable output (KB/MB/GB) → verify: output sorted largest → smallest
- [ ] Handle `permission denied` dirs gracefully (`filepath.SkipDir`) instead of crashing → verify: scanning a restricted folder doesn't stop the whole program
- [ ] Replace `fmt.Println` debug output with `slog` (e.g. log which dirs were skipped and why) → verify: skipped-dir messages show up as structured log lines

## Phase 2 — Medium Project

> Concurrency goes here, not Phase 3 like my Go rules default to — it's the whole point of this
> project, and `-race`/benchmarks below are meaningless until the code is actually parallel.

- [ ] Identify what can be parallelized → verify: can explain which subdirectories are independent
- [ ] Add goroutines + worker pool / `errgroup.SetLimit` → verify: result identical to sequential version, faster on large trees
- [ ] Synchronize shared results (mutex/channel) → verify: `go run -race` reports no warnings
- [ ] Benchmark sequential vs parallel (`go test -bench`) → verify: have numbers showing when parallel wins
- [ ] Set up `.golangci.yml` (golangci-lint config) → verify: `golangci-lint run` passes
- [ ] Set up `gosec` (insecure pattern scan) + `govulncheck` (dependency vuln scan) → verify: both run clean
- [ ] Write table-driven tests (`go test -race ./...`)
- [ ] `Justfile` (`build`, `test`, `lint`, `security`, `check`) wrapping `go mod tidy`, `go vet ./...`, `golangci-lint run`, `goimports -w .`, `go test -race ./...`, `govulncheck ./...`
- [ ] Verify: `just check` all green, push per feature

## Phase 3 — Large Project (optional)

This CLI is realistically "done" after Phase 2 — Phase 3 is optional polish. I'd only come back to it if I want to harden it further, and concurrency shows up once more here (Ctrl+C handling).

- [ ] Handle Ctrl+C (SIGINT) during a long scan via `signal.NotifyContext` → cancel in-flight goroutines instead of an abrupt kill → verify: pressing Ctrl+C mid-scan stops within ~1s, no leftover goroutines
- [ ] Build with version info (`-ldflags`), log version at startup
- [ ] Track tool dependencies with `tool` directive in `go.mod` (goimports, golangci-lint, govulncheck, gosec) instead of a `tools.go` hack

## Backlog

Not committing to these — just don't want to lose the idea.

- JSON output mode, so results can be piped into other tools
- Filter by extension/pattern
- Exclude patterns (`.git`, `node_modules`, ...)
- Interactive TUI (like `ncdu`)
