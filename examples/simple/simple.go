// Package simple provides a simple example.
package main

import (
	"net/http"

	"github.com/cognusion/imageserver"
	imageserver_http "github.com/cognusion/imageserver/http"
	imageserver_http_gift "github.com/cognusion/imageserver/http/gift"
	imageserver_http_image "github.com/cognusion/imageserver/http/image"
	imageserver_image "github.com/cognusion/imageserver/image"
	_ "github.com/cognusion/imageserver/image/gif"
	imageserver_image_gift "github.com/cognusion/imageserver/image/gift"
	_ "github.com/cognusion/imageserver/image/jpeg"
	_ "github.com/cognusion/imageserver/image/png"
	imageserver_testdata "github.com/cognusion/imageserver/testdata"
)

func main() {
	http.Handle("/", &imageserver_http.Handler{
		Parser: imageserver_http.ListParser([]imageserver_http.Parser{
			&imageserver_http.SourceParser{},
			&imageserver_http_gift.ResizeParser{},
			&imageserver_http_image.FormatParser{},
			&imageserver_http_image.QualityParser{},
		}),
		Server: &imageserver.HandlerServer{
			Server: imageserver_testdata.Server,
			Handler: &imageserver_image.Handler{
				Processor: &imageserver_image_gift.ResizeProcessor{},
			},
		},
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
