package proxy

import "fmt"

// Downloader 下载器接口
type Downloader interface {
	// Download 下载
	Download()
}

// ConcreteDownload 具体的下载结构体，实现Downloader接口
type ConcreteDownload struct {
	Url string
}

// Download 下载
func (download ConcreteDownload) Download() {
	fmt.Println(fmt.Sprintf("%s 在下载中", download.Url))
}

// DownloadProxy，下载代理，实现Downloader接口
type DownloadProxy struct {
	Url        string
	downloader Downloader
}

// Download 下载
func (proxy DownloadProxy) Download() {
	fmt.Println(fmt.Sprintf("准备开始下载%s", proxy.Url))
	proxy.downloader.Download()
	fmt.Println(fmt.Sprintf("下载%s完成", proxy.Url))
}
