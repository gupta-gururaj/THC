package main

import (
	"configuartion-service/model"
	"configuartion-service/util"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"

	"github.com/joho/godotenv"
)

type Listener int

func (l *Listener) GetLine(line []byte, reply *model.Reply) error {
	rv := string(line)
	var payLoad model.Payload
	err := json.Unmarshal(line, &payLoad)
	if err != nil {
		fmt.Println("Error in Unmarshaling the payload", err)
	}

	f, _ := payLoad.Data.(map[string]interface{})
	if payLoad.Event == "CREATE_ROOM" {
		util.CreateRoom(f)
		*reply = model.Reply{rv}
	}
	return nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error in loading env variables", err)
	}
	addy, err := net.ResolveTCPAddr("tcp", os.Getenv("ADDRESS")+":"+os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}
	inbound, err := net.ListenTCP("tcp", addy)
	if err != nil {
		log.Fatal(err)
	}
	listener := new(Listener)
	rpc.Register(listener)
	rpc.Accept(inbound)
}
