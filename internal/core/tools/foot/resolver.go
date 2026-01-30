package foot

import "gtr/internal/core/domain/terminal"

func Resolve(term terminal.Terminal) (*Foot, error) {
	return &Foot {
		Foreground: term.Foreground,
		Background: term.Background,

		Regular: term.Regular,
		Bright: term.Bright,
	}, nil;
}
