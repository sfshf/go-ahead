package main

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assetsbfa8d115ce0617d89507412d5393a462f8e9b003 = "<!doctype html>\n\n<body>\n    <p>Can you see this? → {{.Bar}}</p>\n</body>\n"
var _Assets3737a75b5254ed1f6d588b40a3449721f9ea86c2 = "<!doctype html>\n\n<body>\n    <p>Hello, {{.Foo}}</p>\n</body>\n"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"html"}, "/html": []string{"bar.tmpl", "index.tmpl"}}, map[string]*assets.File{
	"/html/index.tmpl": &assets.File{
		Path:     "/html/index.tmpl",
		FileMode: 0x1ff,
		Mtime:    time.Unix(1579943457, 1579943457234902300),
		Data:     []byte(_Assets3737a75b5254ed1f6d588b40a3449721f9ea86c2),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1579943325, 1579943325528486500),
		Data:     nil,
	}, "/html": &assets.File{
		Path:     "/html",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1579943417, 1579943417041652200),
		Data:     nil,
	}, "/html/bar.tmpl": &assets.File{
		Path:     "/html/bar.tmpl",
		FileMode: 0x1ff,
		Mtime:    time.Unix(1579943433, 1579943433583704600),
		Data:     []byte(_Assetsbfa8d115ce0617d89507412d5393a462f8e9b003),
	}}, "")
