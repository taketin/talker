package main

import (
	"os"
	"github.com/BurntSushi/toml"
	"github.com/codegangsta/cli"
)

var Configs Config

type Config struct {
	Dropbox DropboxConfig
	File FileConfig
	Talker TalkerConfig
}

type DropboxConfig struct {
	AppKey string `toml:"appKey"`
	AppSecret string `toml:"appSecret"`
	Token string `toml:"token"`
}

type FileConfig struct {
	RemotePath string `toml:"remotePath"`
	LocalPath string `toml:"localPath"`
}

type TalkerConfig struct {
	Members string `toml:"members"`
}

func main() {
	if _, err = toml.DecodeFile("config.tml", &Configs); err != nil {
		panic(err)
	}

	app := cli.NewApp()
	app.Name = "talker"
	app.Version = Version
	app.Usage = ""
	app.Author = "taketin"
	app.Email = "tksthdnr@gmail.com"
	app.Commands = Commands

	app.Run(os.Args)
}
