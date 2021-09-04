package chatwork

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
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
