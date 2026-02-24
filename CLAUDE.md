# Avro CLI

.NET 9 Terminal.Gui v1 TUI application.

## Tech Stack

- .NET 9 / Terminal.Gui `1.19.0`
- Microsoft.Extensions.DependencyInjection
- Central Package Management (`Directory.Packages.props`)
- Solution format: `.slnx` (XML)

## Architecture

```
src/
  Avro.Cli/          # TUI entry point, DI, windows/views
  Avro.Cli.Core/     # Commands, handlers, interfaces (zero UI dependency)
```

Growth path: Core splits into Domain + Application + Infrastructure when needed.

## Commands

```bash
dotnet build Avro.Cli.slnx                    # build (warnings as errors)
dotnet run --project src/Avro.Cli              # run TUI
task deploy                                    # publish + deploy locally
```

## Terminal.Gui v1 API

Single namespace `Terminal.Gui`. Static application lifecycle:

```csharp
Application.Init();
Application.Run(new MainWindow());
Application.Shutdown();
```

- `Toplevel` — root view (MainWindow inherits from this)
- `Window` — framed container added inside Toplevel
- `MenuBar` — horizontal menu with `MenuBarItem[]` / `MenuItem[]`
- `StatusBar` — bottom bar with `StatusItem[]`
- Separators in menus: `null` entry

## Conventions

- **1 type = 1 file** — every class, interface, struct, record gets its own file
- **`sealed`** by default on all classes
- **File-scoped namespaces** everywhere (enforced via `.editorconfig` as error)
- **CancellationToken** in all async methods
- **Command/Handler pattern** — `ICommand` + `ICommandHandler<T>` in Core
- **DI registration** via `IServiceCollection` extension methods per layer:
  - `AddCoreServices()` in `Avro.Cli.Core`
  - `AddCliServices()` in `Avro.Cli`
  - Future modules: `AddGitCommands()`, `AddDockerCommands()`, etc.

## Verification

- After code changes: `dotnet build Avro.Cli.slnx` — must be 0 warnings, 0 errors
- After adding packages: add version to `Directory.Packages.props` first
- Never claim "done" without running build verification
