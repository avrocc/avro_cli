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

// Create window (Application.Top may be null in 2.0.0-prealpha)
var window = new Window { Title = "Avro CLI (Esc to quit)" };

// Set root view for theme applicator
themeApplicator.SetRootView(window);

MainWindow.Configure(window, themeManager, themeApplicator);

// Apply theme after UI is configured
themeApplicator.ApplyTheme(themeManager.CurrentTheme);
themeManager.ThemeChanged += (_, theme) => themeApplicator.ApplyTheme(theme);

Application.Run(window);

window.Dispose();
Application.Shutdown();
