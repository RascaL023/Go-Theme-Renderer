package register

var RegisteredDomain = map[string]DomainProcessor{}

type DomainProcessor interface {
    Name() string
		Parser
		Resolver
}

func RegisterDomainProcessor(t DomainProcessor) {
    RegisteredDomain[t.Name()] = t;
}

func GetDomainProcessor(name string) (DomainProcessor, bool) {
    p, ok := RegisteredDomain[name];
    return p, ok;
}



