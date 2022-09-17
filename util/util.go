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

var getUrl, updUrl string

func GetRemoteContent() (content string) {
	res, err := http.Get(getUrl)
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
	fmt.Println("配置文件读取成功!")
	fmt.Println("getUrl:", getUrl)
	fmt.Println("updUrl:", updUrl)
}
