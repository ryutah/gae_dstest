package gzip

import (
	app "dstest/middleware"
	"dstest/util"
	"net/http"
	"strconv"

	"google.golang.org/appengine/datastore"

	"github.com/gorilla/mux"
)

func Post(w http.ResponseWriter, r *http.Request) {
	gz := new(Gzip)
	if err := util.ParseBody(r, gz); err != nil {
		util.ResponseBadRequest(w, err.Error())
		return
	}
	s := NewService(app.Ctx())
	key, err := s.Post(gz)
	if err != nil {
		util.ResponseServerError(w, err.Error())
		return
	}
	resp := map[string]int64{"id": key.IntID()}
	w.WriteHeader(http.StatusCreated)
	util.ResponseWithJson(w, resp)
}

func Get(w http.ResponseWriter, r *http.Request) {
	sID := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(sID, 0, 64)
	if err != nil {
		util.ResponseBadRequest(w, err.Error())
		return
	}

	s := NewService(app.Ctx())
	gz, err := s.Get(id)
	if err == datastore.ErrNoSuchEntity {
		http.NotFound(w, r)
		return
	} else if err != nil {
		util.ResponseServerError(w, err.Error())
		return
	}
	util.ResponseWithJson(w, gz)
}
