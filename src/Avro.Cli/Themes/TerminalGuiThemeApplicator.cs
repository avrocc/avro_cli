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
        
        // Ensure color schemes exist
        if (!Colors.ColorSchemes.ContainsKey("Base"))
            Colors.ColorSchemes.Add("Base", baseScheme);
        else
            Colors.ColorSchemes["Base"] = baseScheme;
            
        if (!Colors.ColorSchemes.ContainsKey("Menu"))
            Colors.ColorSchemes.Add("Menu", menuScheme);
        else
            Colors.ColorSchemes["Menu"] = menuScheme;
            
        if (!Colors.ColorSchemes.ContainsKey("Dialog"))
            Colors.ColorSchemes.Add("Dialog", dialogScheme);
        else
            Colors.ColorSchemes["Dialog"] = dialogScheme;
            
        if (!Colors.ColorSchemes.ContainsKey("Error"))
            Colors.ColorSchemes.Add("Error", errorScheme);
        else
            Colors.ColorSchemes["Error"] = errorScheme;
            
        if (!Colors.ColorSchemes.ContainsKey("TopLevel"))
            Colors.ColorSchemes.Add("TopLevel", topLevelScheme);
        else
            Colors.ColorSchemes["TopLevel"] = topLevelScheme;
        
        // Update all existing views
        if (Application.Top != null)
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
