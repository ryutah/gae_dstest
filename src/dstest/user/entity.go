package user

import (
	"google.golang.org/appengine/datastore"
)

type User struct {
	Key   *datastore.Key `json:"-" datastore:"-"`
	ID    int64          `json:"id" datastore:"-"`
	Email string         `json:"email"`
	Name  string         `json:"name"`
}
