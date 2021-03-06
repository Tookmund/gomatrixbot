package main
import (
	"github.com/matrix-org/gomatrix"
	"github.com/tookmund/gomatrixbot"
	"bufio"
	"fmt"
	"os"
)
func main() {
	cli := gomatrixbot.Login()
	roomid := gomatrixbot.Roomid()
	cli.HandleEvent("m.room.message", func(ev *gomatrix.Event) {
		body, ok := ev.Body()
		if ok && ev.Sender != cli.UserID() {
			fmt.Println(ev.Sender, ": ", body)
		}	
	})
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		err := cli.SendText(roomid, scan.Text())
		if err != nil {
			fmt.Println("Send Failed: ", err)
		}
	}	
}
