package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/rpc"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error in loading env variables", err)
	}
	client, err := rpc.Dial("tcp", os.Getenv("ADDRESS")+":"+os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}
	// line := []byte(`hello`)
	room := &Room{
		Type: "open",
		Name: "room1",
		State: State{
			Occupancy: "Moderate",
			Mode:      "mode1",
			IPM:       "ipm-1",
		},
		SetPoint: SetPoint{
			Temperature:      20.7,
			Humidification:   "xyz",
			Dehumidification: "abc",
			Co2:              4.01,
		},
	}
	fmt.Println(room)
	var reply Reply
	var pLoad Payload
	pLoad.Event = "CREATE_ROOM"
	pLoad.Data = *room
	byteData, err := json.Marshal(pLoad)
	if err != nil {
		fmt.Println("Error in Marshaling", err)
	}
	err = client.Call("Listener.GetLine", byteData, &reply)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Reply: %v, Data: %v", reply, reply.Data)
}
