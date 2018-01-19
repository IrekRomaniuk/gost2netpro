package utils

import (
	"flag"
	"fmt"
	"log"
	"testing"
)

var (
	user = flag.String("u", "user", "user")
	pass = flag.String("p", "pass", "pass")
	api  = flag.String("api", "", "api")
	cmd  = flag.String("cmd", "uptime", "cmd")
)

//TestReadFile  : go test -run TestReadFile
func TestReadFile(t *testing.T) {
	list := &Netpro{Hosts: "", User: *user, Pass: *pass, Cmd: *cmd}
	//fmt.Printf("List name: %v\n", list.Name)
	err := list.ReadFile("../targets/pinglist.txt", *user, *pass, *cmd)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("netpro: %v\n", *list)
}

//TestPostPage : go test -run TestPostPage -args -u=user -p=pass -api=''
func TestPostPage(t *testing.T) {
	list := Netpro{Hosts: "1.1.1.1, 2.2.2.2", User: *user, Pass: *pass, Cmd: *cmd}
	fmt.Printf("List: %v\n", list)
	cmd, err := PostPage("https://ubuntu-st2/api/v1/webhooks/netpro", *api, list)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("cmd: %v\n", cmd)
}
