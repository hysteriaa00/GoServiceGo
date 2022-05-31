package db

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/consul/api"
	"os"
)

type PostStore struct {
	cli *api.Client
}

func New() (*PostStore, error) {
	db := os.Getenv("DB")
	dbport := os.Getenv("DBPORT")

	config := api.DefaultConfig()
	config.Address = fmt.Sprintf("%s:%s", db, dbport)
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	return &PostStore{
		cli: client,
	}, nil
}

func (ps *PostStore) Get(id string, version string) (*Config, error) {
	kv := ps.cli.KV()
	key := ccfgIdVer(id, version)
	pair, _, err := kv.Get(key, nil)
	if err != nil {
		return nil, err
	}

	post := &Config{}
	err = json.Unmarshal(pair.Value, post)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (ps *PostStore) GetAll() ([]*Config, error) {
	kv := ps.cli.KV()
	data, _, err := kv.List(all, nil)
	if err != nil {
		return nil, err
	}

	posts := []*Config{}
	for _, pair := range data {
		post := &Config{}
		err = json.Unmarshal(pair.Value, post)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (ps *PostStore) Delete(id, version string) (map[string]string, error) {
	kv := ps.cli.KV()
	_, err := kv.Delete(ccfgIdVer(id, version), nil)
	if err != nil {
		return nil, err
	}

	return map[string]string{"Deleted": id + version}, nil
}

func (ps *PostStore) Post(post *Config) (*Config, error) {
	kv := ps.cli.KV()

	sid, rid := generateKey()
	post.Id = rid

	data, err := json.Marshal(post)
	if err != nil {
		return nil, err
	}

	p := &api.KVPair{Key: sid, Value: data}
	_, err = kv.Put(p, nil)
	if err != nil {
		return nil, err
	}

	return post, nil
}
