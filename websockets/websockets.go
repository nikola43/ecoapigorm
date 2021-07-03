package websockets

import (
	"encoding/json"
	"fmt"
	"github.com/antoniodipinto/ikisocket"
	"github.com/nikola43/ecoapigorm/models"
	"strconv"
)

var SocketInstance *ikisocket.Websocket
var SocketClients map[string]string

func Emit(socketEvent models.SocketEvent, id uint) {
	event, err := json.Marshal(socketEvent)
	if err != nil {
		fmt.Println(err)
	}

	var socketClientId = strconv.FormatUint(uint64(id), 10)
	emitSocketErr := SocketInstance.EmitTo(SocketClients[socketClientId], event)
	if emitSocketErr != nil {
		fmt.Println(emitSocketErr)
	}
}
