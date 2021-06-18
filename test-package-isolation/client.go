package mypkg

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

var baseURL = "https://example.com/api/v2"

type Client struct {
	// fields
	HTTPClient *http.Client
}

func (cli *Client) httpClient() *http.Client {
	if cli.HTTPClient != nil {
		return cli.HTTPClient
	}
	return http.DefaultClient
}

type getResponse struct {
	Value string `json:"value"`
}

func (cli *Client) Get(n int) (string, error) {
	v := url.Values{}
	v.Set("n", strconv.Itoa(n))
	requestURL := baseURL + "/get?" + v.Encode()
	resp, err := cli.httpClient().Get(requestURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var gr getResponse
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&gr); err != nil {
		return "", err
	}
	return gr.Value, nil
}
