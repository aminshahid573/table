# Agent Instructions for gochat

## Project Overview

Go TUI chat application built with the Charm stack (Bubble Tea, Lipgloss, Bubbles).
All files are in `package main`. Module name in `go.mod` is `table`; Go version is 1.25.5.

### File Layout

| File | Responsibility |
|---|---|
| `main.go` | Entry point only (`func main`) |
| `model.go` | `model` struct, `initialModel()`, `Init()`, `Update()`, `recalcLayout()` |
| `view.go` | `View()` rendering method |
| `styles.go` | All lipgloss style declarations |

## Build / Run / Test Commands

```bash
# Install dependencies
go mod tidy

# Build
go build -o app .

# Run
go run main.go

# Run all tests (none exist yet)
go test ./...

# Run a single test by name
go test ./... -run TestName

# Run tests with verbose output
go test -v ./...

# Vet (static analysis)
go vet ./...

# Format code
gofmt -w .
```

There is no Makefile, no linter config (`.golangci.yml`), and no CI/CD pipeline.
No Cursor rules, Copilot instructions, or `.editorconfig` exist.

## Code Style Guidelines

### Imports

Organize imports in two groups separated by a blank line:
1. Standard library
2. External/third-party packages

Alias `bubbletea` as `tea`. Keep imports alphabetically sorted within each group.

```go
import (
    "fmt"
    "os"

    "github.com/charmbracelet/bubbles/textarea"
    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"
)
```

When internal packages are added, use a third group after external imports.

### Formatting

- Use `gofmt` standard formatting.
- The codebase currently uses CRLF line endings; be consistent with existing files.

### Naming Conventions

- **Types**: Unexported, lowercase single-word names for models (`model`).
- **Fields**: `camelCase`, unexported. Descriptive names (`centerRenderedWidth`, not `crw`).
- **Style variables**: `camelCase` with `Style` suffix, grouped in a single `var (...)` block.
  Example: `headerContainerStyle`, `messageBoxStyle`, `leftSidebarStyle`.
- **Constructors**: Use `initialModel()` pattern (returns value, not pointer).
- **Method receivers**: Single letter matching the type (`m` for `model`).
- **Local variables**: Use `:=` short declaration. No single-letter names except
  receivers and short-lived initializers (`ti`, `ta`).

### Types and Structs

- Use flat structs; compose by embedding Bubble Tea component models as value fields.
- Pointer receivers (`*model`) for all methods on the main model.
- Implement `tea.Model` interface implicitly (no explicit interface declaration needed).
- Constructor pattern: instantiate with `New()`, configure fields, return struct literal.

### Error Handling

- At the top level (`main()`), use print-and-exit: `fmt.Println("Error:", err); os.Exit(1)`.
- Do not create custom error types unless needed for specific error discrimination.
- In Bubble Tea programs, errors flow through `tea.Cmd` and messages, not return values.

### Comments

- Use `// --- N. SECTION NAME ---` banner comments to delimit rendering sections in `View()`.
- Write "why" comments for non-obvious layout math (padding, border offsets).
- Add brief label comments above style variable groups.
- No godoc-style comments are currently used; add them for any new exported symbols.

### Logging

- Do not use `log`, `slog`, or third-party loggers. In Bubble Tea, stdout is owned
  by the terminal renderer. Use `tea.Println()` or Bubble Tea's built-in logging
  (`tea.LogToFile`) for debugging only.

## UI Development Rules

### General Guidelines

- Never use commands to send messages when you can directly mutate children or state.
- Keep things simple; do not overcomplicate.
- Create files if needed to separate logic; do not nest models.
- Never do IO or expensive work in `Update`; always use a `tea.Cmd`.
- Never change the model state inside of a command; use messages and update state
  in the main `Update` loop.

### Architecture (Target)

When refactoring into multiple files, follow this structure:

- `model/` - Main UI model and major components (chat, sidebar)
- `chat/` - Chat message item types and renderers
- `dialog/` - Dialog implementations
- `list/` - Generic list component with lazy rendering
- `common/` - Shared utilities and the Common struct
- `styles/` - All style definitions
- `anim/` - Animation system
- `logo/` - Logo rendering

### Components Should Be Dumb

Components should not handle Bubble Tea messages directly. Instead:
- Expose methods for state changes
- Return `tea.Cmd` from methods when side effects are needed
- Handle their own rendering via `Render(width int) string`

### Styling

- Define all styles as package-level `var` declarations using `lipgloss.NewStyle()`.
- Use method chaining for style construction.
- Use `lipgloss.Color("240")` (ANSI) or `lipgloss.Color("#FFFFFF")` (hex) for colors.
- Prefer semantic color names and variables over hardcoded color values.
- Always account for padding and borders in width calculations:
  a border adds 2 to rendered width, padding adds per-side.

### Layout Patterns

- Use `lipgloss.JoinHorizontal()` and `lipgloss.JoinVertical()` for composition.
- Use `lipgloss.Width()` and `lipgloss.Height()` to measure rendered strings.
- Recalculate layout on `tea.WindowSizeMsg`; store dimensions in the model.
- Guard against zero/negative dimensions with minimum value checks.
- Return `"Loading..."` from `View()` before the first `WindowSizeMsg`.

### Key Patterns

- Use `tea.Batch()` when returning multiple commands from `Init()` or `Update()`.
- Use type switches (`switch msg := msg.(type)`) for message handling in `Update()`.
- Collect commands in a `[]tea.Cmd` slice, append each sub-component's cmd, and
  return `tea.Batch(cmds...)`.
- Pass `*common.Common` to components that need styles or app access (when packages exist).

## Known Issues (from Todo.md)

- Status bar width mismatch with other boxes
- Message field horizontal scrolling/wrapping issue
- Message field icons and prompt misalignment on multi-line input
