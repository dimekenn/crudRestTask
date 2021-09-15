package configs

type Configs struct {
	Port string `json:"port"`
	DbUrl string 	`json:"db_url"`
}

func NewConfig() *Configs {
	return &Configs{}
}



