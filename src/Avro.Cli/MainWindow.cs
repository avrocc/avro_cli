namespace Avro.Cli;

public sealed class MainWindow : Window
{
    public MainWindow()
    {
        Title = "Avro CLI (Esc to quit)";

        var label = new Label
        {
            Text = "Welcome to Avro CLI",
            X = Pos.Center(),
            Y = Pos.Center()
        };

        Add(label);
    }
}
