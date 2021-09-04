package chatwork

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
)

type Client struct {
	BaseURL    *url.URL
	HTTPClient *http.Client
	Token      string
	Logger     *log.Logger
}

func NewClient(rawURL string, token string, logger *log.Logger) (*Client, error) {
	parsedURL, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse url: %w", err)
	}

	if token == "" {
		return nil, fmt.Errorf("missing token")
	}

	if logger == nil {
		var discardLogger = log.New(ioutil.Discard, "", log.LstdFlags)
		logger = discardLogger
	}

	return &Client{
		BaseURL:    parsedURL,
		HTTPClient: http.DefaultClient,
		Token:      token,
		Logger:     logger,
	}, nil
}

func (c *Client) newRequest(ctx context.Context, method string, spath string, body io.Reader) (*http.Request, error) {
	u := *c.BaseURL
	u.Path = path.Join(c.BaseURL.Path, spath)

	req, err := http.NewRequestWithContext(ctx, method, u.String(), body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("X-ChatWorkToken", c.Token)

	return req, nil
}

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}

func (c *Client) GetAccount(ctx context.Context) (*GetAccountResponse, error) {
	spath := "/me"
	req, err := c.newRequest(ctx, http.MethodGet, spath, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	var body GetAccountResponse
	if err := decodeBody(res, &body); err != nil {
		return nil, err
	}

	return &body, nil
}

func (c *Client) GetStatus(ctx context.Context) (*GetStatusResponse, error) {
	spath := "/my/status"
	req, err := c.newRequest(ctx, http.MethodGet, spath, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	var body GetStatusResponse
	if err := decodeBody(res, &body); err != nil {
		return nil, err
	}

	return &body, nil
}

func (c *Client) ListTasks(ctx context.Context) (*ListTasksResponse, error) {
	spath := "/my/tasks"
	req, err := c.newRequest(ctx, http.MethodGet, spath, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	var body ListTasksResponse
	if err := decodeBody(res, &body); err != nil {
		return nil, err
	}

	return &body, nil
}

func (c *Client) ListContacts(ctx context.Context) (*ListContactsResponse, error) {
	spath := "/contacts"
	req, err := c.newRequest(ctx, http.MethodGet, spath, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	var body ListContactsResponse
	if err := decodeBody(res, &body); err != nil {
		return nil, err
	}

	return &body, nil
}
