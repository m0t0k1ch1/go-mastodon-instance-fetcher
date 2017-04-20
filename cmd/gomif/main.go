package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/m0t0k1ch1/gomif"
)

const (
	DefaultSpan    = 3600
	DefaultTimeout = 5 // sec
)

func main() {
	os.Exit(run())
}

func run() int {
	var (
		instance string
		span     int
		timeout  int
	)

	flag.StringVar(&instance, "instance", "", "instance name")
	flag.StringVar(&instance, "i", "", "instance name")
	flag.IntVar(&span, "span", DefaultSpan, "time period for fetching instance status")
	flag.IntVar(&timeout, "timeout", DefaultTimeout, "HTTP request timeout (sec)")
	flag.Parse()

	if len(instance) == 0 {
		fmt.Println("instance name is not specified")
		return 1
	}

	client := gomif.NewClient()

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)
	defer cancel()

	statusCh := make(chan *gomif.InstanceStatus, 1)
	errCh := make(chan error, 1)

	go func() {
		status, err := client.FetchLastInstanceStatus(ctx, instance, int64(span))
		if err != nil {
			errCh <- err
			return
		}
		statusCh <- status
	}()

	select {
	case status := <-statusCh:
		b, err := json.Marshal(status)
		if err != nil {
			fmt.Println(err)
			return 1
		}
		fmt.Println(string(b))
		return 0
	case err := <-errCh:
		fmt.Println(err)
		return 1
	case <-ctx.Done():
		fmt.Println(ctx.Err())
		return 1
	}
}
