package register

import (
	"encoding/json"
	"gtr/internal/core/context"
)

type Parser interface {
	Parse(json.RawMessage) (any, error)
}

type Resolver interface {
	Resolve(any, *context.Context) (any, error)
}

type Renderer interface {
	Render(templatePath, outputPath string, data any) error
}
