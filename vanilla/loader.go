package vanilla

import (
	"client-downloader/util"
	"flag"
	"log"
	"path/filepath"
)

type Settings struct {
	version string
}

func (m *Manifest) Download(path string) {
	var versionId string
	flag.StringVar(&versionId, "version", "release", "version minecraft")
	flag.Parse()

	if err := util.GetFromJson("https://launchermeta.mojang.com/mc/game/version_manifest.json", m); err != nil {
		log.Panicf("Manifest loading error: %v", err)
	}

	if versionId == "release" {
		versionId = m.Latest.Release
	} else if versionId == "snapshot" {
		versionId = m.Latest.Snapshot
	}

	var versionInfo *VersionInfo
	for _, info := range m.Versions {
		if info.Id == versionId {
			versionInfo = &info
			break
		}
	}

	if versionInfo == nil {
		log.Panic("not found version")
	}

	var version Version
	if err := util.GetFromJson(versionInfo.Url, &version); err != nil {
		log.Panicf("Version loading error: %v", err)
	}

	file, archive := util.OpenOrCreateZip(path, true)

	defer file.Close()
	defer archive.Close()

	for _, library := range version.Libraries {
		artifact := library.Downloads.Artifact
		util.DownloadAndAddToZip(artifact.Url, filepath.Join("libraries", artifact.Path), archive)
	}

	util.DownloadAndAddToZip(version.Downloads.Client.Url, "client.jar", archive)
}
