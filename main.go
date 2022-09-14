package main

import (
	"PublicClipboard-Client/model"
	"bytes"
	"encoding/json"
	"github.com/atotto/clipboard"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var getUrl, updUrl string

func GetContent() (content string) {
	res, _ := http.Get(getUrl)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var result model.Result
	json.Unmarshal(body, &result)
	return result.Clipboard.Msg
}

func UpdContent(content string) bool {
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
	viper.AddConfigPath("conf")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("viper load fail ...")
		return
	}
	getUrl = viper.GetString("url.getUrl")
	updUrl = viper.GetString("url.updUrl")
}

func main() {
	lastContent, _ := clipboard.ReadAll()
	if viper.GetString("init") == "1" {
		clipboard.WriteAll(GetContent())
	} else {
		UpdContent(lastContent)
	}
	for {
		time.Sleep(time.Second * 5)
		local, _ := clipboard.ReadAll()
		remote := GetContent()

		if local != "" && remote != "" && local != remote {
			if local == lastContent {
				clipboard.WriteAll(remote)
				lastContent = remote
				log.Println("local 《====== remote 同步了远端剪切板")
			} else {
				UpdContent(local)
				lastContent = local
				log.Println("local ======》 remote 更新了远端剪切板")
			}
		}
	}
}
