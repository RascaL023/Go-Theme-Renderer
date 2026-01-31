package main

import (
	"fmt"
	"gtr/internal/core/context"
	"gtr/internal/core/themes/state"
	"gtr/internal/core/themes/theme"
	"gtr/internal/loader"
	"gtr/internal/register"
	"os"
	"path/filepath"
)

func main() {
	// ================================ SOURCE INPUT ================================
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

	// Load the json theme config
	raws, err := loader.LoadJSON[theme.RawTheme](
		filepath.Join("themes", state.Theme.Name, "theme.json"),
	);
	if err != nil {
		panic("Error on loading json theme");
	}
	// ================================ SOURCE INPUT ================================
	theme.BuildResolvedTheme(&raws.Theme);
	ctx := context.Context{
		Theme:   &raws.Theme,
		Domains: make(map[string]any),
	}


	// ================================ DOMAIN PROCESSOR ================================
	for processorName, raw := range raws.Domain {
		processor, ok := register.GetDomainProcessor(processorName);
		if !ok {
			fmt.Println("Unknown processor", processorName);
			continue;
		}

		parsed, err := processor.Parse(raw);
		if err != nil {
			panic("Error on parsing Domain raw");
		}

		resolved, err := processor.Resolve(parsed, &ctx);
		if err != nil {
			panic("Error on resolving domain");
		}

		ctx.Domains[processorName] = resolved;
	}
	// ================================ DOMAIN PROCESSOR ================================


	// ================================ TOOLS PROCESSOR ================================
	for toolName, toolPath := range tools {
		processor, ok := register.GetToolProcessor(toolName);
		if !ok {
			fmt.Println("Unknown tool", toolName);
			continue;
		}

		var parsed any = nil;
		raw, hasConfig := raws.Tools[toolName];
		if hasConfig {
			parsed, err = processor.Parse(raw);
			if err != nil {
				panic(err);
			}
		}

		resolved, err := processor.Resolve(parsed, &ctx);
		if err != nil {
			panic(err)
		}

		err = processor.Render(toolPath.TemplatePath, toolPath.OutputPath, resolved);
		if err != nil {
			panic(err)
		}
		fmt.Println("Success rendering", toolName);
	}
	// ================================ TOOLS PROCESSOR ================================
}
