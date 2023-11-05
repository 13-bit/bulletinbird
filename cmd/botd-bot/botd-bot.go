package main

import (
	"fmt"

	"github.com/13-bit/bulletinbird/cmd/botd-bot/mastodon"
	"github.com/13-bit/bulletinbird/cmd/botd-bot/post"
)

func main() {
	post := post.NewPost()

	fmt.Println(post)

	mastodon.TootBotd(post, "kRZZR4wv-JZVSbRvOAKJ2GqF_1muRjyqzeQc_W0Uggg")
}
