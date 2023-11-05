package mastodon

import (
	"log"

	"github.com/13-bit/bulletinbird/cmd/botd-bot/post"
	"github.com/13-bit/bulletinbird/util"
	"github.com/go-resty/resty/v2"
)

var client *resty.Client

func init() {
	client = resty.New()
}

func TootBotd(post post.Post, token string) {
	resp, err := client.R().SetAuthToken(token).SetHeader("Content-Type", "application/json").SetBody(post.Body).Put("https://botsin.space/api/v1/statuses")

	log.Printf("%+v\n", resp)

	util.CheckError(err)
}
