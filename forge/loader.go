package forge

import (
	"archive/zip"
	"client-downloader/util"
	"client-downloader/vanilla"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"path/filepath"
)

func (m *Manifest) Download(path string) {
	var vanillaManifest vanilla.Manifest
	vanillaManifest.Download(path)

	forgeVersion := flag.Lookup("forge-version").Value.String()
	versionId := flag.Lookup("version").Value.String()

	url := fmt.Sprintf("https://maven.minecraftforge.net/net/minecraftforge/forge/%[1]v-%[2]v/forge-%[1]v-%[2]v-installer.jar", versionId, forgeVersion)

	forgeInstaller := filepath.Join(path, "forge-installer.jar")
	err := util.DownloadFile(url, forgeInstaller, false)
	if err != nil {
		panic(err)
	}

	r, err := zip.OpenReader(forgeInstaller)
	if err != nil {
		panic(err)
	}
	defer r.Close()

	for _, f := range r.File {
		if f.Name == "install_profile.json" {
			rc, err := f.Open()
			if err != nil {
				panic(err)
			}
			defer rc.Close()

			if err := json.NewDecoder(rc).Decode(m); err != nil {
				panic(err)
			}

			break
		}
	}

	log.Printf("%v", m)
}
