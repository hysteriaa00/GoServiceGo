package main

import (
	"WebService/db"
	"encoding/json"
	"io"
	"net/http"

	"github.com/google/uuid"
)

func decodeConfigBody(r io.Reader) (*db.Config, error) {
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	var rt db.Config
	if err := dec.Decode(&rt); err != nil {
		return nil, err
	}
	return &rt, nil
}

func decodeConfigGroupBody(r io.Reader) (*db.ConfigGroup, error) {
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()
	var rt db.ConfigGroup
	if err := dec.Decode(&rt); err != nil {
		return nil, err
	}
	return &rt, nil
}

func renderJSON(w http.ResponseWriter, v interface{}) {
	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func createId() string {
	return uuid.New().String()
}
