package gomif

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchLastInstanceStatus(t *testing.T) {
	expected := &InstanceStatus{
		Date:              1492686962,
		Up:                true,
		Users:             1,
		Statuses:          14,
		Connections:       14,
		OpenRegistrations: false,
		Uptime:            0.9942748091603053,
		HttpsRank:         "A+",
		HttpsScore:        100,
		IPv6:              true,
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal([]*InstanceStatus{expected})
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

	client := NewClient()
	client.SetUri(ts.URL)

	status, err := client.FetchLastInstanceStatus(context.Background(), "mastodon.m0t0k1ch1.com", 3600)
	if err != nil {
		t.Fatalf("should not be fail: %v", err)
	}
	if status.Date != expected.Date {
		t.Errorf("want %d but %d", expected.Date, status.Date)
	}
	if status.Up != expected.Up {
		t.Errorf("want %t but %t", expected.Up, status.Up)
	}
	if status.Users != expected.Users {
		t.Errorf("want %d but %d", expected.Users, status.Users)
	}
	if status.Statuses != expected.Statuses {
		t.Errorf("want %d but %d", expected.Users, status.Users)
	}
	if status.Connections != expected.Connections {
		t.Errorf("want %d but %d", expected.Connections, status.Connections)
	}
	if status.OpenRegistrations != expected.OpenRegistrations {
		t.Errorf("want %t but %t", expected.OpenRegistrations, status.OpenRegistrations)
	}
	if status.Uptime != expected.Uptime {
		t.Errorf("want %f but %f", expected.Uptime, status.Uptime)
	}
	if status.HttpsRank != expected.HttpsRank {
		t.Errorf("want %q but %q", expected.HttpsRank, status.HttpsRank)
	}
	if status.HttpsScore != expected.HttpsScore {
		t.Errorf("want %d but %d", expected.HttpsScore, status.HttpsScore)
	}
	if status.IPv6 != expected.IPv6 {
		t.Errorf("want %t but %t", expected.IPv6, status.IPv6)
	}
}
