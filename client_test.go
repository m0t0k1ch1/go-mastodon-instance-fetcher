package gomif

import (
	"context"
	"testing"
)

func TestFetchInstanceByName(t *testing.T) {
	client := NewClient()

	_, err := client.FetchInstanceByName(context.Background(), "mastodon.m0t0k1ch1.com")
	if err != nil {
		t.Fatal("failed to fetch instance")
	}
}
