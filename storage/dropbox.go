package dropbox

import (
	"github.com/stacktic/dropbox"
)

var appKey string
var appSecret string
var token string

var client *dropbox.Dropbox
var entry *dropbox.Entry
var err error

func SetConfig(aAppKey string, aAppSecret string, aToken string) {
	appKey = aAppKey
	appSecret = aAppSecret
	token = aToken
}

func DownloadFile(src string, dst string) {
	connect()
	client.DownloadToFile(src, dst, "")
}

func UploadFile(src string, dst string) {
	connect()
	if entry, err = client.UploadFile(dst, src, true, ""); err != nil {
		panic(err)
	}
}

func connect() {
	if client == nil {
		client = dropbox.NewDropbox()
		client.SetAppInfo(appKey, appSecret)
		client.SetAccessToken(token)
	}
}
