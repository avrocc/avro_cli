package modules

// Blank imports trigger init() in each module, auto-registering commands.
import (
	_ "avro_cli/internal/modules/git"
	_ "avro_cli/internal/modules/http"
	_ "avro_cli/internal/modules/system"
)
