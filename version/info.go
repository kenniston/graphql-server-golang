package version

import "fmt"

// Default build-time variable.
// These values are overridden via ldflags
var (
	Version   = "0.0.1"
	GitCommit = "unknown-commit"
	BuildTime = "unknown-buildtime"
)

const versionF = `graphql-server
  Version: %s
  GitCommit: %s
  BuildTime: %s
`

// FormattedMessage gets the full formatted version message
func FormattedMessage() string {
	return fmt.Sprintf(versionF, Version, GitCommit, BuildTime)
}