package gomif

import (
	"context"
	"testing"
)

func TestFetchLastInstanceStatus(t *testing.T) {
	client := NewClient()

	_, err := client.FetchLastInstanceStatus(context.Background(), "mastodon.m0t0k1ch1.com", 3600)
	if err != nil {
		t.Fatal("failed to fetch instance status")
	}
}
