using Avro.Cli;

var services = new ServiceCollection()
    .AddCoreServices()
    .AddCliServices()
    .BuildServiceProvider();

using var app = Application.Create().Init();
app.Run(new MainWindow());
