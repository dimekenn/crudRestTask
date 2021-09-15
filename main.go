package main

import (
	"context"
	"crudRestTask/configs"
	"crudRestTask/internal/app"
	"embed"
	"encoding/json"
	_ "github.com/lib/pq"
	"log"
)

//go:embed configs.json
var fs embed.FS

const configName = "configs.json"

func main()  {
	//reading json file for configs
	data, readErr := fs.ReadFile(configName)
	if readErr != nil {
		log.Fatal(readErr)
	}
	//creating config entity to deserialize configs.json
	cfg := configs.NewConfig()
	if unmErr := json.Unmarshal(data, &cfg); unmErr != nil {
		log.Fatal(unmErr)
	}
	//creating context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//channel for errors
	errCh := make(chan error, 1)

	//starting new goroutine for api server
	go app.StartHTTPServer(ctx, cfg, errCh)

	log.Fatalf("Terminated: %s", <-errCh)
}
