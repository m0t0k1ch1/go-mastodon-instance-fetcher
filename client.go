package gomif

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	jmespath "github.com/jmespath/go-jmespath"
)

const (
	DefaultUri     = "https://instances.mastodon.xyz/instances.json"
	DefaultTimeout = 5 // sec
)

var (
	ErrFailedToFetchInstances  = errors.New("failed to fetch instances")
	ErrFailedToExtractInstance = errors.New("failed to extract instance")
)

type Client struct {
	httpClient *http.Client
	uri        string
	timeout    time.Duration // sec
}

type Instance struct {
	Name              string  `json:"name"`
	Uptime            float64 `json:"uptime"`
	Up                bool    `json:"up"`
	HttpsScore        int     `json:"https_score"`
	HttpsRank         string  `json:"https_rank"`
	IPv6              bool    `json:"ipv6"`
	OpenRegistrations bool    `json:"openRegistrations"`
	Users             int     `json:"users"`
	Statuses          int     `json:"statuses"`
	Connections       int     `json:"connections"`
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{},
		uri:        DefaultUri,
		timeout:    DefaultTimeout,
	}
}

func (client *Client) SetUri(uri string) {
	client.uri = uri
}

func (client *Client) SetTimeout(timeout time.Duration) {
	client.timeout = timeout
}

func (client *Client) FetchInstances() ([]*Instance, error) {
	ctx, cancel := context.WithTimeout(context.Background(), client.timeout*time.Second)
	defer cancel()

	req, err := http.NewRequest("GET", client.uri, nil)
	if err != nil {
		return nil, err
	}
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

	var instances []*Instance
	if err := json.Unmarshal(b, &instances); err != nil {
		return nil, err
	}

	return instances, nil
}

func (client *Client) FetchInstanceByName(name string) (*Instance, error) {
	instances, err := client.FetchInstances()
	if err != nil {
		return nil, err
	}

	j, err := jmespath.Compile(fmt.Sprintf("[?name=='%s'] | [0]", name))
	if err != nil {
		return nil, err
	}

	result, err := j.Search(instances)
	if err != nil {
		return nil, err
	}

	instance, ok := result.(*Instance)
	if !ok {
		return nil, ErrFailedToExtractInstance
	}

	return instance, nil
}
