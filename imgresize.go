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

	nresize "github.com/nfnt/resize"
)

const signature = `
 	 _[_]_  
          (")  
      '--( : )--'
        (  :  )
      ""'-...-'""
`

func main() {
	sizePtr := flag.Int("s", 50, "Size in % of the original image.")
	outputPtr := flag.String("o", "out", "Output directory for the images defautl: out.")
	widthPtr := flag.Uint("w", 0, "Width in px of the original image.")
	heightPtr := flag.Uint("h", 0, "Height in px of the original image.")

	flag.Parse()

	var files []string
	if len(flag.Args()) < 1 {
		input := "./*.*"
		var err error
		files, err = filepath.Glob(input)
		if err != nil {
			log.Fatal(err)
			return
		}

	} else {
		for i := 0; i < len(flag.Args()); i++ {
			files = append(files, flag.Args()[i])
		}
	}

	var wg sync.WaitGroup
	for _, f := range files {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			var err error
			var img image.Image

			if *heightPtr == 0 && *widthPtr == 0 {
				img, err = resizeInPercent(f, *sizePtr)
			} else {
				img, err = resize(f, *widthPtr, *heightPtr)
			}

			if err != nil {
				fmt.Println("Error with ", f, ": ", err)
				return
			}

			err = save(*outputPtr, f, img)
			if err != nil {
				fmt.Println("Error with ", f, ": ", err)
				return
			}

			fmt.Println(f, "-> Processed.")
		}(f)

	}

	wg.Wait()
	fmt.Println("done")

	fmt.Println(signature)

}

func encode(extenstion string, w io.Writer, m image.Image) error {
	ext := strings.ToLower(extenstion)

	if ext == ".jpg" || ext == ".jpeg" {
		return jpeg.Encode(w, m, nil)
	}
	if ext == ".png" {
		return png.Encode(w, m)
	}

	return errors.New("unknown format")
}

func resizeInPercent(f string, size int) (image.Image, error) {

	file, err := os.Open(f)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	img, st, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	fmt.Println(st)

	s := uint(img.Bounds().Size().X * size / 100)
	result := nresize.Resize(s, 0, img, nresize.Lanczos3)

	return result, nil
}

func resize(f string, width, height uint) (image.Image, error) {

	file, err := os.Open(f)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	img, st, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	fmt.Println(st)

	result := nresize.Resize(width, height, img, nresize.Lanczos3)

	return result, nil
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
