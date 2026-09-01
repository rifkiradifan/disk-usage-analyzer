# Disk Usage Analyzer

A CLI that recursively calculates directory sizes and lists them sorted from largest to smallest — a Go clone of the Unix `du` command.

> Status: work in progress (Phase 1) — CLI arg parsing and single-level directory size summing are working; recursive traversal (`filepath.WalkDir`) is next.

---

## What it does

Walks a directory tree recursively with `filepath.WalkDir`, aggregates file sizes per directory, and prints results sorted largest-first in human-readable units (KB/MB/GB) — starting as a sequential implementation, then adding goroutine-based parallelism for large directory trees.

---

## Why I'm building it

With an SRE/Cloud Engineering background, I've dealt with disk-space-full incidents in production more than once — it's one of the most common causes of outages. I'm building this to learn Go concurrency fundamentals (goroutines, worker pools, race conditions) from scratch, instead of just running `du` and moving on.

---

## Stack

Planned — nothing wired up yet, see [ROADMAP.md](ROADMAP.md) for sequencing.

| Component | Technology | Notes |
|---|---|---|
| Language | Go 1.25 | |
| Directory traversal | `filepath.WalkDir` (stdlib) | |
| Concurrency | goroutines + `errgroup` | bounded worker pool |
| Linting | golangci-lint | |
| Security scan | govulncheck, gosec | |
| Testing | `go test -race`, benchmarks | |
| Task runner | Just (`Justfile`) | |

---

## Folder Structure

```
disk-usage-analyzer/
├── main.go
├── go.mod
├── .gitignore
├── LICENSE
├── README.md
└── ROADMAP.md
```

---

## Installation

Requires Go 1.25.2 or newer.

```bash
git clone https://github.com/rifkiradifan/disk-usage-analyzer.git
cd disk-usage-analyzer
```

No external dependencies yet — nothing else to install.

---

## How to Run

```bash
go run main.go <path>
```

Sums file sizes in the given directory (one level deep, subdirectories not yet included) and prints the total in bytes. Running without an argument prints a usage message to stderr and exits with code 1.

---

## Roadmap

Detailed execution roadmap is in [ROADMAP.md](ROADMAP.md).

---

## License

MIT — see [LICENSE](LICENSE).
