package user

import (
	app "dstest/middleware"
	"dstest/util"
	"fmt"
	"net/http"
	"strconv"

	"google.golang.org/appengine/datastore"

	"github.com/gorilla/mux"
)

type postResponse struct {
	ID int64 `json:"id"`
}

func Post(w http.ResponseWriter, r *http.Request) {
	u := new(User)
	if err := util.ParseBody(r, u); err != nil {
		util.ResponseBadRequest(w, err.Error())
		return
	}

	s := NewService(app.Ctx())
	key, err := s.Create(u)
	if err != nil {
		util.ResponseServerError(w, err.Error())
		return
	}

	resp := postResponse{key.IntID()}
	w.WriteHeader(http.StatusCreated)
	util.ResponseWithJson(w, resp)
}

func Get(w http.ResponseWriter, r *http.Request) {
	sID := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(sID, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("Cant use user id %v", sID)
		util.ResponseBadRequest(w, msg)
		return
	}

	s := NewService(app.Ctx())
	u, err := s.Get(id)
	if err == datastore.ErrNoSuchEntity {
		http.NotFound(w, r)
		return
	} else if err != nil {
		util.ResponseServerError(w, err.Error())
		return
	}

	util.ResponseWithJson(w, u)
}
