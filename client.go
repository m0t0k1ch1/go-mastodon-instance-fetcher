package gomif

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	jmespath "github.com/jmespath/go-jmespath"
)

const (
	DefaultUri = "https://instances.mastodon.xyz/instances.json"
)

var (
	ErrFailedToFetchInstances  = errors.New("failed to fetch instances")
	ErrFailedToExtractInstance = errors.New("failed to extract instance")
)

type Client struct {
	httpClient *http.Client
	uri        string
}

type Instance struct {
	Name              string  `json:"name"`
	Uptime            float64 `json:"uptime"`
	Up                bool    `json:"up"`
	HttpsScore        float64 `json:"https_score"`
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
	}
}

func (client *Client) SetUri(uri string) {
	client.uri = uri
}

func (client *Client) FetchInstances(ctx context.Context) ([]*Instance, error) {
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

func (client *Client) FetchInstanceByName(ctx context.Context, name string) (*Instance, error) {
	instances, err := client.FetchInstances(ctx)
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
