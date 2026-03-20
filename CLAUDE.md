# Avro CLI

.NET 10 Terminal.Gui v2 TUI application.

## Tech Stack

- .NET 10 (`net10.0`) / Terminal.Gui `2.0.0`
- Microsoft.Extensions.DependencyInjection `10.0.2`
- Central Package Management (`Directory.Packages.props`)
- Solution format: `.slnx` (XML)
- Versioning: GitVersion (GitHubFlow, next version `0.1.0`)

## Architecture

```
src/
  Avro.Cli/              # TUI entry point, DI, windows/views, theme applicator
  Avro.Cli.Core/         # Commands, handlers, theme system, interfaces (zero UI dependency)
themes/                  # .theme files (INI-format color definitions)
```

Growth path: Core splits into Domain + Application + Infrastructure when needed.

### Key Components

- **MainWindow** — static class: `Configure(Window, IThemeManager, IThemeApplicator)` builds menus and content
- **Theme system** — `IThemeManager` / `IThemeLoader` in Core; `IThemeApplicator` in Cli layer
- **SimpleThemeLoader** — INI-format parser for `.theme` files (sections: `[base]`, `[menu]`, `[dialog]`, `[error]`, `[toplevel]`)
- **Command/Handler** — `ICommand` + `ICommandHandler<T>` in Core (ready for future commands)

## Commands

```bash
dotnet build Avro.Cli.slnx                    # build (warnings as errors)
dotnet run --project src/Avro.Cli              # run TUI
task build                                     # build via Taskfile
task run                                       # run TUI via Taskfile
task deploy                                    # publish + deploy locally (osx-arm64)
```

## Terminal.Gui v2 API

Single namespace `Terminal.Gui`. Static application lifecycle:

```csharp
Application.Init();
var window = new Window { Title = "Avro CLI (Esc to quit)" };
// configure window contents via MainWindow.Configure(window, themeManager, themeApplicator)
Application.Run(window);
window.Dispose();
Application.Shutdown();
```

- **v2 changes from v1** — `Window` is instantiated directly (not `Application.Top`); `Run()` takes a `Toplevel` parameter; views are `Dispose()`d after run
- `MainWindow` — static class with `Configure(Window, IThemeManager, IThemeApplicator)` method
- `MenuBar` — horizontal menu with `MenuBarItem[]` / `MenuItem[]`
- **True Color** — theme system uses RGB `new Attribute(new Color(r, g, b), new Color(r, g, b))`
- **Avoid `null` separators** in menus — causes rendering artifacts; use separate top-level menus instead

## Theme System

Themes are INI-format `.theme` files in the `themes/` directory:

```ini
name=Theme Name
author=Author
category=dark

[base]
normal=#c0caf5
focus=#7aa2f7
hot_normal=#bb9af7
hot_focus=#ff9e64
disabled=#565f89
background=#1a1b26
```

Sections: `[base]`, `[menu]`, `[dialog]`, `[error]`, `[toplevel]` — each with `normal`, `focus`, `hot_normal`, `hot_focus`, `disabled`, `background` hex colors.

Built-in themes: Tokyo Night (default), Dracula, Nord, Rose Pine.

## Conventions

- **1 type = 1 file** — every class, interface, struct, record gets its own file (SA1402 enforced as error)
- **`sealed`** by default on all classes
- **File-scoped namespaces** everywhere (enforced via `.editorconfig` as error)
- **CancellationToken** in all async methods
- **Nullable reference types** enabled globally
- **Warnings as errors** — `TreatWarningsAsErrors` is true in `Directory.Build.props`
- **Naming** — PascalCase types, `_camelCase` private fields, `s_camelCase` private static fields, `I`-prefixed interfaces, `T`-prefixed type params
- **Command/Handler pattern** — `ICommand` + `ICommandHandler<T>` in Core
- **DI registration** via `IServiceCollection` extension methods per layer:
  - `AddCoreServices()` in `Avro.Cli.Core`
  - `AddCliServices()` in `Avro.Cli`
  - Future modules: `AddGitCommands()`, `AddDockerCommands()`, etc.
- **Package management** — add version to `Directory.Packages.props` first, then `<PackageReference Include="..." />` (no version attribute) in `.csproj`

## CI/CD

- **ci.yml** — builds on push/PR to main; Ubuntu 24.04, .NET 9 SDK, GitVersion
- **release.yml** — on push to main; multi-platform publish (win-x64, linux-x64, linux-arm64, osx-x64, osx-arm64); creates GitHub Release with artifacts
- Output: self-contained single-file executables, assembly name `avro`

## Verification

- After code changes: `dotnet build Avro.Cli.slnx` — must be 0 warnings, 0 errors
- After adding packages: add version to `Directory.Packages.props` first
- No test projects yet — verify manually via build
- Never claim "done" without running build verification
