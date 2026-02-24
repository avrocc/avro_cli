using Avro.Cli.Core.Themes;

namespace Avro.Cli;

public static class MainWindow
{
    public static void Configure(Toplevel top, IThemeManager themeManager, Themes.IThemeApplicator themeApplicator)
    {
        // Don't set ColorScheme here - let ApplyTheme handle it

        var menuBar = new MenuBar
        {
            Menus =
            [
                new MenuBarItem("_File", new MenuItem[]
                {
                    new("_New Session", "", () => {}),
                    new("_Open...", "", () => {}),
                    new("_Save", "", () => {}),
                    new("_Quit", "", () => Application.RequestStop())
                }),
                new MenuBarItem("_Git", new MenuItem[]
                {
                    new("_Status", "", () => {}),
                    new("_Log", "", () => {})
                }),
                new MenuBarItem("_Docker", new MenuItem[]
                {
                    new("_Containers", "", () => {}),
                    new("_Images", "", () => {})
                }),
                new MenuBarItem("_SSH", new MenuItem[]
                {
                    new("_Connect", "", () => {})
                }),
                new MenuBarItem("_K8s", new MenuItem[]
                {
                    new("_Pods", "", () => {})
                }),
                new MenuBarItem("_Appearance", new MenuItem[]
                {
                    new("Select _Theme...", "", () => ShowThemeSelector(themeManager, themeApplicator))
                }),
                new MenuBarItem("_Help", new MenuItem[]
                {
                    new("_Documentation", "", () => {}),
                    new("Check for _Updates", "", () => {}),
                    new("_About", "", () =>
                        MessageBox.Query("About", "Avro CLI\n\nTerminal UI toolkit for DevOps", "OK"))
                })
            ]
        };

        var statusBar = new StatusBar();
        statusBar.Add(new Shortcut(Key.Q.WithCtrl, "Quit", () => Application.RequestStop()));

        var label = new Label
        {
            Text = "Welcome to Avro CLI",
            X = Pos.Center(),
            Y = Pos.Center()
        };

        top.Add(menuBar);
        top.Add(label);
        top.Add(statusBar);
    }

    private static void ShowThemeSelector(IThemeManager themeManager, Themes.IThemeApplicator themeApplicator)
    {
        var themes = themeManager.AvailableThemes.ToList();
        var currentIndex = themes.FindIndex(t => t.Name == themeManager.CurrentTheme.Name);
        var originalTheme = themeManager.CurrentTheme;
        
        var dialog = new Dialog
        {
            Title = "Select Theme",
            Width = 50,
            Height = Math.Min(themes.Count + 8, 20)
        };
        
        var label = new Label
        {
            Text = "Choose theme (arrow keys preview):",
            X = 1,
            Y = 1
        };
        
        var radioGroup = new RadioGroup
        {
            X = 1,
            Y = 3,
            RadioLabels = themes.Select(t => t.Name).ToArray(),
            SelectedItem = currentIndex >= 0 ? currentIndex : 0
        };
        
        // Live preview on radio button change
        radioGroup.SelectedItemChanged += (sender, args) =>
        {
            var selectedIndex = radioGroup.SelectedItem;
            if (selectedIndex >= 0 && selectedIndex < themes.Count)
            {
                var previewTheme = themes[selectedIndex];
                label.Text = $"Preview: {previewTheme.Name}";
                themeApplicator.ApplyTheme(previewTheme);
            }
        };
        
        var okButton = new Button
        {
            Text = "OK",
            X = Pos.Center() - 8,
            Y = Pos.AnchorEnd(2),
            IsDefault = true
        };
        
        okButton.Accepting += (sender, args) =>
        {
            var selectedTheme = themes[radioGroup.SelectedItem];
            themeManager.SetTheme(selectedTheme.Name);
            themeApplicator.ApplyTheme(selectedTheme);
            Application.RequestStop();
        };
        
        var cancelButton = new Button
        {
            Text = "Cancel",
            X = Pos.Center() + 3,
            Y = Pos.AnchorEnd(2)
        };
        
        cancelButton.Accepting += (sender, args) =>
        {
            themeApplicator.ApplyTheme(originalTheme);
            Application.RequestStop();
        };
        
        dialog.Add(label, radioGroup, okButton, cancelButton);
        
        Application.Run(dialog);
    }
}
