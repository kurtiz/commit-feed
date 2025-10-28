# CommitFeed

A small CLI project written in Go. Provides a foundation for a Cobra-based command-line tool.

## Overview

`commit-feed` is a lightweight CLI application scaffolded with Cobra. It currently contains a root command and basic flag wiring. This repository is intended as the starting point for a tool that works with commit data or feeds.

## Prerequisites

- Go 1.18+ installed and on your PATH (the project uses Go modules).
- A terminal (PowerShell is used in examples below).

## Build

Open PowerShell and run the following from the repository root:

```powershell
Set-Location -Path 'C:\Users\AaronWillDjaba\Documents\go\commit-feed'
go build -v -o commit-feed.exe .
```

This produces `commit-feed.exe` (Windows) in the current directory.

You can also build for your platform with the standard `go build` form:

```powershell
go build ./...
```

## Run / Usage

The project uses Cobra for CLI handling. The root command `Use` value is defined as `commit-feed.git` in `cmd/root.go`.

Example (run built binary):

```powershell
# from project root after building
.\commit-feed.exe --help

# or run directly with `go run` during development
go run ./main.go --help
```

There is currently a simple boolean flag on the root command:

```powershell
# shorthand -t or --toggle
.\commit-feed.exe -t
```

The `--help` output will show available flags and subcommands as the project grows.

## Project layout

- `main.go` - program entrypoint that calls the Cobra `Execute()` helper.
- `cmd/` - Cobra-generated command files. `cmd/root.go` contains the root command.
- `go.mod` / `go.sum` - Go module definition and dependencies.
- `LICENSE` - project license.

## Development notes

- This repository uses Cobra (github.com/spf13/cobra). To add commands, use the Cobra generator or create files under `cmd/` that register subcommands with the root command.
- Keep the module dependencies tidy with `go mod tidy`.

## Contributing

Contributions are welcome. Please open issues or PRs for improvements, features, or bug fixes. Follow standard Go project practices and include tests where appropriate.

## License

This project includes a `LICENSE` file in the repository. Refer to it for the license terms.

## Next steps

- Add meaningful subcommands and real behavior (examples: fetch commit feed, filter, export).
- Add unit tests for commands and business logic.
- Add CI for build and tests.
