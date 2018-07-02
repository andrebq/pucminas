package flags

import (
	"flag"
	"os"
	"time"
)

var (
	duration = flag.Duration("duration", time.Second, "How much time to spend doing simulated readins")
	help     = flag.Bool("h", false, "Show help")
)

// ParseAll parses all registred flags, if help is set,
// the system will print usage and exit with -1
func ParseAll() {
	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(1)
	}
}

// Duration of the experiment
func Duration() time.Duration {
	return *duration
}
