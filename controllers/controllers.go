package controllers

import (
	"Robot/packages/config"
	"fmt"
	"log"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

// type ServerConnection interface {
// 	Connection()
// }

// type server struct {
// 	configs []string
// 	token   []string
// }

func Connection() {
	config.Load()
	configs := oauth1.NewConfig(os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET_KEY"))
	token := oauth1.NewToken(os.Getenv("ACESS_TOKEN"), os.Getenv("ACESS_SECRET_TOKEN"))
	httpClient := configs.Client(oauth1.NoContext, token)

	// Server Client
	client := twitter.NewClient(httpClient)

	// Send Twitter
	sendTweets, _, err := client.Statuses.Update("Testing my twitter robot", nil)
	if err != nil {
		fmt.Println("Twitter upload error", err)
	}
	log.Print(sendTweets.Text)

	// search Tweets
	searchTweets, _, err := client.Search.Tweets(&twitter.SearchTweetParams{
		Query: "#golang",
		Count: 5,
	})
	if err != nil {
		fmt.Println("Search Tweets error", err)
	}

	for _, value := range searchTweets.Statuses {
		log.Print("User name: ", value.User.Name)
		log.Print("Tweet: ", value.Text)

		// retweet from twitter found
		_, _, err := client.Statuses.Retweet(value.ID, nil)
		if err != nil {
			fmt.Println("Retweet error", err)
		}
	}

	// User show
	showUsers, _, err := client.Users.Show(&twitter.UserShowParams{
		ScreenName: "Golang",
	})
	if err != nil {
		fmt.Println(err, "Show User error")
	}
	log.Println(showUsers)

	// list followers

	follower, _, err := client.Followers.List(&twitter.FollowerListParams{})
	if err != nil {
		fmt.Println("Follower error", err)
	}
	fmt.Println("My followers : ", follower)

}
