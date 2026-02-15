package http

import (
	"avro_cli/internal/app/registry"
	"avro_cli/internal/domain"
)

var category = domain.Category{
	Name:        "http",
	Description: "HTTP request utilities",
	Icon:        "\U0001F310",
}

func init() {
	registry.Global().Register(getCmd, postCmd)
}
