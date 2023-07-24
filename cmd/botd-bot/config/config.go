package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/13-bit/bulletinbird/util"
)

var configDir,
	configFilePath string

type Config struct {
	BotdUrl string
}

var cfg Config

func init() {
	homeDir, _ := os.UserHomeDir()
	configDir = fmt.Sprintf("%s/.botd-bot", homeDir)

	configFilePath = fmt.Sprintf("%s/config.json", configDir)

	if !checkConfigExists() {
		log.Printf("%s does not exist, creating...\n", configDir)
		createConfig()
	}

	loadConfig()
}

func checkConfigExists() bool {
	_, err := os.Stat(configDir)

	return !os.IsNotExist(err)
}

func createConfig() {
	if err := os.Mkdir(configDir, os.FileMode(0775)); err != nil {
		util.CheckError(err)
	}

	f, err := os.Create(configFilePath)
	util.CheckError(err)

	defer f.Close()

	cfg := Config{
		BotdUrl: "http://BULLETINBIRDURL:1313/botd",
	}

	cfgJson, _ := json.MarshalIndent(cfg, "", "  ")

	f.Write(cfgJson)

	log.Printf("Config created. Edit BulletinBird server URL in %s...\n", configFilePath)

	os.Exit(0)
}

func loadConfig() {
	f, err := os.Open(configFilePath)
	util.CheckError(err)

	defer f.Close()

	err = json.NewDecoder(f).Decode(&cfg)
	util.CheckError(err)
}

func BotdUrl() string {
	return cfg.BotdUrl
}
