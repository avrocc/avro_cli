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
Application.UseSystemConsole = true;   // REQUIRED on macOS — fixes dropdown rendering
Application.Init();
MainWindow.Configure(Application.Top); // configure Application.Top directly
Application.Run();
Application.Shutdown();
```

- **`UseSystemConsole = true`** — switches from CursesDriver to NetDriver; without this, box-drawing characters render double-width on macOS and dropdowns break
- `Application.Top` — use directly, do NOT subclass `Toplevel`
- `MainWindow` — static class with `Configure(Toplevel top)` method
- `MenuBar` — horizontal menu with `MenuBarItem[]` / `MenuItem[]`
- `StatusBar` — bottom bar with `StatusItem[]`
- **Avoid `null` separators** in menus — causes rendering artifacts; use separate top-level menus instead

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
