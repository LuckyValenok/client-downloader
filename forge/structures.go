package forge

type Manifest struct {
	Comment       []string         `json:"_comment_"`
	Spec          int64            `json:"spec"`
	Profile       string           `json:"profile"`
	Version       string           `json:"version"`
	Path          interface{}      `json:"path"`
	Minecraft     string           `json:"minecraft"`
	ServerJarPath string           `json:"serverJarPath"`
	Data          map[string]Datum `json:"data"`
	Processors    []Processor      `json:"processors"`
	Libraries     []Library        `json:"libraries"`
	Icon          string           `json:"icon"`
	JSON          string           `json:"json"`
	Logo          string           `json:"logo"`
	MirrorList    string           `json:"mirrorList"`
	Welcome       string           `json:"welcome"`
}

type Datum struct {
	Client string `json:"client"`
	Server string `json:"server"`
}

type Library struct {
	Name      string    `json:"name"`
	Downloads Downloads `json:"downloads"`
}

type Downloads struct {
	Artifact Artifact `json:"artifact"`
}

type Artifact struct {
	Path string `json:"path"`
	URL  string `json:"url"`
	Sha1 string `json:"sha1"`
	Size int64  `json:"size"`
}

type Processor struct {
	Sides     []string `json:"sides,omitempty"`
	Jar       string   `json:"jar"`
	Classpath []string `json:"classpath"`
	Args      []string `json:"args"`
	Outputs   *Outputs `json:"outputs,omitempty"`
}

type Outputs struct {
	McSlim  string `json:"{MC_SLIM}"`
	McExtra string `json:"{MC_EXTRA}"`
}
