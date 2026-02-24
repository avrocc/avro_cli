using Avro.Cli.Themes;

namespace Avro.Cli;

public static class DependencyInjection
{
    public static IServiceCollection AddCliServices(this IServiceCollection services)
    {
        services.AddLogging();
        services.AddSingleton<IThemeApplicator, TerminalGuiThemeApplicator>();
        return services;
    }
}
