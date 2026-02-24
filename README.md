# Avro CLI

Terminal UI application built with [Terminal.Gui](https://github.com/gui-cs/Terminal.Gui) v2 and .NET 10.

## Prerequisites

- [.NET 10 SDK](https://dotnet.microsoft.com/download/dotnet/10.0)

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
