# gomif

[![GoDoc](https://godoc.org/github.com/m0t0k1ch1/gomif?status.svg)](https://godoc.org/github.com/m0t0k1ch1/gomif) [![wercker status](https://app.wercker.com/status/3ed695cdfed9a63dd66b823302604041/s/master "wercker status")](https://app.wercker.com/project/byKey/3ed695cdfed9a63dd66b823302604041)

GO [Mastodon](https://github.com/tootsuite/mastodon) Instance status Fetcher

## Examples

### Use as CLI

``` sh
$ go get -u github.com/m0t0k1ch1/gomif/cmd/gomif
$ gomif -i mastodon.m0t0k1ch1.com | jq .
{
  "date": 1492689362,
  "up": true,
  "users": 1,
  "statuses": 15,
  "connections": 17,
  "openRegistrations": false,
  "uptime": 0.9946714031971581,
  "https_rank": "A+",
  "https_score": 100,
  "ipv6": true
}
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
	client := gomif.NewClient()

	status, err := client.FetchLastInstanceStatus(
		context.Background(),
		"mastodon.m0t0k1ch1.com",
		3600, // span (sec)
	)
	if err != nil {
		log.Fatal(err)
	}

	b, err := json.Marshal(status)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
}
```

``` json
{
  "date": 1492689362,
  "up": true,
  "users": 1,
  "statuses": 15,
  "connections": 17,
  "openRegistrations": false,
  "uptime": 0.9946714031971581,
  "https_rank": "A+",
  "https_score": 100,
  "ipv6": true
}
```
