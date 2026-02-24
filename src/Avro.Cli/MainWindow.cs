using Avro.Cli.Core.Themes;

namespace Avro.Cli;

public static class MainWindow
{
    public static void Configure(Toplevel top, IThemeManager themeManager, Themes.IThemeApplicator themeApplicator)
    {
        top.ColorScheme = Colors.Base;

        var statusItem = new StatusItem(Key.Null, "Ready", null);

        var menuBar = new MenuBar(
        [
            new MenuBarItem("_File",
            [
                new MenuItem("_New Session", "", () => SetStatus(statusItem, "New Session — coming soon")),
                new MenuItem("_Open...", "", () => SetStatus(statusItem, "Open — coming soon")),
                new MenuItem("_Save", "", () => SetStatus(statusItem, "Save — coming soon")),
                new MenuItem("_Quit", "", () => Application.RequestStop())
            ]),
            new MenuBarItem("_Git",
            [
                new MenuItem("_Status", "", () => SetStatus(statusItem, "Git Status — coming soon")),
                new MenuItem("_Log", "", () => SetStatus(statusItem, "Git Log — coming soon"))
            ]),
            new MenuBarItem("_Docker",
            [
                new MenuItem("_Containers", "", () => SetStatus(statusItem, "Docker Containers — coming soon")),
                new MenuItem("_Images", "", () => SetStatus(statusItem, "Docker Images — coming soon"))
            ]),
            new MenuBarItem("_SSH",
            [
                new MenuItem("_Connect", "", () => SetStatus(statusItem, "SSH Connect — coming soon"))
            ]),
            new MenuBarItem("_K8s",
            [
                new MenuItem("_Pods", "", () => SetStatus(statusItem, "Kubernetes Pods — coming soon"))
            ]),
            new MenuBarItem("_Appearance",
            [
                new MenuItem("Select _Theme...", "", () => ShowThemeSelector(themeManager, themeApplicator, statusItem))
            ]),
            new MenuBarItem("_Help",
            [
                new MenuItem("_Documentation", "", () => SetStatus(statusItem, "Documentation — coming soon")),
                new MenuItem("Check for _Updates", "", () => SetStatus(statusItem, "Check for Updates — coming soon")),
                new MenuItem("_About", "", () =>
                    MessageBox.Query("About", "Avro CLI\n\nTerminal UI toolkit for DevOps", "_OK"))
            ])
        ]);

        var statusBar = new StatusBar(
        [
            statusItem,
            new StatusItem(Key.CtrlMask | Key.Q, "~^Q~ Quit", () => Application.RequestStop())
        ]);

        var label = new Label("Welcome to Avro CLI")
        {
            X = Pos.Center(),
            Y = Pos.Center()
        };

        top.Add(menuBar, label, statusBar);
    }

    private static void ShowThemeSelector(IThemeManager themeManager, Themes.IThemeApplicator themeApplicator, StatusItem statusItem)
    {
        var themes = themeManager.AvailableThemes.ToList();
        var currentIndex = themes.FindIndex(t => t.Name == themeManager.CurrentTheme.Name);
        var originalTheme = themeManager.CurrentTheme;
        
        var dialog = new Dialog("Select Theme", 50, Math.Min(themes.Count + 8, 20));
        
        var label = new Label("Choose theme (arrow keys preview):")
        {
            X = 1,
            Y = 1
        };
        
        var radioGroup = new RadioGroup(themes.Select(t => NStack.ustring.Make(t.Name)).ToArray())
        {
            X = 1,
            Y = 3,
            SelectedItem = currentIndex >= 0 ? currentIndex : 0
        };
        
        // Live preview on radio button change
        radioGroup.SelectedItemChanged += (prevIndex) =>
        {
            var selectedIndex = radioGroup.SelectedItem;
            if (selectedIndex >= 0 && selectedIndex < themes.Count)
            {
                var previewTheme = themes[selectedIndex];
                label.Text = $"Applying: {previewTheme.Name}...";
                themeApplicator.ApplyTheme(previewTheme);
                Application.Refresh();
                label.Text = $"Preview: {previewTheme.Name}";
            }
        };
        
        var okButton = new Button("_OK", is_default: true)
        {
            X = Pos.Center() - 8,
            Y = Pos.Bottom(dialog) - 3
        };
        
        okButton.Clicked += () =>
        {
            var selectedTheme = themes[radioGroup.SelectedItem];
            themeManager.SetTheme(selectedTheme.Name);
            themeApplicator.ApplyTheme(selectedTheme);
            SetStatus(statusItem, $"Theme: {selectedTheme.Name}");
            Application.RequestStop();
        };
        
        var cancelButton = new Button("_Cancel")
        {
            X = Pos.Center() + 3,
            Y = Pos.Bottom(dialog) - 3
        };
        
        cancelButton.Clicked += () =>
        {
            themeApplicator.ApplyTheme(originalTheme);
            Application.RequestStop();
        };
        
        dialog.Add(label, radioGroup, okButton, cancelButton);
        
        Application.Run(dialog);
    }

    private static MenuItem[] CreateThemeMenuItems(IThemeManager themeManager, Themes.IThemeApplicator themeApplicator, StatusItem statusItem)
    {
        var currentTheme = themeManager.CurrentTheme.Name;
        
        return themeManager.AvailableThemes
            .Select(theme =>
            {
                var isActive = theme.Name == currentTheme;
                var prefix = isActive ? "✓ " : "  ";
                
                return new MenuItem(
                    $"{prefix}{theme.Name}",
                    "",
                    () =>
                    {
                        themeManager.SetTheme(theme.Name);
                        themeApplicator.ApplyTheme(theme);
                        SetStatus(statusItem, $"Theme: {theme.Name}");
                        Application.Refresh();
                    });
            })
            .ToArray();
    }

    private static void SetStatus(StatusItem item, string message)
    {
        item.Title = message;
        Application.Refresh();
    }
}
