package main

import (
	"PublicClipboard-Client/util"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/spf13/viper"
	"log"
	"strconv"
	"time"
)

func main() {
	lastContent, _ := clipboard.ReadAll()
	if viper.GetString("initMode") == "1" {
		clipboard.WriteAll(util.GetContent())
	} else {
		util.UpdContent(lastContent)
	}
	sleepTime, _ := strconv.Atoi(viper.GetString("sleepTime"))
	fmt.Println("程序运行中...")
	for {
		time.Sleep(time.Millisecond * time.Duration(sleepTime))
		local, _ := clipboard.ReadAll()
		remote := util.GetContent()
		if local != "" && remote != "" && local != remote {
			if local == lastContent {
				clipboard.WriteAll(remote)
				lastContent = remote
				log.Println("local <<<=== remote 同步了远端剪切板")
			} else {
				util.UpdContent(local)
				lastContent = local
				log.Println("local ===>>> remote 更新了远端剪切板")
			}
		}
	}
}
