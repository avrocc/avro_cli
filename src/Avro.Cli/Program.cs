using Avro.Cli;
using Avro.Cli.Core.Themes;
using Avro.Cli.Themes;

var services = new ServiceCollection()
    .AddCoreServices()
    .AddCliServices()
    .BuildServiceProvider();

var themeManager = services.GetRequiredService<IThemeManager>();
var themeApplicator = (TerminalGuiThemeApplicator)services.GetRequiredService<IThemeApplicator>();

// Load themes
var themesDirectory = Path.Combine(AppContext.BaseDirectory, "..", "..", "..", "..", "..", "themes");
themesDirectory = Path.GetFullPath(themesDirectory);

if (Directory.Exists(themesDirectory))
{
    themeManager.LoadThemes(themesDirectory);
}

// v2 legacy static pattern (still works)
Application.Init();

// Set root view for theme applicator - use Application.Top
themeApplicator.SetRootView(Application.Top!);

MainWindow.Configure(Application.Top!, themeManager, themeApplicator);

// Apply theme after UI is configured
themeApplicator.ApplyTheme(themeManager.CurrentTheme);
themeManager.ThemeChanged += (_, theme) => themeApplicator.ApplyTheme(theme);

Application.Run();
Application.Shutdown();
