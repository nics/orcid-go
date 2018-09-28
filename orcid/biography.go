package orcid

import (
	"fmt"
	"net/http"
)

type Biography struct {
	Content          *string   `json:"content,omitempty"`
	CreatedDate      *IntValue `json:"created-date,omitempty"`
	LastModifiedDate *IntValue `json:"last-modified-date,omitempty"`
	Path             *string   `json:"path,omitempty"`
	Visibility       *string   `json:"visibility,omitempty"`
}

func (c *Client) Biography(orcid string) (*Biography, *http.Response, error) {
	data := new(Biography)
	path := fmt.Sprintf("%s/biography", orcid)
	res, err := c.get(path, data)
	return data, res, err
}
