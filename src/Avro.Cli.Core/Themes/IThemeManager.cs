namespace Avro.Cli.Core.Themes;

public interface IThemeManager
{
    ThemeDefinition CurrentTheme { get; }
    IReadOnlyList<ThemeDefinition> AvailableThemes { get; }
    void LoadThemes(string themesDirectory);
    void SetTheme(string themeName);
    event EventHandler<ThemeDefinition>? ThemeChanged;
}
