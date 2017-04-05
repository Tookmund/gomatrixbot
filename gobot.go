package gomatrixbot

import (
	"github.com/matrix-org/gomatrix"
	"fmt"
	"os"
)

type Client *gomatrix.Client

func Login() Client {
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
	go sync(cli)
	return cli
}
func sync(cli Client) {
	cligm := gomatrix.Client(*cli)
	for {
		if err := cligm.Sync(); err != nil {
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
