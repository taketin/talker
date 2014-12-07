package main

import (
	"os"
	"fmt"
	"strings"
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
	app := cli.NewApp()
	app.Name = "talker"
	app.Version = Version
	app.Usage = ""
	app.Author = ""
	app.Email = ""
    app.Commands = Commands
    
	app.Run(os.Args)
	if _, err = toml.DecodeFile("config.tml", &Configs); err != nil {
		panic(err)
	}

	members := strings.Split(Configs.Talker.Members, ",")
	talker := chooseTalker(members)
	fmt.Printf("%s \n", talker)
}
