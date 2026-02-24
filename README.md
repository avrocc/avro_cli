# Avro CLI

Terminal UI application built with [Terminal.Gui](https://github.com/gui-cs/Terminal.Gui) v1 and .NET 9.

## Prerequisites

- [.NET 9 SDK](https://dotnet.microsoft.com/download/dotnet/9.0)

## Build & Run

```bash
dotnet build Avro.Cli.slnx
dotnet run --project src/Avro.Cli
```

Press **Esc** to quit.

## Project Structure

```
src/
  Avro.Cli/          # TUI entry point, DI registration
  Avro.Cli.Core/     # Commands, handlers, interfaces (no UI dependency)
```
