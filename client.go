package chatwork

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
)

const BaseURLV2 = "https://api.chatwork.com/v2"

type Client struct {
	URL        *url.URL
	HTTPClient *http.Client
	Token      string
	Logger     *log.Logger
}

func NewClient(urlStr string, token string, logger *log.Logger) (*Client, error) {
	if len(token) == 0 {
		return nil, errors.New("missing token")
	}

	parsedURL, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return nil, err
	}

	var discardLogger = log.New(ioutil.Discard, "", log.LstdFlags)
	if logger == nil {
		logger = discardLogger
	}

	return &Client{
		URL:        parsedURL,
		HTTPClient: http.DefaultClient,
		Token:      token,
		Logger:     logger,
	}, nil
}

func (c *Client) newRequest(ctx context.Context, method string, spath string, body io.Reader) (*http.Request, error) {
	u := *c.URL
	u.Path = path.Join(c.URL.Path, spath)

	req, err := http.NewRequestWithContext(ctx, method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("X-ChatWorkToken", c.Token)

	return req, nil
}

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}
