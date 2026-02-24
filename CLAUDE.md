# Avro CLI

.NET 10 Terminal.Gui v2 TUI application.

## Tech Stack

- .NET 10 / Terminal.Gui `2.0.0-develop.5041`
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
```

## Terminal.Gui v2 API

The legacy static `Application.Init()`/`Application.Shutdown()` is obsolete. Use:

```csharp
using var app = Application.Create().Init();
app.Run(new MainWindow());
```

Key namespaces (not `Terminal.Gui` root):
- `Terminal.Gui.App` — `Application`, `IApplication`
- `Terminal.Gui.ViewBase` — `View`, `Pos`, `Dim`
- `Terminal.Gui.Views` — `Window`, `Label`, `Button`, etc.

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
