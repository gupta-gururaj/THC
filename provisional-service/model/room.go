package model

type Reply struct {
	Data string
}

type Request struct {
	Type string `json:"type"`
}

type Room struct {
	UUID     string
	Type     string
	Name     string
	State    State
	SetPoint SetPoint
}

type State struct {
	Occupancy string
	Mode      string
	IPM       string
}
type SetPoint struct {
	Temperature      float64
	Humidification   string
	Dehumidification string
	Co2              float64
}

type Payload struct {
	Event    string
	RoomData Room
}
