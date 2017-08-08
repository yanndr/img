package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"sync"

	"github.com/yanndr/imgresize"
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
			err := imgresize.Resize(f, *sizePtr)
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
      '--( : )--'
        (  :  )
      ""'-...-'""`)

}
