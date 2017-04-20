package gomif

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	DefaultUri = "https://instances.mastodon.xyz/api/instances/history.json"
)

var (
	span = 3600 // sec

	ErrFailedToFetchInstances = errors.New("failed to fetch instances")
	ErrNoInstanceStatus       = errors.New(fmt.Sprintf("no instance status in last %d seconds", span))
)

type InstanceStatus struct {
	Date              int64   `json:"date"`
	Up                bool    `json:"up"`
	Users             int     `json:"users"`
	Statuses          int     `json:"statuses"`
	Connections       int     `json:"connections"`
	OpenRegistrations bool    `json:"openRegistrations"`
	Uptime            float64 `json:"uptime"`
	HttpsRank         string  `json:"https_rank"`
	HttpsScore        float64 `json:"https_score"`
	IPv6              bool    `json:"ipv6"`
}

type Client struct {
	httpClient *http.Client
	uri        string
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{},
		uri:        DefaultUri,
	}
}

func (client *Client) SetUri(uri string) {
	client.uri = uri
}

func (client *Client) FetchLastInstanceStatusByName(ctx context.Context, name string) (*InstanceStatus, error) {
	now := time.Now().Unix()

	v := url.Values{}
	v.Add("instance", name)
	v.Add("start", strconv.FormatInt(now-int64(span), 10))
	v.Add("end", strconv.FormatInt(now, 10))

	req, err := http.NewRequest("GET", client.uri, nil)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = v.Encode()
	req.WithContext(ctx)

	res, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, ErrFailedToFetchInstances
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var statuses []*InstanceStatus
	if err := json.Unmarshal(b, &statuses); err != nil {
		return nil, err
	}
	if len(statuses) == 0 {
		return nil, ErrNoInstanceStatus
	}

	return statuses[len(statuses)-1], nil
}
