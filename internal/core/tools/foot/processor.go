package foot

import (
	"encoding/json"
	"fmt"
	"gtr/internal/core/context"
	"gtr/internal/core/domain/terminal"
	"gtr/internal/register"
	"gtr/internal/renderer"
)

type FootProcessor struct{}

func (FootProcessor) Name() string { return "foot"; }

func (FootProcessor) Parse(_ json.RawMessage) (any, error) {
	return nil, nil;
}

func (FootProcessor) Resolve(in any, ctx *context.Context) (any, error) {
	terminalCfg, ok := ctx.Domains["terminal"].(*terminal.Terminal)
	if !ok {
		return nil, fmt.Errorf("terminal domain not resolved");
	}
	return Resolve(*terminalCfg);
}

func (FootProcessor) Render(templatePath, outputPath string, data any) error {
	return renderer.Render(templatePath, outputPath, data);
}

func init() { register.RegisterToolProcessor(FootProcessor{}); }

