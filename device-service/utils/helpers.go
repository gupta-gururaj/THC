package helpers

import (
	"fmt"

	"device-service/model"
)

func AddDeviceToRoom(data model.Payload) {
	var device model.Dehum
	device.RoomUUID = data.RoomData.UUID
	fmt.Println("device", device)
}
