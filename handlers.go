package main

import (
	"errors"
	"github.com/gorilla/mux"
	"mime"
	"net/http"
)

func (service *OurService) createConfigHandler(w http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if mediatype != "application/json" {
		err := errors.New("Expect application/json Content-Type")
		http.Error(w, err.Error(), http.StatusUnsupportedMediaType)
		return
	}

	rt, err := decodeBody(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := createId()
	rt.Id = id
	service.Data[id] = rt
	renderJSON(w, rt)
}

func (service *OurService) getAllHandler(w http.ResponseWriter, req *http.Request) {
	allConfigs := []*OurConfig{}
	for _, v := range service.Data {
		allConfigs = append(allConfigs, v)
	}

	renderJSON(w, allConfigs)
}

func (service *OurService) getConfigHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	config, ok := service.Data[id]
	if !ok {
		err := errors.New("key not found")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	renderJSON(w, config)
}

func (service *OurService) delConfigHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	if v, ok := service.Data[id]; ok {
		delete(service.Data, id)
		renderJSON(w, v)
	} else {
		err := errors.New("key not found")
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}
