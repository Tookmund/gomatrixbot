package main
import (
	"github.com/matrix-org/gomatrix"
	"github.com/tookmund/gomatrixbot"
	"fmt"
)

func main() {
	cli := gomatrixbot.Login("https://matrix.org")
	user := gomatrixbot.User()
	syncer := cli.Syncer.(*gomatrix.DefaultSyncer)
	syncer.OnEventType("m.room.message", func(ev *gomatrix.Event) {
		body, ok := ev.Body()
		if ok && ev.Sender != user {
			fmt.Println(ev.Sender, ": ", body)
		}	
	})
	// Do other stuff with cli here...
	for {
	}
}