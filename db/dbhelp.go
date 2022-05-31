package db

import (
	"fmt"
	"github.com/google/uuid"
)

const (
	posts           = "posts/%s"
	all             = "posts"
	configIdVersion = "config/%s/%s"
)

func generateKey() (string, string) {
	id := uuid.New().String()
	return fmt.Sprintf(posts, id), id
}

func ccfgIdVer(id string, version string) string {
	return fmt.Sprintf(configIdVersion, id, version)
}
func constructKey(id string) string {
	return fmt.Sprintf(posts, id)
}
