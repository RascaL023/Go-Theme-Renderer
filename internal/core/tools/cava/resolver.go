package cava

func Resolve(parsed Raw) (*Cava, error) {
	return &Cava {
		Gradient: parsed.GradientColor,
	}, nil;
}

