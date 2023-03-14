package controller

import (
	"encoding/json"
	"fmt"

	"configuartion-service/model"
	"configuartion-service/util"
)

type Listener int

func (l *Listener) GetLine(line []byte, reply *model.Reply) error {
	rv := string(line)
	var payLoad model.Payload
	err := json.Unmarshal(line, &payLoad)
	if err != nil {
		fmt.Println("Error in Unmarshaling the payload", err)
	}
	if payLoad.Event == "CREATE_ROOM" {
		util.CreateRoom(payLoad)
		*reply = model.Reply{rv}
	}
	return nil
}
