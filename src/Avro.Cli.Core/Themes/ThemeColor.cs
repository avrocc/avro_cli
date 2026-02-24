namespace Avro.Cli.Core.Themes;

public sealed record ThemeColor(byte R, byte G, byte B)
{
    public static ThemeColor FromHex(string hex)
    {
        hex = hex.TrimStart('#');
        return new ThemeColor(
            Convert.ToByte(hex.Substring(0, 2), 16),
            Convert.ToByte(hex.Substring(2, 2), 16),
            Convert.ToByte(hex.Substring(4, 2), 16)
        );
    }

    public string ToHex() => $"#{R:X2}{G:X2}{B:X2}";
}
