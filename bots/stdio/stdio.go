package main
import (
	"github.com/matrix-org/gomatrix"
	"github.com/tookmund/gomatrixbot"
	"bufio"
	"fmt"
	"os"
)
func main() {
	cli := gomatrixbot.Login("https://matrix.org")
	roomid := gomatrixbot.RoomId()
	user := gomatrixbot.User()
	syncer := cli.Syncer.(*gomatrix.DefaultSyncer)
	syncer.OnEventType("m.room.message", func(ev *gomatrix.Event) {
		body, ok := ev.Body()
		if ok && ev.Sender != "@"+user+":matrix.org" {
			fmt.Println(ev.Sender, ": ", body)
		}	
	})
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		_, err := cli.SendText(roomid, scan.Text())
		if err != nil {
			fmt.Println("Send Failed: ", err)
		}
	}	
}
