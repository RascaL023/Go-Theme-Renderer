package terminal

import (
	"encoding/json"
	"gtr/internal/core/context"
	"gtr/internal/register"
	"strings"
)

type TerminalProcessor struct{}

func (TerminalProcessor) Name() string { return "terminal"; }

func (TerminalProcessor) Parse(raw json.RawMessage) (any, error) {
	var cfg TerminalInput;
	err := json.Unmarshal(raw, &cfg);
	return cfg, err;
}

func (TerminalProcessor) Resolve(in any, ctx *context.Context) (any, error) {
	inp := in.(TerminalInput);
	if strings.HasPrefix(inp.Foreground, "$theme.") {
		switch strings.TrimPrefix(inp.Foreground, "$theme.") {
		case "palette.text.teritary":
			inp.Foreground = ctx.Theme.Palette.Text.Teritary;
		}
	}

	return Resolve(inp);
}

func init() { register.RegisterDomainProcessor(TerminalProcessor{}); }
