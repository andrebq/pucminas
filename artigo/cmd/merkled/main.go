// deviced simulates up to N device readings from a single device
package main

import (
	"time"

	"github.com/andrebq/pucminas/artigo/merkle"

	"github.com/Sirupsen/logrus"
	"github.com/andrebq/pucminas/artigo/flags"
)

func main() {

	flags.ParseAll()

	start := time.Now()
	var i int
	root := merkle.NewNode()
	var prevRoot *merkle.Node
	var lastKey []byte
	for {
		lastKey = []byte(time.Now().Format(time.RFC3339Nano))
		prevRoot = root
		root = root.Add(lastKey,
			[]byte(time.Now().Format(time.RFC3339Nano)))
		i++
		if time.Now().Sub(start) > flags.Duration() {
			break
		}
	}

	logrus.WithField("nodes", i).Info()

	var diffs int
	start = time.Now()
	for {
		merkle.Diff(root, prevRoot, func(_ merkle.Bytes, _, _ merkle.Bytes) (bool, error) {
			return true, nil
		})
		diffs++
		if time.Now().Sub(start) > flags.Duration() {
			break
		}
	}

	logrus.WithField("nodes", i).WithField("diffs", diffs).Info()
}
