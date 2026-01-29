package main

import (
	"fmt"
	"gtr/internal/core/themes/state"
	"gtr/internal/core/themes/theme"
	"gtr/internal/loader"
	"os"
	"path/filepath"

	_ "gtr/internal/core/domain/terminal"
)

func main() {
	// ================================ FIX ================================
	// Get the state of theme
	assetsMap := os.ExpandEnv("$MYENV/map");
	state, err := loader.LoadJSON[state.State](assetsMap + "/.state.json");
	if err != nil {
		panic("Error on parsing state!");
	}
	fmt.Println("[Go] Using state", state.Version);


	// Get the input - output path of each tools
	tools, err := loader.LoadToolMap(assetsMap + "/path.txt", state);
	if err != nil {
		panic("Error on parsing map path!");
	}
	// ================================ FIX ================================
	
	fmt.Println(tools);

	raws, err := loader.LoadJSON[theme.RawTheme](
		filepath.Join("themes", state.Theme.Name, "theme.json"),
	);
	if err != nil {
		panic("Error on loading theme config");
	}

	resolvedTheme, err := theme.Parse(raws.Theme);
	if err != nil {
		panic("Error on parsing theme");
	}
	fmt.Println(resolvedTheme);

}
