package gcapi

import (
	"context"
	"fmt"

	"github.com/gangozero/hackprague2019-be/generated/models"
	"google.golang.org/api/iterator"
)

func (c *Client) GetProfiles() ([]*models.Profile, error) {
	result := []*models.Profile{}

	ctx := context.Background()
	iter := c.firestore.Collection("profiles").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("Failed to iterate: %v", err)
		}
		data := doc.Data()

		result = append(result, &models.Profile{
			ID:   data["id"].(string),
			Name: data["name"].(string),
		})
	}

	return result, nil
}
