package main

import (
	"WebService/db"
	"errors"
	"fmt"
	"mime"
	"net/http"

	"github.com/gorilla/mux"
)

type postServer struct {
	store *db.PostStore
}

func (configService *postServer) createConfigHandler(w http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if mediatype != "application/json" {
		err := errors.New("expect application/json Content-Type")
		http.Error(w, err.Error(), http.StatusUnsupportedMediaType)
		return
	}

	rt, err := decodeConfigBody(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cfg, err := configService.store.Post(rt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	renderJSON(w, cfg)
}

func (configService *postServer) getAllConfigHandler(w http.ResponseWriter, req *http.Request) {
	allConfigs, err := configService.store.GetAll()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

		renderJSON(w, allConfigs)
	}
}

func (configService *postServer) getConfigHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req)

	id := mux.Vars(req)["id"]
	version := mux.Vars(req)["version"]
	cfg, err := configService.store.Get(id, version)

	if err != nil {
		err := errors.New("key not found")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	renderJSON(w, cfg)
}

func (configService *postServer) delConfigHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	version := mux.Vars(req)["version"]
	cfg, ok := configService.store.Delete(id, version)
	if ok != nil {
		err := errors.New("key not found")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	renderJSON(w, cfg)
}
