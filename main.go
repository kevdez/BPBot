package main

import (
	"flag"
	"fmt"
	// "io/ioutil"

	"hipchat"
)

var (
	message = flag.String("message", "", "The message to send")
	color = flag.String("color", "", "The color of the message")
	room = flag.String("room", "", "The room API ID")
)

func main() {


	flag.Parse()
	if *message == "" || *room == ""{
		flag.PrintDefaults()
		return
	}
	c := hipchat.NewClient("Fzdyz46mmk6tLJRgo19VjGpKKmclqr0YjXKX1989")

	if *color != "" {
		notifRq := &hipchat.NotificationRequest{Message: *message, Notify: true, Color: *color}
		resp, err := c.Room.Notification(*room, notifRq)
		if err != nil {
			fmt.Printf("Error during room notification %q\n", err)
			fmt.Printf("Server returns %+v\n", resp)
			return
		}
		fmt.Println("Sent!")
	} else {
		notifRq := &hipchat.NotificationRequest{Message: *message, Notify: true}
		resp, err := c.Room.Notification(*room, notifRq)
		if err != nil {
			fmt.Printf("Error during room notification %q\n", err)
			fmt.Printf("Server returns %+v\n", resp)
			return
		}
	fmt.Println("Sent!")
	}
}
