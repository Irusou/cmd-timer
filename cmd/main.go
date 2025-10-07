package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gen2brain/beeep"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: mytimer <duration> (e.g. mytimer 10m or mytimer 1h30m)")
		os.Exit(1)
	}

	durationStr := os.Args[1]
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		fmt.Printf("Invalid duration: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("⏰ Timer started for %v...\n", duration)
	time.Sleep(duration)

	msg := fmt.Sprintf("Your %v timer is done!", duration)
	err = beeep.Alert("Timer Finished", msg, "")
	if err != nil {
		fmt.Println("Could not show notification:", err)
	} else {
		fmt.Println("🔔 Timer finished! Notification sent.")
	}
}

