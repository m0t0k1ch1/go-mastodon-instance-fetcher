# gomif

[![GoDoc](https://godoc.org/github.com/m0t0k1ch1/gomif?status.svg)](https://godoc.org/github.com/m0t0k1ch1/gomif) [![wercker status](https://app.wercker.com/status/3ed695cdfed9a63dd66b823302604041/s/master "wercker status")](https://app.wercker.com/project/byKey/3ed695cdfed9a63dd66b823302604041)

GO Mastodon Instance Fetcher

a package for fetching [Mastodon](https://github.com/tootsuite/mastodon) instance information

## Example

``` go
package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/m0t0k1ch1/gomif"
)

func main() {
	client := gomif.NewClient()

	instance, err := client.FetchInstanceByName("mastodon.m0t0k1ch1.com")
	if err != nil {
		log.Fatal(err)
	}

	instanceBytes, err := json.Marshal(instance)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(instanceBytes))
}
```
