package gomif

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
