package components

import "avro_cli/internal/tui/styles"

// StatusBar renders a bottom status bar with breadcrumb and help text.
func StatusBar(breadcrumb string, width int) string {
	text := " " + breadcrumb
	padding := width - len(text)
	if padding > 0 {
		for i := 0; i < padding; i++ {
			text += " "
		}
	}
	return styles.StatusBar.Render(text)
}
