package main

import (
	"fmt"
	"strings"
	"github.com/BurntSushi/toml"
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

	members := strings.Split(Configs.Talker.Members, ",")
	talker := chooseTalker(members)
	fmt.Printf("%s \n", talker)
}
