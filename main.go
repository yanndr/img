package main

import (
	"errors"
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"

	"github.com/nfnt/resize"
)

func main() {
	sizePtr := flag.Int("s", 50, "Size in % of the original image.")

	flag.Parse()

	var files []string
	if len(flag.Args()) < 2 {
		input := "./*.*"
		var err error
		files, err = filepath.Glob(input)
		if err != nil {
			log.Fatal(err)
			return
		}

	} else {
		for i := 1; i < len(flag.Args()); i++ {
			files = append(files, flag.Args()[i])
		}
	}

	var wg sync.WaitGroup
	for _, f := range files {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			err := resizeImg(f, *sizePtr)
			if err != nil {
				fmt.Println("Error with ", f, ": ", err)
				return
			}

			fmt.Println(f, "-> Processed.")
		}(f)

	}

	wg.Wait()
	fmt.Println("done")

	fmt.Println(`
	 _[_]_  
          (")  
       --( : )--'
        (  :  )
      "" -...-'""`)

}

func decode(extenstion string, r io.Reader) (image.Image, error) {
	ext := strings.ToLower(extenstion)
	if ext == ".jpg" || ext == ".jpeg" {
		return jpeg.Decode(r)
	}
	if ext == ".png" {
		return png.Decode(r)
	}

	return nil, errors.New("Doesn't seems to be an image")
}

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

func resizeImg(f string, size int) error {

	file, err := os.Open(f)
	if err != nil {
		return err
	}

	defer file.Close()
	ext := path.Ext(f)
	img, err := decode(ext, file)
	if err != nil {
		return err
	}
	s := uint(img.Bounds().Size().X * size / 100)
	m := resize.Resize(s, 0, img, resize.Lanczos3)
	if ok, err := exist("out"); !ok {
		if err != nil {
			return err
		}
		err := os.Mkdir("out", 0711)
		if err != nil {
			return err
		}
	}
	out, err := os.Create(fmt.Sprintf("out/%v", f))
	if err != nil {
		return err
	}
	defer out.Close()

	// write new image to file
	encode(ext, out, m)

	return nil

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
