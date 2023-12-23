package util

import (
	"archive/zip"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func DownloadAndAddToZip(url, path string, archive *zip.Writer) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	header := &zip.FileHeader{
		Name:   filepath.ToSlash(path),
		Method: zip.Deflate,
	}

	writer, err := archive.CreateHeader(header)
	if err != nil {
		panic(err)
	}

	if _, err = io.Copy(writer, resp.Body); err != nil {
		panic(err)
	}
}

func GetFromJson(url string, target interface{}) error {
	time.Sleep(time.Second) // моджанги терпеть не могут частые запросы, после первого запроса все сдыхает

	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func OpenOrCreateZip(path string, recreate bool) (*os.File, *zip.Writer) {
	var zipfile *os.File
	var err error

	if recreate {
		zipfile, err = os.Create(path)
	} else {
		zipfile, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	}

	if err != nil {
		log.Panicf("Failed to create archive: %v", err)
	}

	return zipfile, zip.NewWriter(zipfile)
}
