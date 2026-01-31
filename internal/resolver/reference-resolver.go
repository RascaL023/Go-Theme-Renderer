package resolver

import "gtr/internal/core/context"

func ResolveVar(s string, ctx *context.Context) string {
	if len(s) < 4 || s[0] != '$' {
		return s
	}

	switch s[1] {
	case 't': // $th.
		if s[2] == 'h' && s[3] == '.' {
			if val, ok := ctx.Theme.Flat[s[4:]]; ok {
				return val
			}
		}
	}

	return s
}

func ResolveVars(arr []string, ctx *context.Context) []string {
    if len(arr) == 0 {
        return arr
    }

    out := make([]string, len(arr))
    for i, v := range arr {
        out[i] = ResolveVar(v, ctx)
    }
    return out
}

