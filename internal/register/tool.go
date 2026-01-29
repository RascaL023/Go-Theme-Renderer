package register

var RegisteredTool = map[string]ToolProcessor{}

type ToolProcessor interface {
    Name() string
		Parser
		Resolver
		Renderer
}

func RegisterToolProcessor(t ToolProcessor) {
    RegisteredTool[t.Name()] = t;
}

func GetToolProcessorProcessor(name string) (ToolProcessor, bool) {
    p, ok := RegisteredTool[name];
    return p, ok;
}


