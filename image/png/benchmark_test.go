package png

import (
	"image/png"
	"testing"

	"github.com/cognusion/imageserver"
	imageserver_image_test "github.com/cognusion/imageserver/image/_test"
	_ "github.com/cognusion/imageserver/image/jpeg"
	"github.com/cognusion/imageserver/testdata"
)

func BenchmarkSize(b *testing.B) {
	enc := &Encoder{}
	for _, tc := range []struct {
		name string
		im   *imageserver.Image
	}{
		{"Small", testdata.Small},
		{"Medium", testdata.Medium},
		{"Large", testdata.Large},
		{"Huge", testdata.Huge},
	} {

		benchmark(b, tc.name, enc, tc.im)
	}
}

func BenchmarkCompressionLevel(b *testing.B) {
	for _, tc := range []struct {
		name string
		cl   png.CompressionLevel
	}{
		{"DefaultCompression", png.DefaultCompression},
		{"NoCompression", png.NoCompression},
		{"BestSpeed", png.BestSpeed},
		{"BestCompression", png.BestCompression},
	} {
		enc := &Encoder{
			CompressionLevel: tc.cl,
		}
		benchmark(b, tc.name, enc, testdata.Medium)
	}
}

func benchmark(b *testing.B, name string, enc *Encoder, im *imageserver.Image) {
	b.Run(name, func(b *testing.B) {
		imageserver_image_test.BenchmarkEncoder(b, enc, im, imageserver.Params{})
	})
}
