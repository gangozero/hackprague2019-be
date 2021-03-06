package gcapi

import (
	"context"
	"fmt"
	"time"

	"github.com/gangozero/hackprague2019-be/generated/models"
	"google.golang.org/api/iterator"
	"google.golang.org/genproto/googleapis/type/latlng"
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

func fs2sample(data map[string]interface{}) *models.Sample {
	geo := data["geo"].(*latlng.LatLng)
	ts := data["ts"].(time.Time)
	return &models.Sample{
		UserID: data["user_id"].(string),
		Grade:  data["grade"].(int64),
		Lat:    geo.Latitude,
		Lon:    geo.Longitude,
		Ts:     ts.Unix(),
	}
}

func sample2fs(profileID string, data *models.Sample) map[string]interface{} {
	return map[string]interface{}{
		"user_id": data.UserID,
		"grade":   data.Grade,
		"ts":      time.Now(),
		"profile": profileID,
		"geo": &latlng.LatLng{
			Latitude:  data.Lat,
			Longitude: data.Lon,
		},
	}
}

func (c *Client) GetGradesByID(profileID string) ([]*models.Sample, error) {
	result := []*models.Sample{}

	ctx := context.Background()
	q := c.firestore.Collection("grades").
		Where("profile", "==", profileID)
	iter := q.Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("Failed to iterate: %v", err)
		}

		result = append(result, fs2sample(doc.Data()))
	}

	return result, nil
}

func (c *Client) GetGradesByIDAndUser(profileID, userID string) ([]*models.Sample, error) {
	result := []*models.Sample{}

	ctx := context.Background()
	q := c.firestore.Collection("grades").
		Where("user_id", "==", userID).
		Where("profile", "==", profileID)
	iter := q.Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("Failed to iterate: %v", err)
		}

		result = append(result, fs2sample(doc.Data()))
	}

	return result, nil
}

func (c *Client) PostGrade(profileID string, data *models.Sample) error {
	ctx := context.Background()
	_, _, err := c.firestore.Collection("grades").
		Add(ctx, sample2fs(profileID, data))
	if err != nil {
		fmt.Errorf("Failed to add sample: %v", err)
	}
	return nil
}
