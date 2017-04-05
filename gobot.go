package main

import (
	"github.com/matrix-org/gomatrix"
	"fmt"
	"os"
	"strings"
	"bytes"
	"io/ioutil"
)

func login(homeserver, username, password string) *gomatrix.Client {
	cli, _ := gomatrix.NewClient(homeserver, "", "")
	resp, err := cli.Login(&gomatrix.ReqLogin{
		Type: "m.login.password",
		User: username,
		Password: password,
	})
	if err != nil {
		panic(err)
	}
	cli.SetCredentials(resp.UserID, resp.AccessToken)
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
