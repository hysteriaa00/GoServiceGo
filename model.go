package main

type OurService struct {
	Data map[string]*OurConfig `json:"data"`
}
type OurConfig struct {
	Entries map[string]string `json:"entries"`
}
