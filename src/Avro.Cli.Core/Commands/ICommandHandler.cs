namespace Avro.Cli.Core.Commands;

public interface ICommandHandler<in TCommand> where TCommand : ICommand
{
    Task ExecuteAsync(TCommand command, CancellationToken cancellationToken);
}
