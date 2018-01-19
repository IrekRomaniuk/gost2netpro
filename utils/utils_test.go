package utils

import (
	"flag"
	"fmt"
	"log"
	"testing"
)

var (
	user = flag.String("u", "", "user")
	pass = flag.String("p", "", "pass")
	api  = flag.String("api", "", "api")
	cmd  = flag.String("cmd", "uptime", "cmd")
)

//TestReadFile  : go test -run TestReadFile
func TestReadFile(t *testing.T) {
	list := &Netpro{Hosts: ""}
	//fmt.Printf("List name: %v\n", list.Name)
	err := list.ReadFile("../targets/pinglist.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("netpro: %v\n", *list)
}

//TestPostPage : go test -run TestPostPage -args -u=user -p=pass
func TestPostPage(t *testing.T) {
	list := Netpro{Hosts: "1.1.1.1, 2.2.2.2", User: *user, Pass: *pass, Cmd: *cmd}
	fmt.Printf("List: %v\n", list)
	ID, err := PostPage("https://ubuntu-st2/api/v1/webhooks/netpro", *api, list)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID: %v\n", ID)
}
