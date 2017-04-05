package main
import (
	"github.com/matrix-org/gomatrix"
	"bufio"
	"fmt"
	"os"
)
func main() {
	user, pass, roomid := getsecret()
	cli := login("https://matrix.org", user, pass)
	syncer := cli.Syncer.(*gomatrix.DefaultSyncer)
	syncer.OnEventType("m.room.message", func(ev *gomatrix.Event) {
		body, ok := ev.Body()
		if ok && ev.Sender != "@"+user+":matrix.org" {
			fmt.Println(ev.Sender, ": ", body)
		}	
	})
 	go sync(cli)
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		_, err := cli.SendText(roomid, scan.Text())
		if err != nil {
			fmt.Println("Send Failed: ", err)
		}
	}	
}
