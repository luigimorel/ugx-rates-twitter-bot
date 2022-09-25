# Introduction

A twitter bot that performs currency conversion and automatically makes a tweet every after 24 hours at 11 AM UTC timezone by relying on a cron job.

## Tools

- Golang
- [API Layer currency convert endpoint](https://api.apilayer.com/)
- [gocron](github.com/go-co-op/gocron) - a package for cron jobs
- [godotenv](github.com/joho/godotenv) - for handling env variables
- [Twitter API wrapper](github.com/dghubble/go-twitter) - for handling twitter events
- [OAuth1](github.com/dghubble/oauth1) - provides a Go implementation of the OAuth 1 spec to allow end-users to authorize a client

### Resources

- [Twitter developers' documentation](https://developer.twitter.com/en)
- [API Layer](https://api.apilayer.com/)
