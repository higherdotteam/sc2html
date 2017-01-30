package main

import "github.com/higherdotteam/sc2html/client"
import "os"

func main() {
	if len(os.Args) == 1 {
		client.ListTeams()
		return
	}
	if len(os.Args) == 2 {
		client.ListRooms(os.Args[1])
		return
	}
	if len(os.Args) == 3 {
		client.SaveHTML(os.Args[1], os.Args[2])
		return
	}
}
