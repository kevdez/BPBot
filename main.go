package main

import (
	"flag"
	"fmt"
	"strconv"
	// "io/ioutil"

	"./hipchat"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

var (
	token  = flag.String("token", "", "The HipChat AuthToken")
	roomId = flag.String("room", "", "The HipChat room id")
	message = flag.String("message", "", "The message to send")
)

func main() {


	flag.Parse()
	if *token == "" || *roomId == "" || *message == "" {
		flag.PrintDefaults()
		return
	}
	c := hipchat.NewClient(*token)

	rooms,_,err := c.Room.List()
	for _, room:= range rooms.Items{
		fmt.Println(room.Name + " " + strconv.Itoa(room.ID))
	}

	notifRq := &hipchat.NotificationRequest{Message: *message, Notify: true, Color: "red" }
	resp, err := c.Room.Notification(*roomId, notifRq)
	if err != nil {
		fmt.Printf("Error during room notification %q\n", err)
		fmt.Printf("Server returns %+v\n", resp)
		return
	}
	fmt.Println("Lol sent!")
}
