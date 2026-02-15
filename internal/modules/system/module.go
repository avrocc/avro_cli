package system

import (
	"avro_cli/internal/app/registry"
	"avro_cli/internal/domain"
)

var category = domain.Category{
	Name:        "system",
	Description: "System information and utilities",
	Icon:        "\U0001F5A5",
}

func init() {
	registry.Global().Register(infoCmd, envCmd, pathCmd, updateCmd)
}
