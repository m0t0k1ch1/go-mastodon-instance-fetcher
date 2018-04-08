package gomif

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestFetchInstanceInformation(t *testing.T) {
	addedAt, _ := time.Parse(time.RFC3339, "2017-04-20T02:22:59.051Z")
	updatedAt, _ := time.Parse(time.RFC3339, "2018-04-08T08:42:34.407Z")
	checkedAt, _ := time.Parse(time.RFC3339, "2018-04-08T08:56:49.894Z")

	expected := &InstanceInformation{
		Id:                "58f81b83165fc85481cbd502",
		Name:              "mastodon.m0t0k1ch1.com",
		AddedAt:           addedAt,
		UpdatedAt:         updatedAt,
		CheckedAt:         checkedAt,
		Uptime:            0.9147309891541013,
		Up:                true,
		Dead:              false,
		Version:           "2.3.2",
		Ipv6:              true,
		HttpsScore:        100,
		HttpsRank:         "A+",
		ObsScore:          115,
		ObsRank:           "A+",
		Users:             2,
		Statuses:          725,
		Connections:       1007,
		OpenRegistrations: false,
		Thumbnail:         "https://mastodon.m0t0k1ch1.com/packs/preview-9a17d32fc48369e8ccd910a75260e67d.jpg",
		ThumbnailProxy:    "https://camo.instances.social/bb2bbd824671db963160cd1ae5d88327349fd0ee/68747470733a2f2f6d6173746f646f6e2e6d3074306b316368312e636f6d2f7061636b732f707265766965772d39613137643332666334383336396538636364393130613735323630653637642e6a7067",
		ActiveUsers:       2,
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(expected)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		return
	}))
	defer ts.Close()

	client := NewClient("")
	client.SetUri(ts.URL)

	status, err := client.FetchInstanceInformation(context.Background(), "mastodon.m0t0k1ch1.com")
	if err != nil {
		t.Fatalf("should not be fail: %v", err)
	}
	if status.Id != expected.Id {
		t.Errorf("want %q but %q", expected.Id, status.Id)
	}
	if status.Name != expected.Name {
		t.Errorf("want %q but %q", expected.Name, status.Name)
	}
	if status.AddedAt.Unix() != expected.AddedAt.Unix() {
		t.Errorf("want %v but %v", status.AddedAt, expected.AddedAt)
	}
	if status.UpdatedAt.Unix() != expected.UpdatedAt.Unix() {
		t.Errorf("want %v but %v", status.UpdatedAt, expected.UpdatedAt)
	}
	if status.CheckedAt.Unix() != expected.CheckedAt.Unix() {
		t.Errorf("want %v but %v", status.CheckedAt, expected.CheckedAt)
	}
	if status.Uptime != expected.Uptime {
		t.Errorf("want %f but %f", expected.Uptime, status.Uptime)
	}
	if status.Up != expected.Up {
		t.Errorf("want %t but %t", expected.Up, status.Up)
	}
	if status.Dead != expected.Dead {
		t.Errorf("want %t but %t", expected.Dead, status.Dead)
	}
	if status.Version != expected.Version {
		t.Errorf("want %q but %q", expected.Version, status.Version)
	}
	if status.Ipv6 != expected.Ipv6 {
		t.Errorf("want %t but %t", expected.Ipv6, status.Ipv6)
	}
	if status.HttpsScore != expected.HttpsScore {
		t.Errorf("want %d but %d", expected.HttpsScore, status.HttpsScore)
	}
	if status.HttpsRank != expected.HttpsRank {
		t.Errorf("want %q but %q", expected.HttpsRank, status.HttpsRank)
	}
	if status.ObsScore != expected.ObsScore {
		t.Errorf("want %d but %d", expected.ObsScore, status.ObsScore)
	}
	if status.ObsRank != expected.ObsRank {
		t.Errorf("want %q but %q", expected.ObsRank, status.ObsRank)
	}
	if status.Users != expected.Users {
		t.Errorf("want %d but %d", expected.Users, status.Users)
	}
	if status.Statuses != expected.Statuses {
		t.Errorf("want %d but %d", expected.Statuses, status.Statuses)
	}
	if status.Connections != expected.Connections {
		t.Errorf("want %d but %d", expected.Connections, status.Connections)
	}
	if status.OpenRegistrations != expected.OpenRegistrations {
		t.Errorf("want %t but %t", expected.OpenRegistrations, status.OpenRegistrations)
	}
	if status.Thumbnail != expected.Thumbnail {
		t.Errorf("want %q but %q", expected.Thumbnail, status.Thumbnail)
	}
	if status.ThumbnailProxy != expected.ThumbnailProxy {
		t.Errorf("want %q but %q", expected.ThumbnailProxy, status.ThumbnailProxy)
	}
	if status.ActiveUsers != expected.ActiveUsers {
		t.Errorf("want %d but %d", expected.ActiveUsers, status.ActiveUsers)
	}
}
