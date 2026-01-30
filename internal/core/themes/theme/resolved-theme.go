package theme

type ThemeType string

const (
	Dark  ThemeType = "dark"
	Light ThemeType = "light"
)

type ResolvedTheme struct {
	Type    ThemeType
	Palette Palette
}



type Palette struct {
	Base        BasePalette
	Surface     SurfacePalette
	Overlay     OverlayPalette
	Accent      AccentPalette
	Status      StatusPalette
	Text        TextPalette
	Border      BorderPalette
	Interaction InteractionPalette
}

type BasePalette struct {
	Bg string
	Fg string
}

type SurfacePalette struct {
	Base   string // surface.0
	Raised string // surface.1
	Hover  string // surface.2
	Active string // surface.3
}

type OverlayPalette struct {
	Subtle  string
	Default string
	Strong  string
}

type AccentPalette struct {
	Primary   string
	Secondary string
	Tertiary  string
}

type StatusPalette struct {
	Success string
	Warning string
	Error   string
	Info    string
}

type TextPalette struct {
	Primary   string
	Secondary string
	Teritary string
	Muted     string
	Disabled  string
}

type BorderPalette struct {
	Default string
	Focus   string
	Urgent  string
}

type InteractionPalette struct {
	HoverBg string
	HoverFg string
	Outline string
}

