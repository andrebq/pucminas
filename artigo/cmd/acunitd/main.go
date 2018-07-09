package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/andrebq/pucminas/artigo/crdt"
	"github.com/andrebq/pucminas/artigo/flags"
	"github.com/andrebq/pucminas/artigo/stopwatch"
	"github.com/rs/xid"
)

var (
	counter uint64
)

func main() {
	flags.ParseAll()

	nodes := make([]crdt.Device, flags.Nodes())

	for i := range nodes {
		nodes[i] = crdt.Device{
			ID:        xid.New().String(),
			LastClock: crdt.NewClock(0),
		}
	}

	acunit := crdt.ACUnit{
		ID: xid.New().String(),
	}
	s := stopwatch.New(flags.Duration())
	logrus.Info("starting")
	count := 0
	for !s.Stop() {
		for _, n := range nodes {
			acunit = acunit.Update(randomControlMessage(acunit.ID, n))
		}
		count++
	}

	logrus.WithField("iterations", count).Info()

	count = 0
	s = stopwatch.New(flags.Duration())
	logrus.Info("second_loop")
	for !s.Stop() {
		status := acunit.Query()
		count++
		_ = status
	}
	logrus.WithField("queries", count).Info()
}

func randomControlMessage(acID string, n crdt.Device) crdt.ControlMessage {
	n.LastClock = n.LastClock.Tick()
	cm := crdt.ControlMessage{
		ControlUnitID: n.ID,
		Status:        counter%2 == 0,
		Clock:         n.LastClock,
		ACUnitID:      acID,
	}
	counter++
	return cm
}
