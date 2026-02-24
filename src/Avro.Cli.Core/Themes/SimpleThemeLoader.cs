namespace Avro.Cli.Core.Themes;

public sealed class SimpleThemeLoader : IThemeLoader
{
    public ThemeDefinition LoadTheme(string filePath)
    {
        var lines = File.ReadAllLines(filePath);
        var data = new Dictionary<string, string>(StringComparer.OrdinalIgnoreCase);
        var sections = new Dictionary<string, Dictionary<string, string>>(StringComparer.OrdinalIgnoreCase);
        var currentSection = "";

        foreach (var rawLine in lines)
        {
            var line = rawLine.Trim();
            if (string.IsNullOrEmpty(line) || line.StartsWith('#'))
                continue;

            if (line.StartsWith('[') && line.EndsWith(']'))
            {
                currentSection = line[1..^1].Trim();
                sections[currentSection] = new Dictionary<string, string>(StringComparer.OrdinalIgnoreCase);
                continue;
            }

            var parts = line.Split('=', 2);
            if (parts.Length != 2)
                continue;

            var key = parts[0].Trim();
            var value = parts[1].Trim();

            if (string.IsNullOrEmpty(currentSection))
                data[key] = value;
            else
                sections[currentSection][key] = value;
        }

        return new ThemeDefinition
        {
            Name = data.GetValueOrDefault("name", "Unknown"),
            Author = data.GetValueOrDefault("author", "Unknown"),
            Base = ParseColorScheme(sections.GetValueOrDefault("base", new())),
            Menu = ParseColorScheme(sections.GetValueOrDefault("menu", new())),
            Dialog = ParseColorScheme(sections.GetValueOrDefault("dialog", new())),
            Error = ParseColorScheme(sections.GetValueOrDefault("error", new())),
            TopLevel = ParseColorScheme(sections.GetValueOrDefault("toplevel", new()))
        };
    }

    public IReadOnlyList<ThemeDefinition> LoadAllThemes(string directory)
    {
        if (!Directory.Exists(directory))
            return Array.Empty<ThemeDefinition>();

        var themeFiles = Directory.GetFiles(directory, "*.theme");
        var themes = new List<ThemeDefinition>();

        foreach (var file in themeFiles)
        {
            try
            {
                var theme = LoadTheme(file);
                themes.Add(theme);
            }
            catch
            {
                // Skip invalid theme files
            }
        }

        return themes;
    }

    private static ColorSchemeDefinition ParseColorScheme(Dictionary<string, string> section)
    {
        return new ColorSchemeDefinition
        {
            Normal = ParseColor(section.GetValueOrDefault("normal", "#808080")),
            Focus = ParseColor(section.GetValueOrDefault("focus", "#ffffff")),
            HotNormal = ParseColor(section.GetValueOrDefault("hot_normal", "#ff00ff")),
            HotFocus = ParseColor(section.GetValueOrDefault("hot_focus", "#ff00ff")),
            Disabled = ParseColor(section.GetValueOrDefault("disabled", "#404040")),
            Background = ParseColor(section.GetValueOrDefault("background", "#000000"))
        };
    }

    private static ThemeColor ParseColor(string hex)
    {
        return ThemeColor.FromHex(hex);
    }
}
