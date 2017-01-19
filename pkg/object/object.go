package object



type Container struct {
	Image string `json:"-"`  // this is contained as a key in the parent map

}

//type Service struct {
//	Name string `json:"-"`  // this is contained as a key in the parent map
//	Containers map[string]*Container
//}
type Service map[string]*Container

type Volumes struct {

}

type OpenCompose struct {
	Version string `json:"version,omitempty"`
	Services map[string]*Service `json:"services"`
	Volumes Volumes `json:"volumes,omitempty"`
}