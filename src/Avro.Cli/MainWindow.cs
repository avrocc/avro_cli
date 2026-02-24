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
                    new ("_Themes", "", () => ShowThemeSelector(window, themeManager, themeApplicator))
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

    private static void ShowThemeSelector(Window parentWindow, IThemeManager themeManager, IThemeApplicator themeApplicator)
    {
        var dialog = new Dialog
        {
            Title = "Select Theme",
            Width = Dim.Percent(60),
            Height = Dim.Percent(60)
        };

        var themes = themeManager.AvailableThemes.ToArray();
        var themeNames = themes.Select(t => t.Name).ToList();
        var currentIndex = Array.FindIndex(themes, t => t.Name == themeManager.CurrentTheme.Name);
        if (currentIndex < 0) currentIndex = 0;

        var listView = new ListView
        {
            X = 1,
            Y = 1,
            Width = Dim.Fill(2),
            Height = Dim.Fill(3),
            CanFocus = true,
            Source = new ListWrapper<string>(new System.Collections.ObjectModel.ObservableCollection<string>(themeNames))
        };
        listView.SelectedItem = currentIndex;

        // Live preview on arrow key navigation
        listView.SelectedItemChanged += (sender, args) =>
        {
            if (args.Item >= 0 && args.Item < themes.Length)
            {
                themeApplicator.ApplyTheme(themes[args.Item]);
            }
        };

        // Enter key = apply and close
        listView.OpenSelectedItem += (sender, args) =>
        {
            var selectedTheme = themes[listView.SelectedItem];
            themeManager.SetTheme(selectedTheme.Name);
            Application.RequestStop();
        };

        var cancelButton = new Button
        {
            Text = "Cancel",
            X = Pos.Center(),
            Y = Pos.AnchorEnd(1)
        };

        cancelButton.Accepting += (sender, args) =>
        {
            themeApplicator.ApplyTheme(themeManager.CurrentTheme);
            Application.RequestStop();
        };

        dialog.Add(listView, cancelButton);
        Application.Run(dialog);
        dialog.Dispose();
    }
}
