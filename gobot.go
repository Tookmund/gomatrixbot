package gomatrixbot

import (
	"github.com/matrix-org/gomatrix"
	"fmt"
	"os"
	"strings"
	"bytes"
	"io/ioutil"
)
var roomid string
var username string

func Login(homeserver string) *gomatrix.Client {
	user, pass, room := getsecret()
	roomid = room
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

func getsecret() (user, pass, room string) {
	roombuf, err := ioutil.ReadFile("roomid")
	if err != nil {
		fmt.Print("Cannot find roomid file!\n")
		os.Exit(0)
	}
	n := bytes.IndexByte(roombuf,'\n')
	room = string(roombuf[:n])
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

func RoomId() string {
	return roomid
}

func User() string {
	return username
}
