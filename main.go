package main

import (
	"client-downloader/fabric"
	"client-downloader/loader"
	"client-downloader/vanilla"
	"errors"
	"flag"
	"os"
)

func getManifest(val string) (loader.Loader, error) {
	switch val {
	case "vanilla":
		return &vanilla.Manifest{}, nil
	case "fabric":
		return &fabric.Manifest{}, nil
	default:
		return nil, errors.New("unknown type of client")
	}
}

func main() {
	var archivePath string
	var manifest loader.Loader
	flag.StringVar(&archivePath, "archive", "test.zip", "location archive")
	flag.Func("type", "type of client (vanilla, fabric)", func(flagValue string) error {
		var err error
		manifest, err = getManifest(flagValue)
		return err
	})
	flag.Parse()

	if manifest == nil {
		flag.Usage()
		os.Exit(0)
	}

	manifest.Download(archivePath)
}
