package http

import (
	"avro_cli/internal/domain"
	"context"
	"fmt"
)

var getCmd = domain.CommandDescriptor{
	Category:    category,
	Name:        "get",
	Description: "Perform an HTTP GET request",
	Args: []domain.ArgDef{
		{Name: "url", Description: "Request URL", Required: true},
	},
	Flags: []domain.ArgDef{
		{Name: "header", Short: "H", Description: "Header in key:value format"},
	},
	Action: func(ctx domain.CommandContext) domain.Result[string] {
		url := ctx.Args["url"]
		if url == "" {
			return domain.Fail[string](&domain.ValidationError{Field: "url", Message: "URL is required"})
		}

		headers := parseHeaders(ctx.Flags["header"])
		status, body, err := ctx.HTTP.Get(context.Background(), url, headers)
		if err != nil {
			return domain.Fail[string](err)
		}

		return domain.Ok(fmt.Sprintf("HTTP %d\n\n%s", status, body))
	},
}

var postCmd = domain.CommandDescriptor{
	Category:    category,
	Name:        "post",
	Description: "Perform an HTTP POST request",
	Args: []domain.ArgDef{
		{Name: "url", Description: "Request URL", Required: true},
		{Name: "body", Description: "Request body", Required: false},
	},
	Flags: []domain.ArgDef{
		{Name: "header", Short: "H", Description: "Header in key:value format"},
	},
	Action: func(ctx domain.CommandContext) domain.Result[string] {
		url := ctx.Args["url"]
		if url == "" {
			return domain.Fail[string](&domain.ValidationError{Field: "url", Message: "URL is required"})
		}

		body := ctx.Args["body"]
		headers := parseHeaders(ctx.Flags["header"])
		status, respBody, err := ctx.HTTP.Post(context.Background(), url, body, headers)
		if err != nil {
			return domain.Fail[string](err)
		}

		return domain.Ok(fmt.Sprintf("HTTP %d\n\n%s", status, respBody))
	},
}

func parseHeaders(raw string) map[string]string {
	headers := make(map[string]string)
	if raw == "" {
		return headers
	}
	// Simple k:v parsing for a single header
	for i := 0; i < len(raw); i++ {
		if raw[i] == ':' {
			key := raw[:i]
			val := raw[i+1:]
			if len(val) > 0 && val[0] == ' ' {
				val = val[1:]
			}
			headers[key] = val
			break
		}
	}
	return headers
}
