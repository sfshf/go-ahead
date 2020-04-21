package templatemethod

import "fmt"

/*
模版方法模式

模版方法模式使用继承机制，把通用步骤和通用方法放到父类中，把具体实现延迟到子类中实现。
使得实现符合开闭原则。

如实例代码中通用步骤在父类中实现（准备、下载、保存、收尾）下载和保存的具体实现留到
子类中，并且提供保存方法的默认实现。

因为Golang不提供继承机制，需要使用匿名组合模拟实现继承。

此处需要注意：因为父类需要调用子类方法，所以子类需要匿名组合父类的同时，父类需要持有
子类的引用。
*/

type Downloader interface {
    Download(uri string)
}

type implement interface {
    download()
    save()
}

type template struct {
    implement
    uri string
}

func (t *template) Download(uri string) {
    t.uri = uri
    fmt.Println("Prepare downloading ...")
    t.implement.download()
    t.implement.save()
    fmt.Println("Finish downloading!")
}

func newTemplate(impl implement) *template {
    return &template{
        implement: impl,
    }
}

type HTTPDownloader struct {
    *template
}

func (hd *HTTPDownloader) download() {
    fmt.Printf("[HTTP] downloading from <%s> ...\n", hd.template.uri)
}

func (*HTTPDownloader) save() {
    fmt.Println("[HTTP] saving ...")
}

func NewHTTPDownloader() Downloader {
    downloader := &HTTPDownloader{}
    downloader.template = newTemplate(downloader)
    return downloader
}

type FTPDownloader struct {
    *template
}

func (fd *FTPDownloader) download() {
    fmt.Printf("[FTP] downloading from <%s> ...\n", fd.template.uri)
}

func (*FTPDownloader) save() {
    fmt.Println("[FTP] saving ...")
}

func NewFTPDownloader() Downloader {
    downloader := &FTPDownloader{}
    downloader.template = newTemplate(downloader)
    return downloader
}
