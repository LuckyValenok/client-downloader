package main

import (
	"client-downloader/fabric"
	"client-downloader/forge"
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
	case "forge":
		return &forge.Manifest{}, nil
	default:
		return nil, errors.New("unknown type of client")
	}
}

func main() {
	var archivePath string
	var manifest loader.Loader
	flag.StringVar(&archivePath, "archive", "test", "location archive")
	flag.Func("type", "type of client (vanilla, fabric, forge)", func(flagValue string) error {
		var err error
		manifest, err = getManifest(flagValue)
		return err
	})
	flag.String("version", "release", "version minecraft")
	flag.String("fabric-loader", "0.15.3", "version of fabric loader")
	flag.String("forge-version", "47.2.20", "version of forge")
	flag.Parse()

	if manifest == nil {
		flag.Usage()
		os.Exit(0)
	}

	manifest.Download(archivePath)
}
