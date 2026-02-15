package git

import (
	"avro_cli/internal/app/registry"
	"avro_cli/internal/domain"
)

var category = domain.Category{
	Name:        "git",
	Description: "Git operations",
	Icon:        "\U0001F500",
}

func init() {
	registry.Global().Register(cloneCmd, statusCmd, logCmd, branchCmd)
}
