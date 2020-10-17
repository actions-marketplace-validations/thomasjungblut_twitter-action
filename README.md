# Attribution

This is a very simple fork of [xorilog/twitter-action](https://github.com/xorilog/twitter-action). It can only tweet a message from the commandline on-behalf of a user and ***that's it***.

# Getting the tokens

You need a [Twitter developer account](https://developer.twitter.com/) and [create an app](https://apps.twitter.com/).

Once you have created a new app, you need to add permissions to `Read and Write` tweets and get four tokens, namely:
* API Key + Secret of the app
* Access Token + Secret for your own user account 

# Use Action

```
action "Send a tweet" {
  uses = "thomasjungblut/twitter-action@master"
  args = ["-message", "test"]
  secrets = ["TWITTER_APP_KEY", "TWITTER_APP_SECRET", "TWITTER_ACCESS_TOKEN", "TWITTER_ACCESS_SECRET"]
}
```

# Install locally

```
go get github.com/thomasjungblut/twitter-action
```

# Usage 

## Commandline with env variables
```
export TWITTER_APP_KEY=xxx
export TWITTER_APP_SECRET=xxx
export TWITTER_ACCESS_TOKEN=xxx
export TWITTER_ACCESS_SECRET=xxx
twitter-action -message "Hello Twitter :)"
```

