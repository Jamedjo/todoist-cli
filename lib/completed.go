package todoist

import (
	"context"
	"net/http"
	"net/url"
)

type Completed struct {
	Items    CompletedItems `json:"items"`
	Projects interface{}    `json:"projects"`
	Sections interface{}    `json:"sections"`
}

func (c *Client) CompletedAll(ctx context.Context, r *Completed) error {
	return c.doApi(ctx, http.MethodPost, "completed/get_all", url.Values{}, &r)
}
