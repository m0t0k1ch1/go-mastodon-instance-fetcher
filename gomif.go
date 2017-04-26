package gomif

import (
	"context"
	"encoding/json"
	"errors"
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
	ErrInstanceNotFound            = errors.New("instance not found")
	ErrFailedToFetchInstanceStatus = errors.New("failed to fetch instanceStatus")
	ErrNoInstanceStatus            = errors.New("no instance status")
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

type Config struct {
	Uri *url.URL
}

type Client struct {
	*http.Client
	config *Config
}

func NewClient() *Client {
	u, _ := url.Parse(DefaultUri)

	return &Client{
		Client: http.DefaultClient,
		config: &Config{
			Uri: u,
		},
	}
}

func (client *Client) SetUri(uri string) error {
	u, err := url.Parse(uri)
	if err != nil {
		return err
	}

	client.config.Uri = u

	return nil
}

func (client *Client) FetchInstanceStatuses(ctx context.Context, name string, start, end int64) ([]*InstanceStatus, error) {
	v := url.Values{}
	v.Add("instance", name)
	v.Add("start", strconv.FormatInt(start, 10))
	v.Add("end", strconv.FormatInt(end, 10))

	client.config.Uri.RawQuery = v.Encode()

	req, err := http.NewRequest(http.MethodGet, client.config.Uri.String(), nil)
	if err != nil {
		return nil, err
	}
	req.WithContext(ctx)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusNotFound {
			return nil, ErrInstanceNotFound
		}
		return nil, ErrFailedToFetchInstanceStatus
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var statuses []*InstanceStatus
	if err := json.Unmarshal(b, &statuses); err != nil {
		return nil, err
	}

	return statuses, nil
}

func (client *Client) FetchLastInstanceStatus(ctx context.Context, name string, span int64) (*InstanceStatus, error) {
	now := time.Now().Unix()

	statuses, err := client.FetchInstanceStatuses(ctx, name, now-span, now)
	if err != nil {
		return nil, err
	}
	if len(statuses) == 0 {
		return nil, ErrNoInstanceStatus
	}

	return statuses[len(statuses)-1], nil
}
