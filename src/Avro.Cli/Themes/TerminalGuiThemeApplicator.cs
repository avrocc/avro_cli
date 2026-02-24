namespace Avro.Cli.Themes;

public sealed class TerminalGuiThemeApplicator : IThemeApplicator
{
    public void ApplyTheme(Core.Themes.ThemeDefinition theme)
    {
        var driver = Application.Driver;

        Colors.Base = CreateColorScheme(driver, theme.Base);
        Colors.Menu = CreateColorScheme(driver, theme.Menu);
        Colors.Dialog = CreateColorScheme(driver, theme.Dialog);
        Colors.Error = CreateColorScheme(driver, theme.Error);
        Colors.TopLevel = CreateColorScheme(driver, theme.TopLevel);
    }

    private static ColorScheme CreateColorScheme(ConsoleDriver driver, Core.Themes.ColorSchemeDefinition def)
    {
        return new ColorScheme
        {
            Normal = driver.MakeAttribute(MapColor(def.Normal), MapColor(def.Background)),
            Focus = driver.MakeAttribute(MapColor(def.Focus), MapColor(def.Background)),
            HotNormal = driver.MakeAttribute(MapColor(def.HotNormal), MapColor(def.Background)),
            HotFocus = driver.MakeAttribute(MapColor(def.HotFocus), MapColor(def.Background)),
            Disabled = driver.MakeAttribute(MapColor(def.Disabled), MapColor(def.Background))
        };
    }

    private static Color MapColor(Core.Themes.ThemeColor color)
    {
        var brightness = (color.R + color.G + color.B) / 3;
        
        if (brightness < 32) return Color.Black;
        if (brightness > 224) return Color.White;
        
        if (color.R > color.G && color.R > color.B)
            return brightness > 128 ? Color.BrightRed : Color.Red;
        
        if (color.G > color.R && color.G > color.B)
            return brightness > 128 ? Color.BrightGreen : Color.Green;
        
        if (color.B > color.R && color.B > color.G)
            return brightness > 128 ? Color.BrightBlue : Color.Blue;
        
        if (color.R > 100 && color.G > 100 && color.B < 100)
            return Color.BrightYellow;
        
        if (color.R > 100 && color.B > 100 && color.G < 100)
            return Color.BrightMagenta;
        
        if (color.G > 100 && color.B > 100 && color.R < 100)
            return Color.BrightCyan;
        
        return brightness > 128 ? Color.Gray : Color.DarkGray;
    }
}
