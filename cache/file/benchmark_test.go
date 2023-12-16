package file

import (
	"testing"

	"github.com/cognusion/imageserver"
	cachetest "github.com/cognusion/imageserver/cache/_test"
	"github.com/cognusion/imageserver/testdata"
)

func BenchmarkGet(b *testing.B) {
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
			cch := newTestCache()
			cachetest.BenchmarkGet(b, cch, 1, tc.im)
		})
	}
}
