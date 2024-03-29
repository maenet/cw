package chatwork

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/go-querystring/query"
)

type PostMessageForm struct {
	Body       string `url:"body"`
	SelfUnread bool   `url:"self_unread,int"`
}

type PostMessageResponseBody struct {
	MessageID string `json:"message_id"`
}

func (c *Client) PostMessage(ctx context.Context, roomID string, form *PostMessageForm) (*PostMessageResponseBody, error) {
	roomIDInt, err := strconv.Atoi(roomID)
	if err != nil || roomIDInt < 0 {
		return nil, fmt.Errorf("roomID must be positive integer: %s", roomID)
	}
	spath := fmt.Sprintf("/rooms/%d/messages", roomIDInt)

	if form == nil {
		return nil, fmt.Errorf("missing form")
	}

	if len(form.Body) == 0 {
		return nil, fmt.Errorf("missing message body")
	}

	payload, err := query.Values(form)
	if err != nil {
		return nil, err
	}

	req, err := c.newRequest(ctx, http.MethodPost, spath, strings.NewReader(payload.Encode()))
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code = %v", res.StatusCode)
	}

	var body PostMessageResponseBody
	if err := decodeBody(res, &body); err != nil {
		return nil, err
	}

	return &body, nil
}
