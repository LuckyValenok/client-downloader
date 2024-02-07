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

	loaderVersion := flag.Lookup("fabric-loader").Value.String()
	versionId := flag.Lookup("version").Value.String()

	if err := util.GetFromJson(fmt.Sprintf("https://meta.fabricmc.net/v2/versions/loader/%v/%v/profile/json", versionId, loaderVersion), m); err != nil {
		log.Panicf("Manifest fabric loading error: %v", err)
	}

	for _, library := range m.Libraries {
		parts := strings.SplitN(library.Name, ":", 3)
		url := fmt.Sprintf("%v/%v/%v/%v/%[3]v-%[4]v.jar", library.Url, strings.ReplaceAll(parts[0], ".", "/"), parts[1], parts[2])
		filePath := strings.ReplaceAll(filepath.Join("libraries", strings.ReplaceAll(parts[0], ".", "/"), parts[1], parts[2], fmt.Sprintf("%v-%v.jar", parts[1], parts[2])), " ", "_")
		if err := util.DownloadFileWithoutSignature(url, filepath.Join(path, filePath)); err != nil {
			panic(err)
		}
	}
}
