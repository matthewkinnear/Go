package main

import (
	"fmt"
	"time"

	fetcher "my-pubsub-app/internal/fetcher"
	pubsub "my-pubsub-app/internal/pub_sub"
	webserver "my-pubsub-app/internal/web_server"
)

func main() {
	// Subscribe to user data
	userChan := pubsub.Subscribe()
	// Assign the userChan to the global variable so it can be accessed in the webserver
	webserver.SetUserChan(userChan)

	// Start the web server
	go func() {
		fmt.Println("Starting web server on http://localhost:8080")
		webserver.StartWebServer()
	}()

	// Continuously fetch user data and publish it every 10 seconds
	for {
		user, err := fetcher.FetchUser()
		if err != nil {
			fmt.Println("Error fetching user data:", err)
			continue
		}
		fmt.Printf("Fetched user: %+v\n", user)
		pubsub.Publish(user)
		time.Sleep(10 * time.Second) // Add a delay to fetch a new user every 10 seconds
	}
}
