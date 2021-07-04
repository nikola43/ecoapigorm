package websockets

import (
	"encoding/json"
	"fmt"
	"github.com/antoniodipinto/ikisocket"
	"strconv"
)

type SocketEvent struct {
	Type   string      `json:"type"`
	Action string      `json:"action"`
	Data   interface{} `json:"data"`
}


var SocketInstance *ikisocket.Websocket
var SocketClients map[string]string

func  Emit(socketEvent SocketEvent, id uint) {
	event, err := json.Marshal(socketEvent)
	if err != nil {
		fmt.Println(err)
	}

	var socketClientId = strconv.FormatUint(uint64(id), 10)
	if uuid, found := SocketClients[socketClientId]; found {
		emitSocketErr := SocketInstance.EmitTo(uuid, event)
		if emitSocketErr != nil {
			fmt.Println(emitSocketErr)
		}
	}
}
