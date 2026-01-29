package terminal

func Resolve(raw TerminalInput) (*Terminal, error) {
    return &Terminal{
        Foreground: raw.Foreground,
        Background: raw.Background,
				SelectionFg: raw.Selection.Foreground,
				SelectionBg: raw.Selection.Background,
        Cursor:     raw.Cursor,
        Opacity:    raw.Opacity,
        Regular:    raw.Palette.Normal,
        Bright:     raw.Palette.Bright,
    }, nil;
}

