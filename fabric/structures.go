package fabric

type Manifest struct {
	Id           string `json:"id"`
	InheritsFrom string `json:"inheritsFrom"`
	ReleaseTime  string `json:"releaseTime"`
	Time         string `json:"time"`
	Type         string `json:"type"`
	MainClass    string `json:"mainClass"`
	Arguments    struct {
		Game []interface{} `json:"game"`
		Jvm  []string      `json:"jvm"`
	} `json:"arguments"`
	Libraries []Library `json:"libraries"`
}

type Library struct {
	Name   string `json:"name"`
	Url    string `json:"url"`
	Md5    string `json:"md5,omitempty"`
	Sha1   string `json:"sha1,omitempty"`
	Sha256 string `json:"sha256,omitempty"`
	Sha512 string `json:"sha512,omitempty"`
	Size   int    `json:"size,omitempty"`
}
