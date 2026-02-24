using Avro.Cli.Core.Themes;

namespace Avro.Cli.Core;

public static class DependencyInjection
{
    public static IServiceCollection AddCoreServices(this IServiceCollection services)
    {
        services.AddSingleton<IThemeLoader, SimpleThemeLoader>();
        services.AddSingleton<IThemeManager, ThemeManager>();
        return services;
    }
}
