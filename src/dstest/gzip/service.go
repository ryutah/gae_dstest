package gzip

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

const Kind = "Gzip"

func NewService(ctx context.Context) GzipServivce {
	return &gzipService{ctx}
}

type GzipServivce interface {
	Post(gz *Gzip) (*datastore.Key, error)
	Get(id int64) (*Gzip, error)
}

type gzipService struct {
	ctx context.Context
}

func (g *gzipService) Post(gz *Gzip) (*datastore.Key, error) {
	key := NewKey(g.ctx)
	return datastore.Put(g.ctx, key, gz)
}

func (g *gzipService) Get(id int64) (*Gzip, error) {
	key := Key(g.ctx, id)
	gz := new(Gzip)
	err := datastore.Get(g.ctx, key, gz)
	return gz, err
}

func NewKey(ctx context.Context) *datastore.Key {
	return datastore.NewIncompleteKey(ctx, Kind, nil)
}

func Key(ctx context.Context, id int64) *datastore.Key {
	return datastore.NewKey(ctx, Kind, "", id, nil)
}
