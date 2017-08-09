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
	formatPtr := flag.String("f", "", "Force the format of the output: png or jpg. if empty it will keep the input image format.")

	flag.Parse()

	format := strings.ToLower(*formatPtr)

	if format != "" && (format != "png" && format != "jpg" && format != "jpeg") {
		log.Fatal("error, unknow output format ", format)
	}

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
			processFile(f, format, *outputPtr, *sizePtr, *widthPtr, *heightPtr)
		}(f)
	}

	wg.Wait()
	fmt.Println("done")

	fmt.Println(signature)

}

func processFile(file, format, outputDir string, size int, width, height uint) {
	var err error
	var outputImg image.Image

	inputImg, _, err := openImage(file)
	if err != nil {
		fmt.Println("Error with ", file, ": ", err)
		return
	}

	if height == 0 && width == 0 {
		width := uint(inputImg.Bounds().Size().X * size / 100)
		outputImg = resize.Resize(width, 0, inputImg, resize.Lanczos3)
	} else {
		outputImg = resize.Resize(width, height, inputImg, resize.Lanczos3)
	}

	var outpufilename string
	if format != "" {
		ext := path.Ext(file)
		outpufilename = strings.Replace(file, ext, fmt.Sprintf(".%v", format), -1)
		fmt.Println(outpufilename)
	} else {

		outpufilename = file
	}

	if err != nil {
		fmt.Println("Error with ", file, ": ", err)
		return
	}

	err = save(outputDir, outpufilename, outputImg)
	if err != nil {
		fmt.Println("Error with ", outpufilename, ": ", err)
		return
	}

	fmt.Println(outpufilename, "-> Processed.")
}

func encode(extension string, w io.Writer, m image.Image) error {
	ext := strings.ToLower(extension)

	if ext == ".jpg" || ext == ".jpeg" {
		return jpeg.Encode(w, m, nil)
	}
	if ext == ".png" {
		return png.Encode(w, m)
	}

	return errors.New("unknown format")
}

func openImage(file string) (image.Image, string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, "", err
	}

	defer f.Close()
	img, ext, err := image.Decode(f)
	if err != nil {
		return nil, "", err
	}

	return img, ext, nil

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

	outputFile, err := os.Create(fmt.Sprintf("%v/%v", dir, filename))
	if err != nil {
		return err
	}
	defer outputFile.Close()

	return encode(ext, outputFile, img)
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
