package imgresize

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path"
	"strings"

	"github.com/nfnt/resize"
)

func encode(extenstion string, w io.Writer, m image.Image) error {
	ext := strings.ToLower(extenstion)

	if ext == ".jpg" || ext == ".jpeg" {
		return jpeg.Encode(w, m, nil)
	}
	if ext == ".png" {
		return png.Encode(w, m)
	}

	return errors.New("Doesn't seems to be an image")
}

func Resize(f string, size int) error {

	file, err := os.Open(f)
	if err != nil {
		return err
	}

	defer file.Close()
	img, st, err := image.Decode(file)
	if err != nil {
		return err
	}
	fmt.Println(st)

	s := uint(img.Bounds().Size().X * size / 100)
	m := resize.Resize(s, 0, img, resize.Lanczos3)

	return save("out", f, m)

}

func save(dir, filename string, img image.Image) error {

	if ok, err := exist(dir); !ok {
		if err != nil {
			return err
		}
		err := os.Mkdir(dir, 0711)
		if err != nil {
			return err
		}
	}

	ext := path.Ext(filename)
	out, err := os.Create(fmt.Sprintf("%v/%v", dir, filename))
	if err != nil {
		return err
	}
	defer out.Close()

	// write new image to file
	return encode(ext, out, img)
}

func exist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
