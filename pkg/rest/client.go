package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
)

type Client struct {
	client *http.Client
	host   string
	port   int
	token  string
	ssl    bool
}

func New(host string, port int, token string, useSSL bool) *Client {
	return &Client{
		client: &http.Client{},
		host:   host,
		port:   port,
		token:  token,
		ssl:    useSSL,
	}
}

func Get[R any](c *Client, p string, values url.Values) (*R, error) {
	scheme := "http"
	if c.ssl {
		scheme = "https"
	}

	addPath := ""
	if p == "" {
		addPath = "/"
	}
	u := url.URL{
		Scheme: scheme,
		Host:   fmt.Sprintf("%s:%d", c.host, c.port),
		Path:   path.Join("api", p) + addPath,
	}

	u.RawQuery = values.Encode()

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	// add authorization header to the req
	req.Header.Add("Authorization", "Bearer "+c.token)
	req.Header.Add("content-type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	var rtn R

	err = json.NewDecoder(resp.Body).Decode(&rtn)
	return &rtn, err
}
