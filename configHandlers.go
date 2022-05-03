package main

import (
	"errors"
	"mime"
	"net/http"

	"github.com/gorilla/mux"
)

func (configService *ConfigService) createConfigHandler(w http.ResponseWriter, req *http.Request) {
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

	id := createId()
	rt.Id = id
	configService.Data[id] = rt
	w.WriteHeader(http.StatusCreated)
	renderJSON(w, rt)
}

func (configService *ConfigService) getAllConfigHandler(w http.ResponseWriter, req *http.Request) {
	allConfigs := []*Config{}
	for _, v := range configService.Data {
		allConfigs = append(allConfigs, v)
	}

	renderJSON(w, allConfigs)
}

func (configService *ConfigService) getConfigHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	config, ok := configService.Data[id]
	if !ok {
		err := errors.New("key not found")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	renderJSON(w, config)
}

func (configService *ConfigService) delConfigHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	if v, ok := configService.Data[id]; ok {
		delete(configService.Data, id)
		renderJSON(w, v)
	} else {
		err := errors.New("key not found")
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}
