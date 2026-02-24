using Avro.Cli;
using Avro.Cli.Core.Themes;
using Avro.Cli.Themes;

var services = new ServiceCollection()
    .AddCoreServices()
    .AddCliServices()
    .BuildServiceProvider();

var themeManager = services.GetRequiredService<IThemeManager>();
var themeApplicator = services.GetRequiredService<IThemeApplicator>();

Application.UseSystemConsole = true;
Application.Init();

themeApplicator.ApplyTheme(themeManager.CurrentTheme);
themeManager.ThemeChanged += (_, theme) => themeApplicator.ApplyTheme(theme);

MainWindow.Configure(Application.Top);
Application.Run();
Application.Shutdown();
