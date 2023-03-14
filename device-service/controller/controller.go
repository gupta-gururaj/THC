package controller

import (
	"encoding/json"
	"fmt"

	"device-service/model"
	helpers "device-service/utils"
)

type Listener int

func (l *Listener) GetLine(line []byte, reply *model.Reply) error {
	var pLoad model.Payload
	err := json.Unmarshal(line, &pLoad)
	if err != nil {
		fmt.Println("Error in Unmarshaling payload:", err)
	}
	fmt.Println("rv", pLoad)
	helpers.AddDeviceToRoom(pLoad)
	return nil
}
