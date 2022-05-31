package main

//
//import (
//	"WebService/db"
//	"errors"
//	"mime"
//	"net/http"
//
//	"github.com/gorilla/mux"
//)
//
//func (service *postServer) createConfigGroupHandler(w http.ResponseWriter, req *http.Request) {
//	contentType := req.Header.Get("Content-Type")
//	mediatype, _, err := mime.ParseMediaType(contentType)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	if mediatype != "application/json" {
//		err := errors.New("expect application/json Content-Type")
//		http.Error(w, err.Error(), http.StatusUnsupportedMediaType)
//		return
//	}
//
//	rt, err := decodeConfigGroupBody(req.Body)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	id := createId()
//	rt.Id = id
//	service.Data[id] = rt
//	w.WriteHeader(http.StatusCreated)
//	renderJSON(w, rt)
//}
//
//func (service *postServer) getAllConfigGroupHandler(w http.ResponseWriter, req *http.Request) {
//	allConfigs := []*db.ConfigGroup{}
//	for _, v := range service.Data {
//		allConfigs = append(allConfigs, v)
//	}
//
//	renderJSON(w, allConfigs)
//}
//
//func (service *postServer) getConfigGroupHandler(w http.ResponseWriter, req *http.Request) {
//	id := mux.Vars(req)["id"]
//	config, ok := service.Data[id]
//	if !ok {
//		err := errors.New("key not found")
//		http.Error(w, err.Error(), http.StatusNotFound)
//		return
//	}
//	renderJSON(w, config)
//}
//
//func (service *postServer) delConfigGroupHandler(w http.ResponseWriter, req *http.Request) {
//	id := mux.Vars(req)["id"]
//	if v, ok := service.Data[id]; ok {
//		delete(service.Data, id)
//		renderJSON(w, v)
//	} else {
//		err := errors.New("key not found")
//		http.Error(w, err.Error(), http.StatusNotFound)
//	}
//}
