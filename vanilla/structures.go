package vanilla

import "time"

type VersionInfo struct {
	Id          string    `json:"id"`
	Type        string    `json:"type"`
	Url         string    `json:"url"`
	Time        time.Time `json:"time"`
	ReleaseTime time.Time `json:"releaseTime"`
}

type Manifest struct {
	Latest struct {
		Release  string `json:"release"`
		Snapshot string `json:"snapshot"`
	} `json:"latest"`
	Versions []VersionInfo `json:"versions"`
}

type Version struct {
	Arguments struct {
		Game []interface{} `json:"game"`
		Jvm  []interface{} `json:"jvm"`
	} `json:"arguments"`

	AssetIndex struct {
		Id        string `json:"id"`
		Sha1      string `json:"sha1"`
		Size      int    `json:"size"`
		TotalSize int    `json:"totalSize"`
		Url       string `json:"url"`
	} `json:"assetIndex"`

	Assets string `json:"assets"`

	ComplianceLevel int `json:"complianceLevel"`

	Downloads struct {
		Client         File `json:"client"`
		ClientMappings File `json:"client_mappings"`
		Server         File `json:"server"`
		ServerMappings File `json:"server_mappings"`
	} `json:"downloads"`
	Id          string `json:"id"`
	JavaVersion struct {
		Component    string `json:"component"`
		MajorVersion int    `json:"majorVersion"`
	} `json:"javaVersion"`
	Libraries []struct {
		Downloads struct {
			Artifact Artifact `json:"artifact"`
		} `json:"downloads"`
		Name  string `json:"name"`
		Rules []Rule `json:"rules,omitempty"`
	} `json:"libraries"`
	Logging struct {
		Client struct {
			Argument string     `json:"argument"`
			File     FileWithId `json:"file"`
			Type     string     `json:"type"`
		} `json:"client"`
	} `json:"logging"`
	MainClass              string    `json:"mainClass"`
	MinimumLauncherVersion int       `json:"minimumLauncherVersion"`
	ReleaseTime            time.Time `json:"releaseTime"`
	Time                   time.Time `json:"time"`
	Type                   string    `json:"type"`
}

type File struct {
	Sha1 string `json:"sha1"`
	Size int    `json:"size"`
	Url  string `json:"url"`
}

type Artifact struct {
	File
	Path string `json:"path"`
}

type FileWithId struct {
	File
	Id string `json:"id"`
}

type Rule struct {
	Action string `json:"action"`
	Os     struct {
		Name string `json:"name"`
	} `json:"os"`
}

type Asset struct {
	Hash string `json:"hash"`
	Size int    `json:"size"`
}

type JSONData struct {
	Objects map[string]Asset `json:"objects"`
}
