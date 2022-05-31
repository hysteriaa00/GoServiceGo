package db

type ConfigService struct {
	Data map[string]*Config `json:"data"`
}

type ConfigGroupService struct {
	Data map[string]*ConfigGroup `json:"data"`
}

type Config struct {
	Id      string            `json:"id"`
	Entries map[string]string `json:"entries"`
	//Version string            `json:"version"`
}

type ConfigGroup struct {
	Id    string   `json:"id"`
	Items []Config `json:"items"`
}
