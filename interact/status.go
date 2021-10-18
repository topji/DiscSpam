package interact

import (
	"Raid-Client/utils"
	"fmt"
	"time"
)

var tokens []string
var updating bool

func ChangeStatus(token string) {
	if !updating {
		time.Sleep(100 * time.Millisecond)
		ws := utils.SetupWebSocket(token)
		go utils.RecieveIncomingPayloads(ws, token)
		for {
			if utils.WSConnected {
				utils.SetStatus(utils.Status, ws)
				fmt.Printf("%s %s\n", white(token), green("| Successfully set the status"))
				ws.Close()
				break
			}
		}
	}
}

// Loop presence message every 60s
func loopMessage() {
	for {
		time.Sleep(100 * time.Second)
		updating = true
		for _, tkn := range tokens {
			time.Sleep(100 * time.Millisecond)
			ws := utils.SetupWebSocket(tkn)
			go utils.RecieveIncomingPayloads(ws, tkn)

			for {
				if utils.WSConnected {
					utils.SetStatus(utils.Status, ws)
					ws.Close()
					break
				}
			}
		}
		updating = false
	}
}
