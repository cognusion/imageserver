# Image Server
An image server toolkit in Go (Golang)

Was written by [Pierre Durand](https://github.com/pierrre/) but abandoned. I plan on keeping up with this.

[![GoDoc](https://godoc.org/github.com/cognusion/imageserver?status.svg)](https://godoc.org/github.com/cognusion/imageserver)

## Features
- HTTP server
- Resize ([GIFT](https://github.com/disintegration/gift), [nfnt resize](https://github.com/nfnt/resize), [Graphicsmagick](http://www.graphicsmagick.org/))
- Rotate
- Crop
- Convert (JPEG, GIF (animated), PNG , BMP, TIFF, ...)
- Cache ([groupcache](https://github.com/golang/groupcache), [Redis](https://github.com/garyburd/redigo), [Memcache](https://github.com/bradfitz/gomemcache), in memory)
- Gamma correction
- HMAC URL Signer/Verifier
- AWS S3 Source
- Fully modular

## Examples
- [Simple](https://github.com/cognusion/imageserver/blob/master/examples/simple/simple.go)
- [Advanced](https://github.com/cognusion/imageserver/blob/master/examples/advanced/advanced.go)
- [Cache](https://github.com/cognusion/imageserver/blob/master/examples/cache/cache.go)
- [Groupcache](https://github.com/cognusion/imageserver/blob/master/examples/groupcache/groupcache.go)
- [HTTP Source](https://github.com/cognusion/imageserver/blob/master/examples/httpsource/httpsource.go)
- [Mandelbrot](https://github.com/pierrre/mandelbrot/blob/master/examples/httpserver/httpserver.go) 

## Backward compatibility
Starting with v1.0, the API is guaranteed through v1.
