package thumbnailer

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func BenchmarkVips(b *testing.B) {
	benchmark(b, ThumbnailWithVips)
}

func BenchmarkNFNT(b *testing.B) {
	benchmark(b, ThumbnailWithNFNT)
}

func benchmark(b *testing.B, fn func(string, string, uint, uint) error) {
	dir, _ := ioutil.TempDir("", "")
	for i := 0; i < b.N; i++ {
		if err := fn(filepath.Join(dir, fmt.Sprintf("totem_640_480_%d.jpg", i)), "totem.jpg", 640, 480); err != nil {
			b.Log(err)
			b.FailNow()
		}
	}
}
