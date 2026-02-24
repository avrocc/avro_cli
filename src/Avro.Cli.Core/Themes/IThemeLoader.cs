namespace Avro.Cli.Core.Themes;

public interface IThemeLoader
{
    ThemeDefinition LoadTheme(string filePath);
    IReadOnlyList<ThemeDefinition> LoadAllThemes(string directory);
}
