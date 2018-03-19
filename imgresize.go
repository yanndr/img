package main

// func main() {

// 	var versionPtr bool

// 	sizePtr := flag.Int("s", 50, "Size in % of the original image.")
// 	outputPtr := flag.String("o", "out", "Output directory for the images defautl: out.")
// 	widthPtr := flag.Uint("w", 0, "Width in px of the original image.")
// 	heightPtr := flag.Uint("h", 0, "Height in px of the original image.")
// 	formatPtr := flag.String("f", "", "Force the format of the output: png or jpg. if empty it will keep the input image format.")
// 	flag.BoolVar(&versionPtr, "version", false, "Display the version")
// 	flag.BoolVar(&versionPtr, "v", false, "Display the version")

// 	flag.Usage = func() {
// 		fmt.Fprintf(os.Stdout, "%s - Image resize %s\n", os.Args[0], version)
// 		fmt.Fprintf(os.Stdout, "%s [file ...] [options]\n", os.Args[0])
// 		flag.PrintDefaults()
// 	}

// 	flag.Parse()

// 	if versionPtr {
// 		fmt.Println("imgresize version: ", version)
// 		fmt.Println(signature)
// 		return
// 	}

// 	format := strings.ToLower(*formatPtr)

// 	if format != "" && (format != "png" && format != "jpg" && format != "jpeg") {
// 		log.Fatal("error, unknow output format ", format)
// 	}

// 	var files []string
// 	if len(flag.Args()) < 1 {

// 		flag.Usage()
// 		return

// 	}

// 	for i := 0; i < len(flag.Args()); i++ {
// 		files = append(files, flag.Args()[i])
// 	}

// 	var wg sync.WaitGroup
// 	for _, f := range files {
// 		wg.Add(1)
// 		go func(f string) {
// 			defer wg.Done()
// 			processFile(f, format, *outputPtr, *sizePtr, *widthPtr, *heightPtr)
// 		}(f)
// 	}

// 	wg.Wait()
// 	fmt.Println("done")
// }
