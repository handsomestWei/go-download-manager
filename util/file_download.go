package util

import (
	"github.com/cavaliercoder/grab"
	"net/http"
	"time"
	"github.com/handsomestWei/go-download-manager/log"
)

// @see https://github.com/cavaliercoder/grab
func FileDownLoad(dst, urlStr string) bool {

	var resp *grab.Response
	// 重试三次
	for i := 1; i < 4; i++ {
		client := grab.NewClient()
		req, _ := grab.NewRequest(dst, urlStr)
		log.Infof("Downloading %v...", req.URL())
		resp := client.Do(req)
		if !isRspOk(resp) {
			time.Sleep(10 * time.Second)
			log.Errorf("Download failed, retry %d", i)
			continue
		} else {
			break
		}
	}

	if !isRspOk(resp) {
		log.Error("Download failed")
		return false
	}
	log.Infof("%v", resp.HTTPResponse.Status)
	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()

Loop:
	for {
		select {
		case <-t.C:
			log.Infof("transferred %f", resp.Progress())
		case <-resp.Done:
			// download is complete
			break Loop
		}
	}

	if err := resp.Err(); err != nil {
		log.Errorf("Download failed: %v", err)
		return false
	}

	log.Infof("Download %v saved to %s", resp.Filename, dst)
	return true
}

func isRspOk(resp *grab.Response) bool {
	if resp == nil || resp.HTTPResponse == nil || resp.HTTPResponse.StatusCode != http.StatusOK {
		return false
	} else {
		return true
	}
}
