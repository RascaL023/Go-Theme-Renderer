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

func GetToolProcessor(name string) (ToolProcessor, bool) {
    p, ok := RegisteredTool[name];
    return p, ok;
}


