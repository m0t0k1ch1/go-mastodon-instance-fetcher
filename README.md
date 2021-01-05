# gomif

[![GoDoc](https://godoc.org/github.com/m0t0k1ch1/gomif?status.svg)](https://godoc.org/github.com/m0t0k1ch1/gomif) [![wercker status](https://app.wercker.com/status/3ed695cdfed9a63dd66b823302604041/s/master "wercker status")](https://app.wercker.com/project/byKey/3ed695cdfed9a63dd66b823302604041)

a [Mastodon](https://github.com/tootsuite/mastodon) instance information fetcher for Go

## Examples

First you need to [get a instances.social API token](https://instances.social/api/token).

### Use as CLI

``` sh
$ go get -u github.com/m0t0k1ch1/gomif/cmd/gomif
$ GOMIF_TOKEN=<your token> gomif -i mastodon.m0t0k1ch1.com
```

### Use in code

``` go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/m0t0k1ch1/gomif"
)

func main() {
	client := gomif.NewClient("your token")

	info, err := client.FetchInstanceInformation(
		context.Background(),
		"mastodon.m0t0k1ch1.com",
	)
	if err != nil {
		log.Fatal(err)
	}

	b, err := json.Marshal(info)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
}
```
