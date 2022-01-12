package http

import (
	"fmt"
	"github.com/kelseyhightower/confd/log"
	"io/ioutil"
	"net/http"
	"strings"
)

// Client provides a shell for the yaml client
type Client struct {
	url string
}

func NewHttpClient(url string) (*Client, error) {
	return &Client{url: strings.TrimSuffix(url, "/") + "/"}, nil
}
func (c *Client) GetValues(keys []string) (map[string]string, error) {
	client := http.DefaultClient
	vars := make(map[string]string)
	for idx, key := range keys {
		url := c.url + strings.TrimSuffix(strings.TrimPrefix(key, "/"), "*")
		log.Debug(fmt.Sprintf("key -- :[%d], %s,  %s", idx, key, url))
		value, err := c.getValue(client, url)
		if err != nil {
			return nil, err
		}
		vars[key] = value
	}
	log.Debug(fmt.Sprintf("Key Map: %#v", vars))
	return vars, nil
}

func (c *Client) getValue(client *http.Client, url string) (string, error) {
	response, err := client.Get(url)
	//goland:noinspection GoUnhandledErrorResult
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
func (c *Client) WatchPrefix(prefix string, keys []string, waitIndex uint64, stopChan chan bool) (uint64, error) {
	<-stopChan
	return 0, nil
}
