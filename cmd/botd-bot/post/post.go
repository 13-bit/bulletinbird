package post

import (
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/13-bit/bulletinbird/birds"
	"github.com/13-bit/bulletinbird/botd"
	"github.com/13-bit/bulletinbird/cmd/botd-bot/config"
	"github.com/13-bit/bulletinbird/util"
)

type PostTemplate struct {
	LifeHistoryMapping map[string]string `json:"lifeHistoryMapping"`
	PostTextTemplates  []string          `json:"postTextTemplates"`
}

type Post struct {
	Body string
	Url  string
}

var postTemplate PostTemplate

//go:embed resources/*
var resourcesFS embed.FS

func init() {
	f, err := resourcesFS.Open("resources/bot-post-text.json")
	util.CheckError(err)

	defer f.Close()

	err = json.NewDecoder(f).Decode(&postTemplate)
	util.CheckError(err)

	rand.Seed(time.Now().UnixNano())
}

func NewPost() Post {
	var post Post

	resp, err := http.Get(config.BotdUrl())
	util.CheckError(err)

	defer resp.Body.Close()
	bodyBytes, _ := io.ReadAll(resp.Body)

	var botd botd.Botd
	json.Unmarshal(bodyBytes, &botd)

	post.Url = botd.Bird.GuideUrl
	post.Body = generatePostBody(botd.Bird)

	return post
}

func generatePostBody(bird birds.Bird) string {
	randIndex := rand.Intn(len(postTemplate.PostTextTemplates))
	time := time.Now()

	fmt.Println(bird)

	body := postTemplate.PostTextTemplates[randIndex]
	body = strings.ReplaceAll(body, "{{date}}", time.Format("Monday, January 2, 2006"))
	body = strings.ReplaceAll(body, "{{common-name}}", bird.CommonName)
	body = strings.ReplaceAll(body, "{{scientific-name}}", bird.ScientificName)
	body = strings.ReplaceAll(body, "{{habitat}}", postTemplate.LifeHistoryMapping[bird.Habitat])
	body = strings.ReplaceAll(body, "{{nest}}", postTemplate.LifeHistoryMapping[bird.Nesting])
	body = strings.ReplaceAll(body, "{{diet}}", postTemplate.LifeHistoryMapping[bird.Food])

	return body
}
