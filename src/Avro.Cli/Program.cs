using Avro.Cli;
using Avro.Cli.Core.Themes;
using Avro.Cli.Themes;

var services = new ServiceCollection()
    .AddCoreServices()
    .AddCliServices()
    .BuildServiceProvider();

var themeManager = services.GetRequiredService<IThemeManager>();
var themeApplicator = services.GetRequiredService<IThemeApplicator>();

// Load themes from themes directory
var themesDirectory = Path.Combine(AppContext.BaseDirectory, "..", "..", "..", "..", "..", "themes");
themesDirectory = Path.GetFullPath(themesDirectory);

if (Directory.Exists(themesDirectory))
{
    themeManager.LoadThemes(themesDirectory);
}

Application.Init();

// Initialize color schemes AFTER Application.Init()
themeApplicator.ApplyTheme(themeManager.CurrentTheme);
themeManager.ThemeChanged += (_, theme) => themeApplicator.ApplyTheme(theme);

MainWindow.Configure(Application.Top!, themeManager, themeApplicator);
Application.Run();
Application.Shutdown();
