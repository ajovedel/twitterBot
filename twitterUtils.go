package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/ChimeraCoder/anaconda"
)

func authenticate() *anaconda.TwitterApi {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)
	return api
}

func postTweetRandomTime(api *anaconda.TwitterApi) {
	/*
			rand.Seed(time.Now().UTC().UnixNano())
			randInt := rand.Int()

		// wait random time and post
		fmt.Printf("Will sleep for %d second and then post\n", randInt%10)
		time.Sleep(time.Hour * time.Duration(randInt%23))
	*/
	api.PostTweet("yay marriott!! #MembersGetIt #RewardsPoints #FindYourRoute #MRpoints #MarriottRewards", nil)

	fmt.Print("Just posted at: ")
	fmt.Println(time.Now().Format(time.RFC3339))
}

func retweetRandomTime(api *anaconda.TwitterApi) {
	timeMinus24H := fmt.Sprintf("%s", time.Now().AddDate(0, 0, -1).Format("2006-01-02"))
	tweetSearch := fmt.Sprintf("from:marriottrewards since:%s", timeMinus24H)
	tweets, _ := api.GetSearch(tweetSearch, nil)

	for _, tweet := range tweets.Statuses {
		tweetText := strings.ToLower(tweet.FullText)
		if strings.Contains(tweet.FullText, "RT") || strings.Contains(tweetText, "#membersgetit") ||
			strings.Contains(tweetText, "#rewardspoints") || strings.Contains(tweetText, "#findyourroute") || strings.Contains(tweetText, "#mrpoints") {
			fmt.Printf("Retweeting: %s\n", tweet.FullText)
			api.Retweet(tweet.Id, false)
		}
	}
}
