package cava

import (
	"encoding/json"
	"gtr/internal/core/context"
	"gtr/internal/register"
	"gtr/internal/renderer"
)

type CavaProcessor struct{}

func (CavaProcessor) Name() string { return "cava"; }

func (CavaProcessor) Parse(raw json.RawMessage) (any, error) {
	var cfg Raw;
	err := json.Unmarshal(raw, &cfg);
	return cfg, err;

}

func (CavaProcessor) Resolve(in any, ctx *context.Context) (any, error) {
	return Resolve(in.(Raw));
}

func (CavaProcessor) Render(templatePath, outputPath string, data any) error {
	return renderer.Render(templatePath, outputPath, data);
}

func init() { register.RegisterToolProcessor(CavaProcessor{}); }


