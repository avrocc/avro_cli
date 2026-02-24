namespace Avro.Cli.Core.Themes;

public sealed class ThemeManager : IThemeManager
{
    private readonly IThemeLoader _themeLoader;
    private readonly List<ThemeDefinition> _themes = new();
    private ThemeDefinition _currentTheme;

    public ThemeManager(IThemeLoader themeLoader)
    {
        _themeLoader = themeLoader;
        _currentTheme = CreateDefaultTheme();
        _themes.Add(_currentTheme);
    }

    public ThemeDefinition CurrentTheme => _currentTheme;

    public IReadOnlyList<ThemeDefinition> AvailableThemes => _themes;

    public event EventHandler<ThemeDefinition>? ThemeChanged;

    public void LoadThemes(string themesDirectory)
    {
        var loadedThemes = _themeLoader.LoadAllThemes(themesDirectory);
        
        foreach (var theme in loadedThemes)
        {
            if (_themes.All(t => t.Name != theme.Name))
                _themes.Add(theme);
        }
    }

    public void SetTheme(string themeName)
    {
        var theme = _themes.FirstOrDefault(t => t.Name.Equals(themeName, StringComparison.OrdinalIgnoreCase));
        if (theme is null)
            throw new ArgumentException($"Theme '{themeName}' not found", nameof(themeName));

        _currentTheme = theme;
        ThemeChanged?.Invoke(this, _currentTheme);
    }

    private static ThemeDefinition CreateDefaultTheme()
    {
        return new ThemeDefinition
        {
            Name = "Default",
            Author = "Avro CLI",
            Base = new ColorSchemeDefinition
            {
                Normal = new ThemeColor(128, 128, 128),
                Focus = new ThemeColor(255, 255, 255),
                HotNormal = new ThemeColor(255, 0, 255),
                HotFocus = new ThemeColor(255, 0, 255),
                Disabled = new ThemeColor(64, 64, 64),
                Background = new ThemeColor(0, 0, 0)
            },
            Menu = new ColorSchemeDefinition
            {
                Normal = new ThemeColor(255, 255, 255),
                Focus = new ThemeColor(0, 0, 0),
                HotNormal = new ThemeColor(255, 0, 255),
                HotFocus = new ThemeColor(255, 0, 255),
                Disabled = new ThemeColor(64, 64, 64),
                Background = new ThemeColor(64, 64, 64)
            },
            Dialog = new ColorSchemeDefinition
            {
                Normal = new ThemeColor(255, 255, 255),
                Focus = new ThemeColor(0, 0, 0),
                HotNormal = new ThemeColor(0, 255, 255),
                HotFocus = new ThemeColor(0, 255, 255),
                Disabled = new ThemeColor(64, 64, 64),
                Background = new ThemeColor(64, 64, 64)
            },
            Error = new ColorSchemeDefinition
            {
                Normal = new ThemeColor(255, 255, 255),
                Focus = new ThemeColor(255, 255, 0),
                HotNormal = new ThemeColor(255, 255, 0),
                HotFocus = new ThemeColor(255, 255, 0),
                Disabled = new ThemeColor(128, 128, 128),
                Background = new ThemeColor(255, 0, 0)
            },
            TopLevel = new ColorSchemeDefinition
            {
                Normal = new ThemeColor(128, 128, 128),
                Focus = new ThemeColor(255, 255, 255),
                HotNormal = new ThemeColor(255, 0, 255),
                HotFocus = new ThemeColor(255, 0, 255),
                Disabled = new ThemeColor(64, 64, 64),
                Background = new ThemeColor(0, 0, 0)
            }
        };
    }
}
