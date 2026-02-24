namespace Avro.Cli.Themes;

public sealed class TerminalGuiThemeApplicator : IThemeApplicator
{
    public void ApplyTheme(Core.Themes.ThemeDefinition theme)
    {
        // Terminal.Gui v2 supports True Color!
        var baseScheme = CreateColorScheme(theme.Base);
        var menuScheme = CreateColorScheme(theme.Menu);
        var dialogScheme = CreateColorScheme(theme.Dialog);
        var errorScheme = CreateColorScheme(theme.Error);
        var topLevelScheme = CreateColorScheme(theme.TopLevel);
        
        Colors.ColorSchemes["Base"] = baseScheme;
        Colors.ColorSchemes["Menu"] = menuScheme;
        Colors.ColorSchemes["Dialog"] = dialogScheme;
        Colors.ColorSchemes["Error"] = errorScheme;
        Colors.ColorSchemes["TopLevel"] = topLevelScheme;
        
        // Update all existing views
        UpdateViewColors(Application.Top);
    }
    
    private static void UpdateViewColors(View? view)
    {
        if (view == null) return;
        
        // Apply color schemes based on view type
        if (view is MenuBar)
            view.ColorScheme = Colors.ColorSchemes["Menu"];
        else if (view is Dialog)
            view.ColorScheme = Colors.ColorSchemes["Dialog"];
        else if (view is Toplevel)
            view.ColorScheme = Colors.ColorSchemes["TopLevel"];
        else
            view.ColorScheme = Colors.ColorSchemes["Base"];
        
        // Recursively update children
        foreach (var subview in view.Subviews)
        {
            UpdateViewColors(subview);
        }
    }

    private static ColorScheme CreateColorScheme(Core.Themes.ColorSchemeDefinition def)
    {
        return new ColorScheme
        {
            Normal = new Terminal.Gui.Attribute(MapColor(def.Normal), MapColor(def.Background)),
            Focus = new Terminal.Gui.Attribute(MapColor(def.Focus), MapColor(def.Background)),
            HotNormal = new Terminal.Gui.Attribute(MapColor(def.HotNormal), MapColor(def.Background)),
            HotFocus = new Terminal.Gui.Attribute(MapColor(def.HotFocus), MapColor(def.Background)),
            Disabled = new Terminal.Gui.Attribute(MapColor(def.Disabled), MapColor(def.Background))
        };
    }

    private static Color MapColor(Core.Themes.ThemeColor color)
    {
        // Terminal.Gui v2 True Color support
        return new Color(color.R, color.G, color.B);
    }
}
