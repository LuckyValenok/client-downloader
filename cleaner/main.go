package main

import (
	"client-downloader/util"
	"flag"
	"log"
	"os"
	"path"
	"strings"
)

func main() {
	var dir string
	flag.StringVar(&dir, "dir", "test", "dir path")
	flag.Parse()

	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".jar") {
			if err := util.RemoveSignatureFiles(path.Join(dir, file.Name())); err != nil {
				panic(err)
			}
		}
	}
}
