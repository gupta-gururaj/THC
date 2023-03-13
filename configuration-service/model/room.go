package model

type Reply struct {
	Data string
}

type Payload struct {
	Event string
	Data  interface{}
}

type Room struct {
	Id       string
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
