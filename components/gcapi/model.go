package gcapi

import "cloud.google.com/go/firestore"

type Client struct {
	firestore *firestore.Client
}