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
        // Terminal.Gui v2 supports True Color (24-bit RGB)!
        _baseScheme = CreateColorScheme(theme.Base);
        _menuScheme = CreateColorScheme(theme.Menu);
        _dialogScheme = CreateColorScheme(theme.Dialog);
        _topLevelScheme = CreateColorScheme(theme.TopLevel);
        
        // Update all existing views
        if (_rootView != null)
        {
            UpdateViewColors(_rootView);
        }
    }
    
    private void UpdateViewColors(View? view)
    {
        if (view == null) return;
        
        // Apply color schemes based on view type
        if (_baseScheme != null && _menuScheme != null && _dialogScheme != null && _topLevelScheme != null)
        {
            view.ColorScheme = view switch
            {
                MenuBar => _menuScheme,
                Dialog => _dialogScheme,
                Window => _topLevelScheme,
                _ => _baseScheme
            };
            
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
            Focus = CreateAttribute(def.Background, def.Focus),
            HotNormal = CreateAttribute(def.HotNormal, def.Background),
            HotFocus = CreateAttribute(def.Background, def.HotFocus),
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
