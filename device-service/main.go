package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"

	"github.com/joho/godotenv"

	"device-service/controller"
)

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
	listener := new(controller.Listener)
	rpc.Register(listener)
	rpc.Accept(inbound)
}
