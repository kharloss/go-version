package build

import (
	"fmt"
	"runtime"
	"time"
)

var (
	// Application supplied by the linker
	Application = "goexecutable"
	// BuildDate supplied by the linker
	BuildDate = fmt.Sprintf("%s", time.Now().UTC().Format("2006-01-02T15:04:05+00:00"))
	// Version supplied by the linker
	Version = "v0.0.0"
	// Revision supplied by the linker
	Revision = "00000000"
	// GoVersion supplied by the runtime
	GoVersion = runtime.Version()
)

// init returns the version
func init() {
	fmt.Printf("%s version %s git revision %s go version %s build date %s\n", Application, Version, Revision, GoVersion, BuildDate)
}
