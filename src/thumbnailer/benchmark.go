package thumbnailer

import (
	"fmt"
	"image/jpeg"
	"os"
	"os/exec"
	"strings"

	"github.com/nfnt/resize"
)

func ThumbnailWithVips(dst, src string, w, h uint) error {
	// vipsthumbnail can be installed on ubuntu from package libvips-tools
	c := exec.Command(
		"vipsthumbnail",
		"-o", strings.Replace(dst, "%", "%%", -1),
		"-s", fmt.Sprintf("%dx%d", w, h),
		src,
	)
	err := c.Start()
	if err != nil {
		return err
	}
	return c.Wait()
}

func ThumbnailWithNFNT(dst, src string, width, height uint) error {
	r, err := os.Open(src)
	if err != nil {
		return err
	}
	defer r.Close()
	img, err := jpeg.Decode(r)
	if err != nil {
		return err
	}
	t := resize.Thumbnail(width, height, img, resize.NearestNeighbor)

	w, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer w.Close()

	return jpeg.Encode(w, t, nil)
}
