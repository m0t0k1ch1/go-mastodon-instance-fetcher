package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/m0t0k1ch1/gomif"
)

const (
	GomifTokenKey = "GOMIF_TOKEN"
)

func main() {
	os.Exit(run())
}

func run() int {
	var instance string

	flag.StringVar(&instance, "instance", "", "instance name")
	flag.StringVar(&instance, "i", "", "instance name")
	flag.Parse()

	if len(instance) == 0 {
		return fail(fmt.Errorf("instance name is not specified"))
	}

	client := gomif.NewClient(os.Getenv(GomifTokenKey))

	info, err := client.FetchInstanceInformation(context.Background(), instance)
	if err != nil {
		return fail(err)
	}

	b, err := json.Marshal(info)
	if err != nil {
		return fail(err)
	}
	fmt.Println(string(b))

	return success()
}

func success() int {
	return 0
}

func fail(err error) int {
	fmt.Println(err)
	return 1
}
