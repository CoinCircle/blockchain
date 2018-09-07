package blockchain

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	API_ROOT         = "https://blockchain.info"
	TESTNET_API_ROOT = "https://testnet.blockchain.info"
)

type Options struct {
	UseTestnet bool
}

type Client struct {
	*http.Client
	options *Options
}

func (c *Client) loadResponse(path string, i interface{}, formatJson bool) error {
	var apiRoot = API_ROOT
	if c.options.UseTestnet {
		apiRoot = TESTNET_API_ROOT
	}

	full_path := apiRoot + path
	if formatJson {
		full_path = apiRoot + path + "?format=json"
	}

	fmt.Println("querying..." + full_path)
	rsp, e := c.Get(full_path)
	if e != nil {
		return e
	}

	defer rsp.Body.Close()

	b, e := ioutil.ReadAll(rsp.Body)
	if e != nil {
		return e
	}
	if rsp.Status[0] != '2' {
		return fmt.Errorf("expected status 2xx, got %s: %s", rsp.Status, string(b))
	}

	return json.Unmarshal(b, &i)
}

func New(opts *Options) (*Client, error) {
	return &Client{
		Client: &http.Client{
			Timeout: time.Second * 10,
		},
		options: opts,
	}, nil
}
