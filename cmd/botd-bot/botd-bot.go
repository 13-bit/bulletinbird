package main

import (
	"fmt"

	"github.com/13-bit/bulletinbird/cmd/botd-bot/post"
)

func main() {
	post := post.NewPost()

	fmt.Println(post)
}
