package cmd

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
	"sync"

	"github.com/nfnt/resize"
	"github.com/spf13/cobra"
)

var (
	size              int
	width, height     uint
	outputDir, format string
)

var resizeCmd = &cobra.Command{
	Use:   "resize [files ...]",
	Short: "Resize the images to a percentage or a value.",
	RunE: func(cmd *cobra.Command, args []string) error {
		var files []string

		for i := 0; i < len(args); i++ {
			files = append(files, args[i])
		}

		var wg sync.WaitGroup
		for _, f := range files {
			wg.Add(1)
			go func(f string) {
				defer wg.Done()
				processFile(f, format, outputDir, size, width, height)
			}(f)
		}

		wg.Wait()

		return nil
	},
	Args: cobra.MinimumNArgs(1),
}

func init() {
	resizeCmd.Flags().StringVarP(&outputDir, "out", "o", "out", "Output directory for the images defautl: out.")
	resizeCmd.Flags().StringVarP(&format, "format", "f", "", "Force the format of the output: png or jpg. if empty it will keep the input image format.")
	resizeCmd.Flags().IntVarP(&size, "size", "s", 50, "Size in % of the original image.")
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
