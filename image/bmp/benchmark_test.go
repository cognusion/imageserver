package bmp

import (
	"testing"

	"github.com/cognusion/imageserver"
	imageserver_image_test "github.com/cognusion/imageserver/image/_test"
	_ "github.com/cognusion/imageserver/image/jpeg"
	"github.com/cognusion/imageserver/testdata"
)

func Benchmark(b *testing.B) {
	enc := &Encoder{}
	params := imageserver.Params{}
	for _, tc := range []struct {
		name string
		im   *imageserver.Image
	}{
		{"Small", testdata.Small},
		{"Medium", testdata.Medium},
		{"Large", testdata.Large},
		{"Huge", testdata.Huge},
	} {
		b.Run(tc.name, func(b *testing.B) {
			imageserver_image_test.BenchmarkEncoder(b, enc, tc.im, params)
		})
	}
}
