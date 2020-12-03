package main

import (
	"fmt"
	"github.com/shima-park/agollo"
	"os"
	"time"
	"github.com/handsomestWei/go-download-manager/config"
	"github.com/handsomestWei/go-download-manager/log"
	"github.com/handsomestWei/go-download-manager/util"
)

func main() {
	config.InitConfig("config.properties")

	err := agollo.InitWithDefaultConfigFile(
		agollo.WithLogger(agollo.NewLogger(agollo.LoggerWriter(os.Stdout))),
		agollo.AutoFetchOnCacheMiss(),
		agollo.FailTolerantOnBackupExists(),
		agollo.ConfigServerRefreshIntervalInSecond(20*time.Second),
		agollo.LongPollerInterval(20*time.Second),
	)
	if err != nil {
		panic(err)
	}
	// 连接apollo配置中心
	errorCh := agollo.Start()
	log.Infof("agollo start watch key %s, current value is %v", config.Conf.WatchKey, agollo.Get(config.Conf.WatchKey))

	stop := make(chan bool)
	watchNSCh := agollo.WatchNamespace(config.Conf.WatchNs, stop)

	// 监听指定key的变化
	go func() {
		for {
			select {
			case err := <-errorCh:
				log.Error(err)
			case resp := <-watchNSCh:
				log.Infof("Watch Namespace %s %v", config.Conf.WatchNs, resp)
				if resp.Changes != nil {
					for _, v := range resp.Changes {
						if config.Conf.WatchKey == v.Key {
							fileVersion := fmt.Sprintf("%v", v.Value)
							log.Infof("%s change to %v", config.Conf.WatchKey, fileVersion)
							util.FileDownLoad(config.Conf.DownLoadDst, formatDownLoadUrl(fileVersion))
						}
					}
				}
			case <-time.After(120 * time.Second):
				log.Infof("%s current value is %v", config.Conf.WatchKey, agollo.Get(config.Conf.WatchKey))
			}
		}
	}()

	select {}

	agollo.Stop()
}

// 例：http://www.xxx.com/public/install/1.0.0/xxx.zip
func formatDownLoadUrl(fileVersion string) string {
	return config.Conf.DownLoadUrl+fileVersion+"/"+config.Conf.DownLoadFileName
}
