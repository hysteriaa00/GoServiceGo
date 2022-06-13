package db

import (
	"fmt"
	"github.com/google/uuid"
)

const (
	config          = "config/%s"
	all             = "config"
	configIdVersion = "config/%s/%s"
)

func generateKey() (string, string) {
	id := uuid.New().String()
	return fmt.Sprintf(config, id), id
}
func createNewConfigWithVersionKey(version string) (string, string) {
	id := uuid.New().String()
	return fmt.Sprintf(configIdVersion, id, version), id
}
func createConfigWithIdAndVersionKey(id string, version string) string {
	return fmt.Sprintf(configIdVersion, id, version)
}

func constructKey(id string) string {
	return fmt.Sprintf(config, id)
}
