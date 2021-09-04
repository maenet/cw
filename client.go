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

func (c *Client) newRequest(ctx context.Context, method string, pathstr string, body io.Reader) (*http.Request, error) {
	u := *c.BaseURL
	u.Path = path.Join(c.BaseURL.Path, pathstr)

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

func (c *Client) callAPI(ctx context.Context, method string, pathstr string, body io.Reader, responseBody interface{}) error {
	req, err := c.newRequest(ctx, http.MethodGet, pathstr, body)
	if err != nil {
		return err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	if err := decodeBody(res, &responseBody); err != nil {
		return err
	}

	return nil
}

func (c *Client) GetAccount(ctx context.Context) (*GetAccountResponse, error) {
	var res GetAccountResponse
	if err := c.callAPI(ctx, http.MethodGet, "/me", nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) GetStatus(ctx context.Context) (*GetStatusResponse, error) {
	var res GetStatusResponse
	if err := c.callAPI(ctx, http.MethodGet, "/my/status", nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) ListMyTasks(ctx context.Context) (*ListMyTasksResponse, error) {
	var res ListMyTasksResponse
	if err := c.callAPI(ctx, http.MethodGet, "/my/tasks", nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) ListContacts(ctx context.Context) (*ListContactsResponse, error) {
	var res ListContactsResponse
	if err := c.callAPI(ctx, http.MethodGet, "/contacts", nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// func (c *Client) ListRooms(ctx context.Context) (*ListContactsResponse, error) { return nil, nil }
// func (c *Client) CreateRoom(ctx context.Context, name string, desc string, icon string, adminIDs []int, memberIDs []int, readonlyIDs []int, link bool, linkCode string, linkNeedCcceptance bool) (*ListContactsResponse, error) { return nil, nil }
// func (c *Client) GetRoom(ctx context.Context, roomID int) (*ListContactsResponse, error) { return nil, nil }
// func (c *Client) UpdateRoom(ctx context.Context, roomID int, name string, desc string, icon string) (*ListContactsResponse, error) { return nil, nil }
// func (c *Client) DeleteRoom(ctx context.Context, roomID int, actionType string) (*ListContactsResponse, error) { return nil, nil }
//
// func (c *Client) ListMembers(ctx context.Context, roomID int) (*ListContactsResponse, error) { return nil, nil }
// func (c *Client) UpdateMember(ctx context.Context, roomID int, adminIDs []int, memberIDs []int, readonlyIDs []int) (*ListContactsResponse, error) { return nil, nil }
//
// func (c *Client) ListMessages(ctx context.Context, roomID int, force bool) (*ListContactsResponse, error) { return nil, nil }
// func (c *Client) CreateMessage(ctx context.Context, roomID int, body string, selfUnread bool) (*ListContactsResponse, error) { return nil, nil }
// func (c *Client) MarkMessagesAsRead(ctx context.Context, roomID int, messageID string) (*ListContactsResponse, error) { return nil, nil }
// func (c *Client) MarkMessagesAsUnread(ctx context.Context, roomID int, messageID string) (*ListContactsResponse, error) { return nil, nil }
// func (c *Client) GetMessage(ctx context.Context, roomID int, messageID string) (*ListContactsResponse, error) { return nil, nil }
// func (c *Client) UpdateMessage(ctx context.Context, roomID int, messageID string, body string) (*ListContactsResponse, error) { return nil, nil }
// func (c *Client) DeleteMessage(ctx context.Context, roomID int, messageID string) (*ListContactsResponse, error) { return nil, nil }
//
// func (c *Client) ListTasks(ctx context.Context, roomID int, accountID int, assignorAccountID int, status string) (*ListContactsResponse, error) { return nil, nil }
// func (c *Client) CreateTask(ctx context.Context, roomID int, body string, limit int, limitType string, toIds []int) (*ListContactsResponse, error) { return nil, nil }
// func (c *Client) GetTask(ctx context.Context, roomID int, taskID int) (*ListContactsResponse, error) { return nil, nil }
// func (c *Client) UpdateTaskStatus(ctx context.Context, roomID int, taskID int, body string) (*ListContactsResponse, error) { return nil, nil }
//
// func (c *Client) ListFiles(ctx context.Context, roomID int, accountID int) (*ListContactsResponse, error) { return nil, nil }
// func (c *Client) UploadFile(ctx context.Context, roomID int, fpath string, message string) (*ListContactsResponse, error) { return nil, nil }
// func (c *Client) GetFile(ctx context.Context, roomID int, fileId int) (*ListContactsResponse, error) { return nil, nil }
//
// func (c *Client) GetLink(ctx context.Context, roomID int) (*ListContactsResponse, error) { return nil, nil }
// func (c *Client) CreateLink(ctx context.Context, roomID int, code string, desc string, need_acceptance bool) (*ListContactsResponse, error) { return nil, nil }
// func (c *Client) UpdateLink(ctx context.Context, roomID int, code string, desc string, need_acceptance bool) (*ListContactsResponse, error) { return nil, nil }
// func (c *Client) DeleteLink(ctx context.Context, roomID int) (*ListContactsResponse, error) { return nil, nil }

// func (c *Client) ListIncomingRequests(ctx context.Context) (*ListContactsResponse, error) { return nil, nil }
// func (c *Client) AcceptIncomingRequest(ctx context.Context, requestID int) (*ListContactsResponse, error) { return nil, nil }
// func (c *Client) RejectIncomingRequest(ctx context.Context, requestID int) (*ListContactsResponse, error) { return nil, nil }
