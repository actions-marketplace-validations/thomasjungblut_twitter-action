package main

import (
	"flag"
	"github.com/coreos/pkg/flagutil"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"log"
	"os"
)

func checkNonEmptyOrFatal(s string, varName string) {
	if s == "" {
		log.Fatalf("%s was not passed", varName)
	}
}

func main() {
	flags := flag.NewFlagSet("tweet", flag.ExitOnError)
	appKey := flags.String("app-key", "", "Twitter App Key")
	appSecret := flags.String("app-secret", "", "Twitter App Secret")
	accessToken := flags.String("access-token", "", "Twitter Personal Access Token")
	accessSecret := flags.String("access-secret", "", "Twitter Personal Access Secret")
	tweetMessage := flags.String("message", "", "Tweet Message")
	err := flags.Parse(os.Args[1:])
	if err != nil {
		log.Fatalf("error during argument parsing: %v", err)
	}

	err = flagutil.SetFlagsFromEnv(flags, "TWITTER")
	if err != nil {
		log.Fatalf("error when setting flags from env variables: %v", err)
	}

	checkNonEmptyOrFatal(*appKey, "app-key")
	checkNonEmptyOrFatal(*appSecret, "app-secret")
	checkNonEmptyOrFatal(*accessToken, "access-token")
	checkNonEmptyOrFatal(*accessSecret, "access-secret")

	// validate a content is available and does not exceed 280 chars, because those are the rules of twitter
	if len(*tweetMessage) == 0 {
		log.Fatal("Your tweet is empty!")
	} else if len(*tweetMessage) > 280 {
		log.Fatal("Tweet must be less than 280 characters long")
	}

	config := oauth1.NewConfig(*appKey, *appSecret)
	token := oauth1.NewToken(*accessToken, *accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)
	tweet, _, err := client.Statuses.Update(*tweetMessage, nil)

	if err != nil {
		log.Fatalf("error while posting tweet: %v", err)
	}

	log.Printf("Successfully sent tweet: '%s' at %s", *tweetMessage, tweet.CreatedAt)
}
