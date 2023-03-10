package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"

	"github.com/joho/godotenv"
)

type Listener int

func (l *Listener) GetLine(line []byte, reply *Reply) error {
	rv := string(line)
	var payLoad Payload
	err := json.Unmarshal(line, &payLoad)
	if err != nil {
		fmt.Println("Error in Unmarshaling the payload", err)
	}

	f, _ := payLoad.Data.(map[string]interface{})
	if payLoad.Event == "CREATE_ROOM" {
		createRoom(f)
		*reply = Reply{rv}
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
