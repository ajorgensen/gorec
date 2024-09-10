package gorec

type Request struct {
	Name    string            `yaml:"name"`
	Type    string            `yaml:"type"`
	Method  string            `yaml:"method"`
	URL     string            `yaml:"url"`
	Headers map[string]string `yaml:"headers"`
	Body    string            `yaml:"body"`
}
