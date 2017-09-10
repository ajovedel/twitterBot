package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello World!")

	api := authenticate()
	api.PostTweet("Hello world again!!!", nil)
	retweetRandomTime(api)
	for {
		// post #marriott tweet at random time for 50 points
		go postTweetRandomTime(api)

		// retweet certain marriott tweets every 24 hours
		go retweetRandomTime(api)

		// wait 1 day
		time.Sleep(24 * time.Hour)
		fmt.Print("Finished waiting 24 hours. looping again at: ")
		fmt.Println(time.Now().Format(time.RFC3339))
	}
}
