namespace Avro.Cli.Core.Themes;

public enum ThemeCategory
{
    Dark,
    Light
}

public sealed record ThemeDefinition
{
    public required string Name { get; init; }
    public required string Author { get; init; }
    public ThemeCategory Category { get; init; } = ThemeCategory.Dark;
    public required ColorSchemeDefinition Base { get; init; }
    public required ColorSchemeDefinition Menu { get; init; }
    public required ColorSchemeDefinition Dialog { get; init; }
    public required ColorSchemeDefinition Error { get; init; }
    public required ColorSchemeDefinition TopLevel { get; init; }
}
