package core

// TODO marshall method
type Events struct {
	Events []Ewent7W
}

// TODO use other struct; 7W is a way to present info not collect
type Ewent7W struct {
	Who       string
	DidWhat   string
	When      string
	Where     string
	Wherefrom string
	WhereTo   string
	What      string
}

type Agent interface {
	// какая архитектура у агента?
	// методы по типам собираемых событий
	// события мб должны отправляться в канал, из которого читает Core
	GetEvents() <-chan Events
}
