package util

import (
	"encoding/json"
	"fmt"
	"log"
	"net/rpc"
	"os"

	uuid "github.com/satori/go.uuid"

	"configuartion-service/model"
)

type Reply struct {
	Data string
}

func CreateRoom(data model.Payload) {
	// Entry in DB
	newUuid := uuid.NewV4()
	var reply Reply
	fmt.Println("newUuid", newUuid)
	fmt.Println("In Create Room")
	byteId, err := json.Marshal(newUuid)
	if err != nil {
		fmt.Println("Errror in Marshaling uuid:", err)
	}
	data.RoomData.UUID = string(byteId)
	fmt.Println(data)

	client, err := rpc.Dial("tcp", os.Getenv("ADDRESS")+":"+os.Getenv("DEVICE_SERVICE_PORT"))
	if err != nil {
		log.Fatal(err)
	}

	byteData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error in Marshaling data", err)
	}

	err = client.Call("Listener.GetLine", byteData, &reply)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Reply: %v, Data: %v", reply, reply.Data)
}
