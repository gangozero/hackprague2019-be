package gcapi

import (
	"context"

	"cloud.google.com/go/firestore"
)

func NewClient(projectID string) (*Client, error) {

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	return &Client{
		firestore: client,
	}, nil
}

func (c *Client) Close() {
	c.firestore.Close()
}
