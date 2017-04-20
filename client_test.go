package gomif

import (
	"context"
	"testing"
)

func TestFetchInstanceByName(t *testing.T) {
	name := "mastodon.m0t0k1ch1.com"

	client := NewClient()
	instance, err := client.FetchInstanceByName(context.Background(), name)
	if err != nil {
		t.Fatal("failed to fetch instance")
	}

	if instance.Name != name {
		t.Errorf("unexpected instance name")
	}
}
