package main

type OurService struct {
	Data map[string]*OurConfig `json:"data"`
}

type OurConfig struct {
	Id      string            `json:"id"`
	Entries map[string]string `json:"entries"`
}
