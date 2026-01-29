package terminal

import (
	"encoding/json"
	"gtr/internal/core/context"
	"gtr/internal/register"
)

type TerminalProcessor struct{}

func (TerminalProcessor) Name() string { return "terminal"; }

func (TerminalProcessor) Parse(raw json.RawMessage) (any, error) {
	var cfg TerminalInput;
	err := json.Unmarshal(raw, &cfg);
	return cfg, err;
}

func (TerminalProcessor) Resolve(in any, ctx context.Context) (any, error) {
	return Resolve(in.(TerminalInput));
}

func init() { register.RegisterDomainProcessor(TerminalProcessor{}); }
