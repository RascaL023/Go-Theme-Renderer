package context

import "gtr/internal/core/themes/theme"


type Context struct {
	Theme	   *theme.ResolvedTheme
	Domains  map[string]any
}

