package main

import (
	"context"
	"os"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/andrebq/pucminas/artigo/ctrlc"
	"github.com/andrebq/pucminas/artigo/extract"
)

func main() {
	ctx, cancel := ctrlc.Watch(context.Background())
	defer cancel()

	input := extract.NewContextRead(ctx, time.Second, os.Stdin)
	defer input.Close()

	pipe := extract.NewPipe(os.Stdout, input)

	pipe.WriteHeader()
	for {
		err := pipe.Copy()
		if extract.IsTimeout(err) {
			logrus.WithError(err).Fatal("aborting")
		} else if extract.IsCancel(err) {
			logrus.WithError(err).Fatal("cancel")
		}
		if err != nil {
			logrus.WithError(err).Error("error copying data")
		}
	}
}
