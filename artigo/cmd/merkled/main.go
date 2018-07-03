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
	for {
		root = root.Add([]byte(time.Now().Format(time.RFC3339Nano)),
			[]byte(time.Now().Format(time.RFC3339Nano)))
		i++
		if time.Now().Sub(start) > flags.Duration() {
			break
		}
	}

	logrus.WithField("nodes", i).Info()
}
