package object

type Port struct {
	Address       string
	ContainerPort int
	HostPort      int
	ServicePort   int
	Protocol      string
}

type Mapping struct {
	Port Port
	Type string
	Name string
}

type EnvVariable struct {
	Key   string
	Value string
}

type Container struct {
	Name        string
	Image       string
	Environment []EnvVariable
	Mappings    []Mapping
}

type Service struct {
	Name       string
	Containers []Container
}

type Volume struct {
	// TODO: remove tags when we have Go 1.8
	Name string `json:"name"`
	Size string `json:"size,omitempty"`
	Mode string `json:"mode,omitempty"`
}

type OpenCompose struct {
	Version  string
	Services []Service
	Volumes  []Volume
}
