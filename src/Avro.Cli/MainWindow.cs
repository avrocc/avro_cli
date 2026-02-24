using Avro.Cli.Core.Themes;
using Avro.Cli.Themes;

namespace Avro.Cli;

public static class MainWindow
{
    public static void Configure(Window window, IThemeManager themeManager, IThemeApplicator themeApplicator)
    {
        if (window == null)
        {
            throw new ArgumentNullException(nameof(window), "Window is null!");
        }
        
        // MenuBar
        var menu = new MenuBar
        {
            Menus =
            [
                new MenuBarItem("_File", new MenuItem[]
                {
                    new ("_Quit", "", () => Application.RequestStop(), null, null, KeyCode.Q | KeyCode.CtrlMask)
                }),
                new MenuBarItem("_Appearance", new MenuItem[]
                {
                    new MenuBarItem("_Themes", BuildThemeSubmenus(themeManager, themeApplicator))
                })
            ]
        };

        // Main label
        var label = new Label
        {
            Text = "Avro CLI - Your productivity companion",
            X = Pos.Center(),
            Y = Pos.Center()
        };

        window.Add(menu);
        window.Add(label);
    }

    private static MenuItem[] BuildThemeSubmenus(IThemeManager themeManager, IThemeApplicator themeApplicator)
    {
        var themes = themeManager.AvailableThemes;
        var darkThemes = themes.Where(t => t.Category == ThemeCategory.Dark).OrderBy(t => t.Name).ToArray();
        var lightThemes = themes.Where(t => t.Category == ThemeCategory.Light).OrderBy(t => t.Name).ToArray();

        var items = new List<MenuItem>();

        // Dark section header
        if (darkThemes.Length > 0)
        {
            items.Add(new MenuItem("── Dark ──", "", null) { CanExecute = () => false });
            items.AddRange(darkThemes.Select(t => CreateThemeMenuItem(t, themeManager, themeApplicator)));
        }

        // Light section header
        if (lightThemes.Length > 0)
        {
            items.Add(new MenuItem("── Light ──", "", null) { CanExecute = () => false });
            items.AddRange(lightThemes.Select(t => CreateThemeMenuItem(t, themeManager, themeApplicator)));
        }

        return items.ToArray();
    }

    private static MenuItem CreateThemeMenuItem(ThemeDefinition theme, IThemeManager themeManager, IThemeApplicator themeApplicator)
    {
        return new MenuItem(theme.Name, "", () =>
        {
            themeManager.SetTheme(theme.Name);
            themeApplicator.ApplyTheme(theme);
        });
    }
}
