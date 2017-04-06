package gomatrixbot

import (
	"github.com/matrix-org/gomatrix"
	"fmt"
	"os"
)

type Client struct {
	client *gomatrix.Client
}
type EventCallback gomatrix.OnEventListener

func Login() *Client {
	user, pass, homeserver := getsecret()
	cli, _ := gomatrix.NewClient(homeserver, "", "")
	resp, err := cli.Login(&gomatrix.ReqLogin{
		Type: "m.login.password",
		User: user,
		Password: pass,
	})
	if err != nil {
		panic(err)
	}
	cli.SetCredentials(resp.UserID, resp.AccessToken)
	client := &Client{cli}
	go sync(client)
	return client
}

func sync(cli *Client) {
	for {
		if err := cli.client.Sync(); err != nil {
			fmt.Println("Sync() returned ", err)
		}
	}
}

func getsecret() (user, pass, homeserver string) {
	file, err := os.Open("login")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Fscan(file, &user, &pass, &homeserver)
	return 
}

func (cli Client) HandleEvent(event string, callback EventCallback) {
	syncer := cli.client.Syncer.(*gomatrix.DefaultSyncer)
	syncer.OnEventType(event, gomatrix.OnEventListener(callback))
}

func (cli Client) SendText(roomid, text string) error {
	_, err := cli.client.SendText(roomid, text)
	return err
}

func (cli Client) UserID() string {
	return cli.client.UserID
}
