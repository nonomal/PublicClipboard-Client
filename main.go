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
		err := clipboard.WriteAll(util.GetRemoteContent())
		log.Println("local <<<=== remote 同步了远端剪切板")
		if err != nil {
			log.Println("启动失败，本地剪切板写入失败,", err)
			return
		}
	} else {
		util.UpdRemoteContent(lastContent)
		log.Println("local ===>>> remote 更新了远端剪切板")
	}
	sleepTime, _ := strconv.Atoi(viper.GetString("sleepTime"))
	fmt.Println("程序运行中...")
	for {
		time.Sleep(time.Millisecond * time.Duration(sleepTime))
		local, _ := clipboard.ReadAll()
		remote := util.GetRemoteContent()
		if local != "" && remote != "" && local != remote {
			if local == lastContent {
				err := clipboard.WriteAll(remote)
				if err != nil {
					log.Println("本地剪切板写入失败,", err)
					continue
				}
				lastContent = remote
				log.Println("local <<<=== remote 同步了远端剪切板")
			} else {
				util.UpdRemoteContent(local)
				lastContent = local
				log.Println("local ===>>> remote 更新了远端剪切板")
			}
		}
	}
}
