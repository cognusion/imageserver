package _test

import (
	"io/ioutil"
	"testing"

	"github.com/cognusion/imageserver"
	imageserver_image "github.com/cognusion/imageserver/image"
)

// BenchmarkEncoder is a helper to benchmark imageserver/image.Encoder.
func BenchmarkEncoder(b *testing.B, enc imageserver_image.Encoder, im *imageserver.Image, params imageserver.Params) {
	nim, err := imageserver_image.Decode(im)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := enc.Encode(ioutil.Discard, nim, params)
		if err != nil {
			b.Fatal(err)
		}
	}
}
