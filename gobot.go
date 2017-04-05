package gomatrixbot

import (
	"github.com/matrix-org/gomatrix"
	"fmt"
	"os"
	"strings"
	"io/ioutil"
)
var username string

func Login(homeserver string) *gomatrix.Client {
	user, pass := getsecret()
	username = user
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
func sync(cli *gomatrix.Client) {
	for {
		if err := cli.Sync(); err != nil {
			fmt.Println("Sync() returned ", err)
		}
	}
}

func getsecret() (user, pass string) {
	userbuf, err := ioutil.ReadFile("password")
	if err != nil {
		fmt.Print("Cannot find password file!\n")
		os.Exit(0)
	}
	fulluser := string(userbuf[:])
	userarray := strings.Split(fulluser, "\n")
	user = userarray[0]
	pass = userarray[1]
	return 
}

func User() string {
	return username
}
