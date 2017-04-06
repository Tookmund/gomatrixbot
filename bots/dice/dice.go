package main
import (
	"github.com/matrix-org/gomatrix"
	"github.com/tookmund/gomatrixbot"
	"github.com/tookmund/dice"
	"fmt"
	"os"
	"strconv"
)
func main() {
	cli := gomatrixbot.Login()
	roomid := roomid()
	cli.HandleEvent("m.room.message", func(ev *gomatrix.Event) {
		body, ok := ev.Body()
		if ok {
			fmt.Println(ev.Sender, ": ", body)
			if ev.Sender != cli.UserID() {
				var sides, number int
				_, err := fmt.Sscanf(body,"%dd%d",&number, &sides)
				if err == nil {
					cli.SendText(roomid, body+": "+strconv.Itoa(dice.Roll(sides,number)))
				}	
			}	
		}	
	})
	for {
	}
}

func roomid() (room string) {
	file, err := os.Open("roomid")
	if err != nil {
		fmt.Println("Cannot open roomid!")
		os.Exit(0)
	}
	_, err = fmt.Fscan(file,&room)
	if err != nil {
		panic(err)
	}
	return
}
