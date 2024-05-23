package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)

	os.Mkdir("../../static/js", os.ModePerm)

	// Bundle htmx.org min
	{
		fd, err := os.Open("../../node_modules/htmx.org/dist/htmx.min.js")
		if err != nil {
			log.Fatalln(err)
		}

		defer fd.Close()

		bundled, err := os.Create("../../static/js/htmx.min.js")
		if err != nil {
			log.Fatalln(err)
		}

		defer bundled.Close()

		io.Copy(bundled, fd)
	}
}
