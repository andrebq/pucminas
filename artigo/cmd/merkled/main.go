// deviced simulates up to N device readings from a single device
package main

import (
	"time"

	"github.com/andrebq/pucminas/artigo/merkle"
	"github.com/andrebq/pucminas/artigo/stopwatch"

	"github.com/Sirupsen/logrus"
	"github.com/andrebq/pucminas/artigo/flags"
)

func main() {

	flags.ParseAll()

	var i int
	root := merkle.NewNode()
	var prevRoot *merkle.Node
	var lastKey []byte
	logrus.Info("starting")
	s := stopwatch.New(flags.Duration())
	for !s.Stop() {
		lastKey = []byte(time.Now().Format(time.RFC3339Nano))
		prevRoot = root
		root = root.Add(lastKey,
			[]byte(time.Now().Format(time.RFC3339Nano)))
		i++
	}

	logrus.WithField("nodes", i).Info()

	var diffs int
	s = stopwatch.New(flags.Duration())
	logrus.Info("second_loop")
	for !s.Stop() {
		merkle.Diff(root, prevRoot, func(_ merkle.Bytes, _, _ merkle.Bytes) (bool, error) {
			return true, nil
		})
		diffs++
	}

	logrus.WithField("nodes", i).WithField("diffs", diffs).Info()
}
