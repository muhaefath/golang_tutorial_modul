package main

import (
	"context"
)

type WebhookHandler interface {
	NewOrder(ctx context.Context, body []byte) error
	UpsertOrder(ctx context.Context, body []byte) error
	CancelOrder(ctx context.Context, body []byte) error
	GetFullOrder(ctx context.Context) error
	DeleteWebhook(ctx context.Context, storeID int) error
}
