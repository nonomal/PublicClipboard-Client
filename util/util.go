package util

import (
	"PublicClipboard-Client/model"
	"bytes"
	"encoding/json"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
)

var getUrl, updUrl string

func GetRemoteContent() (content string) {
	res, _ := http.Get(getUrl)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(res.Body)
	body, _ := io.ReadAll(res.Body)
	var result model.Result
	json.Unmarshal(body, &result)
	return result.Clipboard.Msg
}

func UpdRemoteContent(content string) bool {
	info := make(map[string]string)
	info["content"] = content
	bytesData, _ := json.Marshal(info)
	reader := bytes.NewReader(bytesData)
	http.Post(updUrl, "application/json;charset=UTF-8", reader)
	return true
}

func init() {
	viper.SetConfigName("conf")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./conf")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("viper load fail ...")
		return
	}
	getUrl = viper.GetString("url.getUrl")
	updUrl = viper.GetString("url.updUrl")
}
