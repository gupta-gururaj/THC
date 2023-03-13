package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/rpc"
	"os"

	"github.com/gin-gonic/gin"

	"provisional-service/model"
)

func ABC(c *gin.Context) {
	var req model.Request
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("ERROR Reading data from request: ", err)
	}
	fmt.Println("ctx", string(jsonData))
	err = json.Unmarshal(jsonData, &req)
	if err != nil {
		fmt.Println("ERROR Unmarshaling data from request: ", err)
	}
	client, err := rpc.Dial("tcp", os.Getenv("ADDRESS")+":"+os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}
	// line := []byte(`hello`)
	room := &model.Room{
		Type: req.Type,
		Name: "room1",
		State: model.State{
			Occupancy: "Moderate",
			Mode:      "mode1",
			IPM:       "ipm-1",
		},
		SetPoint: model.SetPoint{
			Temperature:      20.7,
			Humidification:   "xyz",
			Dehumidification: "abc",
			Co2:              4.01,
		},
	}
	fmt.Println(room)
	var reply model.Reply
	var pLoad model.Payload
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
	c.Data(http.StatusOK, gin.MIMEJSON, jsonData)
}
