package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/evanoberholster/imagemeta"
)

var tZero time.Time

func main() {
	log.SetFlags(0)

	flagDoIt := flag.Bool("doit", false, "Do it")
	flagFixExt := flag.Bool("fixext", false, "Correct extension to .jpg")
	flag.Parse()

	for _, path := range flag.Args() {
		if getCT(path) != "image/jpeg" {
			log.Println(path, "not a JPEG image")
			continue
		}

		creationDate, err := getCreat(path)
		if err != nil {
			log.Println(path, err)
			continue
		}

		log.Println(path, "exif creation date", creationDate)

		if *flagDoIt {
			err = os.Chtimes(path, tZero, creationDate)
			if err != nil {
				log.Println(path, err)
				continue
			}

			if *flagFixExt && filepath.Ext(path) != ".jpg" {
				ext := filepath.Ext(path)
				newPath, _ := strings.CutSuffix(path, ext)

				err = os.Rename(path, newPath+".jpg")
				if err != nil {
					log.Println(path, err)
					continue
				}
			}
		}
	}
}

func getCT(path string) string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	buf := make([]byte, 512)

	io.ReadAtLeast(f, buf, 512)

	return http.DetectContentType(buf)
}

func getCreat(path string) (t time.Time, err error) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()

	data, err := imagemeta.Decode(f)
	if err != nil {
		return
	}

	t = data.CreateDate()
	return
}
