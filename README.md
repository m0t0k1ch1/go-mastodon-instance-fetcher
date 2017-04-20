# gomif

[![GoDoc](https://godoc.org/github.com/m0t0k1ch1/gomif?status.svg)](https://godoc.org/github.com/m0t0k1ch1/gomif) [![wercker status](https://app.wercker.com/status/3ed695cdfed9a63dd66b823302604041/s/master "wercker status")](https://app.wercker.com/project/byKey/3ed695cdfed9a63dd66b823302604041)

GO Mastodon Instance Fetcher

a package for fetching [Mastodon](https://github.com/tootsuite/mastodon) instance information

## Example

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
	client := gomif.NewClient()

	ctx := context.Background()

	instance, err := client.FetchInstanceByName(ctx, "mastodon.m0t0k1ch1.com")
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

``` json
{
  "name": "mastodon.m0t0k1ch1.com",
  "uptime": 99.40828402366864,
  "up": true,
  "https_score": 100,
  "https_rank": "A+",
  "ipv6": true,
  "openRegistrations": false,
  "users": 1,
  "statuses": 13,
  "connections": 14
}
```
