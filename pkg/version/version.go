package version

import (
	"fmt"
	"runtime"
)

type Info struct {
	GitTag      string `json:"gitTag"`
	GitCommit   string `json:"gitCommit"`
	GitTreeSate string `json:"gitTreeState"`
	BuildDate   string `json:"buildDate"`
	GoVersion   string `json:"goVersion"`
	Compiler    string `json:"compiler"`
	Platform    string `json:"platform"`
}

func Get() Info {
	return Info{
		GitTag:      gitTag,
		GitCommit:   gitCommit,
		GitTreeSate: gitTreeState,
		BuildDate:   buildDate,
		GoVersion:   runtime.Version(),
		Compiler:    runtime.Compiler,
		Platform:    fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
