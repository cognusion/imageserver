// Package cache provides a cache example.
//
// It shows how to use a single cache or several caches together.
package main

import (
	"crypto/sha256"
	"flag"
	"net/http"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/disintegration/gift"
	"github.com/go-redis/redis"
	"github.com/cognusion/imageserver"
	imageserver_cache "github.com/cognusion/imageserver/cache"
	imageserver_cache_file "github.com/cognusion/imageserver/cache/file"
	imageserver_cache_memcache "github.com/cognusion/imageserver/cache/memcache"
	imageserver_cache_memory "github.com/cognusion/imageserver/cache/memory"
	imageserver_cache_redis "github.com/cognusion/imageserver/cache/redis"
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

var (
	flagHTTP     = ":8080"
	flagMemory   = int64(128 * (1 << 20))
	flagRedis    = "localhost:6379"
	flagMemcache = "localhost:11211"
	flagFile     = ""
)

func main() {
	parseFlags()
	startHTTPServer()
}

func parseFlags() {
	flag.StringVar(&flagHTTP, "http", flagHTTP, "HTTP")
	flag.Int64Var(&flagMemory, "memory", flagMemory, "Memory")
	flag.StringVar(&flagRedis, "redis", flagRedis, "Redis")
	flag.StringVar(&flagMemcache, "memcache", flagMemcache, "Memcache")
	flag.StringVar(&flagFile, "file", flagFile, "File")
	flag.Parse()
}

func startHTTPServer() {
	http.Handle("/", http.StripPrefix("/", newImageHTTPHandler()))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(flagHTTP, nil)
	if err != nil {
		panic(err)
	}
}

func newImageHTTPHandler() http.Handler {
	return &imageserver_http.Handler{
		Parser: imageserver_http.ListParser([]imageserver_http.Parser{
			&imageserver_http.SourcePathParser{},
			&imageserver_http_gift.ResizeParser{},
			&imageserver_http_image.FormatParser{},
			&imageserver_http_image.QualityParser{},
		}),
		Server: newServer(),
	}
}

func newServer() imageserver.Server {
	srv := imageserver_testdata.Server
	srv = newServerImage(srv)
	srv = newServerFile(srv)
	srv = newServerMemcache(srv)
	srv = newServerRedis(srv)
	srv = newServerMemory(srv)
	return srv
}

func newServerImage(srv imageserver.Server) imageserver.Server {
	return &imageserver.HandlerServer{
		Server: srv,
		Handler: &imageserver_image.Handler{
			Processor: &imageserver_image_gift.ResizeProcessor{
				DefaultResampling: gift.LanczosResampling,
			},
		},
	}
}

func newServerMemcache(srv imageserver.Server) imageserver.Server {
	if flagMemcache == "" {
		return srv
	}
	cl := memcache.New(flagMemcache)
	var cch imageserver_cache.Cache = &imageserver_cache_memcache.Cache{Client: cl}
	cch = &imageserver_cache.IgnoreError{Cache: cch}
	cch = &imageserver_cache.Async{Cache: cch}
	kg := imageserver_cache.NewParamsHashKeyGenerator(sha256.New)
	return &imageserver_cache.Server{
		Server:       srv,
		Cache:        cch,
		KeyGenerator: kg,
	}
}

func newServerRedis(srv imageserver.Server) imageserver.Server {
	if flagRedis == "" {
		return srv
	}
	client := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs: []string{"localhost:6379"},
	})
	var cch imageserver_cache.Cache = &imageserver_cache_redis.Cache{
		Client: client,
		Expire: 7 * 24 * time.Hour,
	}
	cch = &imageserver_cache.IgnoreError{Cache: cch}
	cch = &imageserver_cache.Async{Cache: cch}
	kg := imageserver_cache.NewParamsHashKeyGenerator(sha256.New)
	kg = &imageserver_cache.PrefixKeyGenerator{
		KeyGenerator: kg,
		Prefix:       "image:",
	}
	return &imageserver_cache.Server{
		Server:       srv,
		Cache:        cch,
		KeyGenerator: kg,
	}
}

func newServerMemory(srv imageserver.Server) imageserver.Server {
	if flagMemory <= 0 {
		return srv
	}
	cch := imageserver_cache_memory.New(flagMemory)
	kg := imageserver_cache.NewParamsHashKeyGenerator(sha256.New)
	return &imageserver_cache.Server{
		Server:       srv,
		Cache:        cch,
		KeyGenerator: kg,
	}
}

func newServerFile(srv imageserver.Server) imageserver.Server {
	if flagFile == "" {
		return srv
	}
	cch := imageserver_cache_file.Cache{Path: flagFile}
	kg := imageserver_cache.NewParamsHashKeyGenerator(sha256.New)
	return &imageserver_cache.Server{
		Server:       srv,
		Cache:        &cch,
		KeyGenerator: kg,
	}
}
