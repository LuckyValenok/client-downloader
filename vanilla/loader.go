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
	versionId := flag.Lookup("version").Value.String()

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

	for _, library := range version.Libraries {
		forWindows := len(library.Rules) == 0
		for _, rule := range library.Rules {
			if rule.Os.Name == "windows" {
				forWindows = true
			}
		}
		if !forWindows {
			continue
		}
		artifact := library.Downloads.Artifact
		if err := util.DownloadFileWithoutSignature(artifact.Url, filepath.Join(path, "libraries", artifact.Path)); err != nil {
			panic(err)
		}
	}

	if err := util.DownloadFileWithoutSignature(version.Downloads.Client.Url, filepath.Join(path, "client.jar")); err != nil {
		panic(err)
	}
}
