using Avro.Cli;

var services = new ServiceCollection()
    .AddCoreServices()
    .AddCliServices()
    .BuildServiceProvider();

Application.UseSystemConsole = true;
Application.Init();
MainWindow.Configure(Application.Top);
Application.Run();
Application.Shutdown();
