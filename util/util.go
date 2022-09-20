package util

import (
	"PublicClipboard-Client/model"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
)

var readUrl, writeUrl string

func GetRemoteContent() (content string) {
	res, err := http.Get(readUrl)
	if err != nil {
		return ""
	}
	defer res.Body.Close()
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
	http.Post(writeUrl, "application/json;charset=UTF-8", reader)
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
	readUrl = viper.GetString("url.readUrl")
	writeUrl = viper.GetString("url.writeUrl")
	fmt.Println("配置文件读取成功!")
	fmt.Println("readUrl:", readUrl)
	fmt.Println("writeUrl:", writeUrl)
}
