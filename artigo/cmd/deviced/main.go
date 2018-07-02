// deviced simulates up to N device readings from a single device
package main

import (
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/andrebq/pucminas/artigo/crdt"
	"github.com/andrebq/pucminas/artigo/flags"
)

func main() {

	flags.ParseAll()

	signer, err := crdt.NewSigner()
	if err != nil {
		logrus.WithError(err).Fatal("unable to start signer")
	}

	reading := crdt.Reading{
		Clock: crdt.Clock{
			Epoch: 0,
			Unix:  time.Now().UnixNano() / int64(time.Millisecond),
		},
		Stats: crdt.Stats{
			Count:   100,
			Current: 25,
			Max:     28,
			Min:     20,
			Sum:     25,
		},
		DeviceID: signer.Identity().SignerID,
	}

	start := time.Now()
	var i int
	for {
		_, err := crdt.EncodeAndSign(reading, signer)
		if err != nil {
			logrus.WithError(err).Fatal("unable to perform reading")
		}
		i++
		if time.Now().Sub(start) > flags.Duration() {
			break
		}
	}

	logrus.WithField("readings", i).Info()
}
