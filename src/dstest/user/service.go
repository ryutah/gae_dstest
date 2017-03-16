package user

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

const Kind = "User"

func NewService(ctx context.Context) UserService {
	return &userService{ctx}
}

type UserService interface {
	Create(user *User) (*datastore.Key, error)
	Get(id int64) (*User, error)
	List(query *datastore.Query) ([]*User, error)
	Update(id int64, user *User) error
	Delete(id int64) error
	Count() (int, error)
}

type userService struct {
	ctx context.Context
}

func (u *userService) Create(user *User) (*datastore.Key, error) {
	key := NewKey(u.ctx)
	return datastore.Put(u.ctx, key, user)
}

func (u *userService) Get(id int64) (*User, error) {
	key := Key(u.ctx, id)
	user := new(User)
	log.Debugf(u.ctx, "Get User %v", user.Key)
	if err := datastore.Get(u.ctx, key, user); err != nil {
		return nil, err
	}
	user.Key, user.ID = key, key.IntID()
	return user, nil
}

func (u *userService) List(query *datastore.Query) ([]*User, error) {
	var users []*User
	_, err := query.GetAll(u.ctx, &users)
	return users, err
}

func (u *userService) Update(id int64, user *User) error {
	key := Key(u.ctx, id)
	_, err := datastore.Put(u.ctx, key, user)
	return err
}

func (u *userService) Delete(id int64) error {
	key := Key(u.ctx, id)
	return datastore.Delete(u.ctx, key)
}

func (u *userService) Count() (int, error) {
	keys, err := Query().KeysOnly().GetAll(u.ctx, nil)
	return len(keys), err
}

func NewKey(ctx context.Context) *datastore.Key {
	return datastore.NewIncompleteKey(ctx, Kind, nil)
}

func Key(ctx context.Context, id int64) *datastore.Key {
	return datastore.NewKey(ctx, Kind, "", id, nil)
}

func Query() *datastore.Query {
	return datastore.NewQuery(Kind)
}
