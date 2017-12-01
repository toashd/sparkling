package sparkling

import "fmt"

type version struct {
	Major, Minor, Patch int    // semantic versioning scheme
	Label               string // optional build label
	Name                string // optional build name
}

// Version information. Updated manually before each release.
var Version = version{0, 1, 0, "", ""}

// Build information. Populated at build-time.
var (
	// Build is the SHA-1 (short) of the most recent commit hash.
	Build string

	// BuildDate is the date and time of the build. Should be in YYYYMMDDHHMMSS format.
	BuildDate string

	// GoVersion is the go version the binary was build with.
	GoVersion string
)

// Prints the version and build information.
func (v version) String() string {
	if v.Label != "" {
		return fmt.Sprintf("sparkling version %d.%d.%d-%s \"%s\"\nbuild: %s, %s, go%s", v.Major, v.Minor, v.Patch, v.Label, v.Name, Build, BuildDate, GoVersion)
	}
	return fmt.Sprintf("sparkling version %d.%d.%d\nbuild: %s, %s, go%s", v.Major, v.Minor, v.Patch, Build, BuildDate, GoVersion)
}
