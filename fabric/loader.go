package fabric

import (
	"client-downloader/util"
	"client-downloader/vanilla"
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"strings"
)

func (m *Manifest) Download(path string) {
	var vanillaManifest vanilla.Manifest
	vanillaManifest.Download(path)

	var loaderVersion string
	flag.StringVar(&loaderVersion, "fabric-loader", "0.15.3", "version of fabric loader")
	flag.Parse()

	versionId := flag.Lookup("version").Value.String()

	if err := util.GetFromJson(fmt.Sprintf("https://meta.fabricmc.net/v2/versions/loader/%v/%v/profile/json", versionId, loaderVersion), m); err != nil {
		log.Panicf("Manifest fabric loading error: %v", err)
	}

	file, archive := util.OpenOrCreateZip(path, false)

	defer file.Close()
	defer archive.Close()

	for _, library := range m.Libraries {
		parts := strings.SplitN(library.Name, ":", 3)
		url := fmt.Sprintf("%v/%v/%v/%v/%[3]v-%[4]v.jar", library.Url, strings.ReplaceAll(parts[0], ".", "/"), parts[1], parts[2])
		path := strings.ReplaceAll(filepath.Join("libraries", strings.ReplaceAll(parts[0], ".", "/"), parts[1], parts[2], fmt.Sprintf("%v-%v.jar", parts[1], parts[2])), " ", "_")
		util.DownloadAndAddToZip(url, path, archive)
	}
}
