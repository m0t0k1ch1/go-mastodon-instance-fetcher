package gomif

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	DefaultUri = "https://instances.social/api/1.0/instances/show"
)

type InstanceInformation struct {
	Id                string    `json:"id"`
	Name              string    `json:"name"`
	AddedAt           time.Time `json:"added_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	CheckedAt         time.Time `json:"checked_at"`
	Uptime            float64   `json:"uptime"`
	Up                bool      `json:"up"`
	Dead              bool      `json:"dead"`
	Version           string    `json:"version"`
	Ipv6              bool      `json:"ipv6"`
	HttpsScore        int       `json:"https_score"`
	HttpsRank         string    `json:"https_rank"`
	ObsScore          int       `json:"obs_score"`
	ObsRank           string    `json:"obs_rank"`
	Users             int       `json:"users"`
	Statuses          int       `json:"statuses"`
	Connections       int       `json:"connections"`
	OpenRegistrations bool      `json:"open_registrations"`
	Thumbnail         string    `json:"thumbnail"`
	ThumbnailProxy    string    `json:"thumbnail_proxy"`
	ActiveUsers       int       `json:"active_users"`
}

type ErrorResponse struct {
	Error *Error `json:"error"`
}

type Error struct {
	Message string `json:"message"`
}

type Config struct {
	Uri   string
	Token string
}

type Client struct {
	*http.Client
	config *Config
}

func NewClient(token string) *Client {
	return &Client{
		Client: http.DefaultClient,
		config: &Config{
			Uri:   DefaultUri,
			Token: token,
		},
	}
}

func (client *Client) SetUri(uri string) {
	client.config.Uri = uri
}

func (client *Client) FetchInstanceInformation(ctx context.Context, name string) (*InstanceInformation, error) {
	u, err := url.Parse(client.config.Uri)
	if err != nil {
		return nil, err
	}

	v := url.Values{}
	v.Add("name", name)

	u.RawQuery = v.Encode()

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.WithContext(ctx)
	req.Header.Set("Authorization", "Bearer "+client.config.Token)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		var resp ErrorResponse
		if err := json.Unmarshal(b, &resp); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf(resp.Error.Message)
	}

	var info InstanceInformation
	if err := json.Unmarshal(b, &info); err != nil {
		return nil, err
	}

	return &info, nil
}
