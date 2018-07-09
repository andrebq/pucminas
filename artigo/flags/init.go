package flags

import (
	"flag"
	"os"
	"time"

	"github.com/Sirupsen/logrus"
)

var (
	duration      = flag.Duration("duration", time.Second, "How much time to spend doing simulated readins")
	nodes         = flag.Uint("nodes", 1, "How many nodes should generate data (not valid for all simulations")
	fulltimestamp = flag.Bool("fullts", true, "Use a full timestamp for logrus")
	jsonformat    = flag.Bool("json", false, "Use JSONFormatter to output logging info")
	help          = flag.Bool("h", false, "Show help")
)

// ParseAll parses all registred flags, if help is set,
// the system will print usage and exit with -1
func ParseAll() {
	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(1)
	}

	if *fulltimestamp {
		fmter := &logrus.TextFormatter{}
		fmter.FullTimestamp = *fulltimestamp

		logrus.SetFormatter(fmter)
	}
}

// Duration of the experiment
func Duration() time.Duration {
	return *duration
}

// Nodes returns how many different nodes are writing data
func Nodes() int {
	return int(*nodes)
}
