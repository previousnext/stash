package main

import (
	"flag"
	"fmt"

	"github.com/ttacon/pretty"
	"github.com/ttacon/stash"
)

var (
	username = flag.String("u", "", "username")
	password = flag.String("p", "", "password")
	host     = flag.String("h", "http://localhost:7990", "host to query")

	group = flag.String("g", "", "group to get")
)

func main() {
	flag.Parse()
	c := stash.NewClient(*username, *password, *host)
	group, err := c.GroupService().DeleteGroup(*group)
	if err != nil {
		fmt.Println("failed to retrieve user: ", err)
		return
	}

	pretty.Println(group)
}
