using Avro.Cli.Core.Themes;

namespace Avro.Cli.Themes;

public sealed class TerminalGuiThemeApplicator : IThemeApplicator
{
    private View? _rootView;
    private ColorScheme? _baseScheme;
    private ColorScheme? _menuScheme;
    private ColorScheme? _dialogScheme;
    private ColorScheme? _topLevelScheme;

    public void SetRootView(View rootView)
    {
        _rootView = rootView;
    }

    public void ApplyTheme(ThemeDefinition theme)
    {
        Console.WriteLine($"[DEBUG] ApplyTheme called: {theme.Name}");
        
        // Terminal.Gui v2 supports True Color (24-bit RGB)!
        _baseScheme = CreateColorScheme(theme.Base);
        _menuScheme = CreateColorScheme(theme.Menu);
        _dialogScheme = CreateColorScheme(theme.Dialog);
        _topLevelScheme = CreateColorScheme(theme.TopLevel);
        
        Console.WriteLine($"[DEBUG] ColorSchemes created. Base normal fg: {_baseScheme.Normal.Foreground.R},{_baseScheme.Normal.Foreground.G},{_baseScheme.Normal.Foreground.B}");
        
        // Update all existing views
        if (_rootView != null)
        {
            Console.WriteLine($"[DEBUG] Updating view colors. Root view type: {_rootView.GetType().Name}");
            UpdateViewColors(_rootView);
            Console.WriteLine("[DEBUG] UpdateViewColors completed");
        }
        else
        {
            Console.WriteLine("[DEBUG] WARNING: _rootView is null!");
        }
    }
    
    private void UpdateViewColors(View? view)
    {
        if (view == null) return;
        
        // Apply color schemes based on view type
        if (_baseScheme != null && _menuScheme != null && _dialogScheme != null && _topLevelScheme != null)
        {
            var oldScheme = view.ColorScheme;
            view.ColorScheme = view switch
            {
                MenuBar => _menuScheme,
                Dialog => _dialogScheme,
                Window => _topLevelScheme,
                _ => _baseScheme
            };
            
            Console.WriteLine($"[DEBUG] View {view.GetType().Name}: old scheme={oldScheme?.Normal.Foreground.R}, new scheme={view.ColorScheme.Normal.Foreground.R}");
            
            // Force redraw after color change
            view.SetNeedsDraw();
        }
        
        // Recursively update children
        foreach (var subview in view.Subviews)
        {
            UpdateViewColors(subview);
        }
    }

    private static ColorScheme CreateColorScheme(ColorSchemeDefinition def)
    {
        return new ColorScheme
        {
            Normal = CreateAttribute(def.Normal, def.Background),
            Focus = CreateAttribute(def.Focus, def.Background),
            HotNormal = CreateAttribute(def.HotNormal, def.Background),
            HotFocus = CreateAttribute(def.HotFocus, def.Background),
            Disabled = CreateAttribute(def.Disabled, def.Background)
        };
    }

    private static Terminal.Gui.Attribute CreateAttribute(ThemeColor fg, ThemeColor bg)
    {
        // Terminal.Gui v2 True Color (24-bit RGB)
        var foreground = new Color(fg.R, fg.G, fg.B);
        var background = new Color(bg.R, bg.G, bg.B);
        
        return new Terminal.Gui.Attribute(foreground, background);
    }
}
