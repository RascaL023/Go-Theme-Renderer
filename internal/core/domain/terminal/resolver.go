package terminal

import (
	"gtr/internal/core/context"
	"gtr/internal/resolver"
)

func Resolve(raw TerminalInput, ctx *context.Context) (*Terminal, error) {
    return &Terminal{
        Foreground: resolver.ResolveVar(raw.Foreground, ctx),
        Background: resolver.ResolveVar(raw.Background, ctx),

				SelectionFg: resolver.ResolveVar(raw.Selection.Foreground, ctx),
				SelectionBg: resolver.ResolveVar(raw.Selection.Background, ctx),

        Cursor:     resolver.ResolveVar(raw.Cursor, ctx),
        Opacity:    raw.Opacity,

        Regular:    resolver.ResolveVars(raw.Palette.Normal, ctx),
        Bright:     resolver.ResolveVars(raw.Palette.Bright, ctx),
    }, nil;
}

