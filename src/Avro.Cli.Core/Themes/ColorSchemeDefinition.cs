namespace Avro.Cli.Core.Themes;

public sealed record ColorSchemeDefinition
{
    public required ThemeColor Normal { get; init; }
    public required ThemeColor Focus { get; init; }
    public required ThemeColor HotNormal { get; init; }
    public required ThemeColor HotFocus { get; init; }
    public required ThemeColor Disabled { get; init; }
    public required ThemeColor Background { get; init; }
}
