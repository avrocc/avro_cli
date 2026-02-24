namespace Avro.Cli;

public static class DependencyInjection
{
    public static IServiceCollection AddCliServices(this IServiceCollection services)
    {
        services.AddLogging();
        return services;
    }
}
