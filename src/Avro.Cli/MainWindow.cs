namespace Avro.Cli;

public static class MainWindow
{
    public static void Configure(Toplevel top)
    {
        ApplyTheme();
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

    private static void ApplyTheme()
    {
        var driver = Application.Driver;

        Colors.Base = new ColorScheme
        {
            Normal = driver.MakeAttribute(Color.Gray, Color.Black),
            Focus = driver.MakeAttribute(Color.White, Color.Black),
            HotNormal = driver.MakeAttribute(Color.BrightMagenta, Color.Black),
            HotFocus = driver.MakeAttribute(Color.BrightMagenta, Color.Black),
            Disabled = driver.MakeAttribute(Color.DarkGray, Color.Black)
        };

        Colors.Menu = new ColorScheme
        {
            Normal = driver.MakeAttribute(Color.White, Color.DarkGray),
            Focus = driver.MakeAttribute(Color.Black, Color.Gray),
            HotNormal = driver.MakeAttribute(Color.BrightMagenta, Color.DarkGray),
            HotFocus = driver.MakeAttribute(Color.BrightMagenta, Color.Gray),
            Disabled = driver.MakeAttribute(Color.DarkGray, Color.DarkGray)
        };

        Colors.Dialog = new ColorScheme
        {
            Normal = driver.MakeAttribute(Color.White, Color.DarkGray),
            Focus = driver.MakeAttribute(Color.Black, Color.Gray),
            HotNormal = driver.MakeAttribute(Color.BrightCyan, Color.DarkGray),
            HotFocus = driver.MakeAttribute(Color.BrightCyan, Color.Gray),
            Disabled = driver.MakeAttribute(Color.DarkGray, Color.DarkGray)
        };

        Colors.Error = new ColorScheme
        {
            Normal = driver.MakeAttribute(Color.White, Color.Red),
            Focus = driver.MakeAttribute(Color.BrightYellow, Color.Red),
            HotNormal = driver.MakeAttribute(Color.BrightYellow, Color.Red),
            HotFocus = driver.MakeAttribute(Color.BrightYellow, Color.BrightRed),
            Disabled = driver.MakeAttribute(Color.Gray, Color.Red)
        };

        Colors.TopLevel = new ColorScheme
        {
            Normal = driver.MakeAttribute(Color.Gray, Color.Black),
            Focus = driver.MakeAttribute(Color.White, Color.Black),
            HotNormal = driver.MakeAttribute(Color.BrightMagenta, Color.Black),
            HotFocus = driver.MakeAttribute(Color.BrightMagenta, Color.Black),
            Disabled = driver.MakeAttribute(Color.DarkGray, Color.Black)
        };
    }

    private static void SetStatus(StatusItem item, string message)
    {
        item.Title = message;
        Application.Refresh();
    }
}
