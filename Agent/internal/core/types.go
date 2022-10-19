package core

type Event struct {
	FileName string `json:"filename"`
	Event    string `json:"event"`
	Ts       int64  `json:"ts"`
}

type HostData struct {
	HostName string `json:"hostname"`
	Os       string `json:"os"`
	// TODO ETC
}

type EventsChunk struct {
	HostData HostData `json:"host"`
	Event    []Event  `json:"events"`
}
