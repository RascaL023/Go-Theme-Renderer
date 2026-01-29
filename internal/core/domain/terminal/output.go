package terminal

type Terminal struct {
    Foreground string
    Background string

		SelectionFg string
		SelectionBg string

    Cursor     string
    Opacity    float64

    Regular []string
    Bright  []string
}

