# gomif

[![GoDoc](https://godoc.org/github.com/m0t0k1ch1/gomif?status.svg)](https://godoc.org/github.com/m0t0k1ch1/gomif) [![wercker status](https://app.wercker.com/status/3ed695cdfed9a63dd66b823302604041/s/master "wercker status")](https://app.wercker.com/project/byKey/3ed695cdfed9a63dd66b823302604041)

GO Mastodon Instance status Fetcher

a package for fetching [Mastodon](https://github.com/tootsuite/mastodon) instance status

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

	instance, err := client.FetchLastInstanceStatusByName(ctx, "mastodon.m0t0k1ch1.com")
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
  "date": 1492688462,
  "up": true,
  "users": 1,
  "statuses": 15,
  "connections": 15,
  "openRegistrations": false,
  "uptime": 0.994535519125683,
  "https_rank": "A+",
  "https_score": 100,
  "ipv6": true
}
```
