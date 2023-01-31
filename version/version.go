package version

import (
	"fmt"
)

const (
	// ServiceName 服务名称
	ServiceName = "demo"
)

var (
	GitTag    string = "v1.0.0"
	GitCommit string
	GitBranch string
	BuildTime string
	GoVersion string
)

// FullVersion show the version info
func FullVersion() string {
	version := fmt.Sprintf("Version   : %s\nBuild Time: %s\nGit Branch: %s\nGit Commit: %s\nGo Version: %s\n", GitTag, BuildTime, GitBranch, GitCommit, GoVersion)
	return version
}

// Short 版本缩写
func Short() string {
	return fmt.Sprintf("%s[%s %s]", GitTag, BuildTime, GitCommit)
}
