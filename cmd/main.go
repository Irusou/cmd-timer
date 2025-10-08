package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gen2brain/beeep"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: time <duration> <title> <message>")
		os.Exit(1)
	}

	durationStr := os.Args[1]

	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		fmt.Printf("Invalid duration: %v\n", err)
		os.Exit(1)
	}

	time.Sleep(duration)

	title := ""
	msg := ""
	
	if len(os.Args) == 3 {
		title = os.Args[2]
	} 
	
	if len(os.Args) == 4 {
		title = os.Args[2]
		msg = os.Args[3]
	} 

	err = beeep.Alert(title, msg, "")

	if err != nil {
		fmt.Println("Could not show notification:", err)
	} else {
		fmt.Println("🔔 Timer finished! Notification sent.")
	}
}

