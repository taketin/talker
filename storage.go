package main

import (
	"os"
	"bytes"
	"bufio"
	"io/ioutil"
	"github.com/taketin/talker/storage"
)

var fp *os.File
var err error

func finishedMemberFromStrage() ([]string) {
	setConfigToDropbox()
	dropbox.DownloadFile(Configs.File.RemotePath, os.Getenv("HOME") + Configs.File.LocalPath)
	return readFile(os.Getenv("HOME") + Configs.File.LocalPath)
}

func saveFinishedMember(finishedMember []string) {
	writeFile(finishedMember)
	setConfigToDropbox()
	dropbox.UploadFile(Configs.File.RemotePath, os.Getenv("HOME") + Configs.File.LocalPath)
	removeFile()
}

func setConfigToDropbox() {
	dropbox.SetConfig(Configs.Dropbox.AppKey, Configs.Dropbox.AppSecret, Configs.Dropbox.Token)
}

func readFile(src string) ([]string) {
	var finishedMember []string
	if fp, err = os.Open(os.Getenv("HOME") + Configs.File.LocalPath); err != nil {
		createFile()
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		finishedMember = append(finishedMember, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return finishedMember
}

func createFile() {
	if fp, err = os.OpenFile(os.Getenv("HOME") + Configs.File.LocalPath, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666); err != nil {
		panic(err)
	}
}

func writeFile(finishedMember []string) {
	var buffer bytes.Buffer
	for _, val := range finishedMember {
		buffer.WriteString(val + "\n")
	}
	content := []byte(buffer.String())
	if err = ioutil.WriteFile(os.Getenv("HOME") + Configs.File.LocalPath, content, os.ModePerm); err != nil {
		panic(err)
	}
}

func removeFile() {
	if err = os.Remove(os.Getenv("HOME") + Configs.File.LocalPath); err != nil {
		panic(err)
	}
}
