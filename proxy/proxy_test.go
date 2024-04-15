package proxy

import (
	"testing"
)

func Test(t *testing.T) {
	t.Run("proxy-statics: ", StaticsProxy)
}

func StaticsProxy(t *testing.T) {
	url := "http://baidu.com"
	download := &ConcreteDownload{Url: url}
	proxy := &DownloadProxy{
		Url:        url,
		downloader: download,
	}
	proxy.Download()
}
