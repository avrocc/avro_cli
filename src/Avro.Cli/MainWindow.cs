using Avro.Cli.Core.Themes;
using Avro.Cli.Themes;

namespace Avro.Cli;

public static class MainWindow
{
    public static void Configure(Toplevel top, IThemeManager themeManager, IThemeApplicator themeApplicator)
    {
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
                    new ("_Themes", "", () => ShowThemeSelector(top, themeManager, themeApplicator))
                })
            ]
        };

        // StatusBar
        var statusBar = new StatusBar();
        statusBar.Add(new Shortcut(KeyCode.Q | KeyCode.CtrlMask, "Quit", () => Application.RequestStop()));

        // Main label
        var label = new Label
        {
            Text = "Avro CLI - Your productivity companion",
            X = Pos.Center(),
            Y = Pos.Center()
        };

        top.Add(menu);
        top.Add(statusBar);
        top.Add(label);
    }

    private static void ShowThemeSelector(Toplevel parentWindow, IThemeManager themeManager, IThemeApplicator themeApplicator)
    {
        var dialog = new Dialog
        {
            Title = "Select Theme",
            Width = Dim.Percent(60),
            Height = Dim.Percent(60)
        };

        var themes = themeManager.AvailableThemes.ToArray();
        var themeNames = themes.Select(t => t.Name).ToArray();
        var currentIndex = Array.FindIndex(themes, t => t.Name == themeManager.CurrentTheme.Name);
        if (currentIndex < 0) currentIndex = 0;

        var radioGroup = new RadioGroup
        {
            X = 1,
            Y = 1,
            Width = Dim.Fill(2),
            Height = Dim.Fill(2),
            RadioLabels = themeNames
        };

        if (currentIndex >= 0 && currentIndex < themeNames.Length)
        {
            radioGroup.SelectedItem = currentIndex;
        }

        // Live preview on selection change
        radioGroup.SelectedItemChanged += (sender, args) =>
        {
            var selectedTheme = themes[args.SelectedItem];
            themeApplicator.ApplyTheme(selectedTheme);
        };

        var okButton = new Button
        {
            Text = "OK",
            X = Pos.Center() - 10,
            Y = Pos.AnchorEnd(1),
            IsDefault = true
        };

        okButton.Accepting += (sender, args) =>
        {
            var selectedTheme = themes[radioGroup.SelectedItem];
            themeManager.SetTheme(selectedTheme.Name);
            Application.RequestStop();
        };

        var cancelButton = new Button
        {
            Text = "Cancel",
            X = Pos.Center() + 2,
            Y = Pos.AnchorEnd(1)
        };

        cancelButton.Accepting += (sender, args) =>
        {
            // Revert to original theme
            themeApplicator.ApplyTheme(themeManager.CurrentTheme);
            Application.RequestStop();
        };

        dialog.Add(radioGroup, okButton, cancelButton);
        Application.Run(dialog);
        dialog.Dispose();
    }
}
