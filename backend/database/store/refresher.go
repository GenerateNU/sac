package store

import (
	"context"
	"time"
)

type Refresher interface {
	Get(ctx context.Context, userID string) (string, error)
	Set(ctx context.Context, userID string, token string, expiry time.Duration) error
	Delete(ctx context.Context, userID string) error
	Exists(ctx context.Context, userID string) (bool, error)
}

type Refresh struct {
	store Storer
}

func NewRefresh(store Storer) *Refresh {
	return &Refresh{
		store: store,
	}
}

func (r *Refresh) Get(ctx context.Context, userID string) (string, error) {
	return r.store.Get(ctx, userID)
}

func (r *Refresh) Set(ctx context.Context, userID string, token string, expiry time.Duration) error {
	return r.store.Set(ctx, userID, token, expiry)
}

func (r *Refresh) Delete(ctx context.Context, userID string) error {
	return r.store.Delete(ctx, userID)
}

func (r *Refresh) Exists(ctx context.Context, userID string) (bool, error) {
	return r.store.Exists(ctx, userID)
}
