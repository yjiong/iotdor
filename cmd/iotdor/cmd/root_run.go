package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var ctx context.Context
var cancel context.CancelFunc

func run(cmd *cobra.Command, args []string) error {
	ctx = context.Background()
	ctx, cancel = context.WithCancel(ctx)
	defer cancel()

	tasks := []func() error{
		setLogLevel,
		printStartMessage,
		getDataSrc,
		//setDataProcess,
		//setHjClient,
		//startAPI,
		//chkReInit,
		//startGorilla,
	}
	for _, t := range tasks {
		if err := t(); err != nil {
			log.Fatal(err)
		}
	}
	sigChan := make(chan os.Signal, 1)
	exitChan := make(chan struct{}, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	log.WithField("signal", <-sigChan).Info("signal received")
	go func() {
		log.Warning("stopping iotdor process")
		exitChan <- struct{}{}
	}()
	select {
	case <-exitChan:
	case s := <-sigChan:
		log.WithField("signal", s).Info("signal received, stopping immediately")
	}

	return nil
}
