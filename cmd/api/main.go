package main

import (
	"log"
	"os"

	"github.com/vitorpcruz/sso/config"
	"github.com/vitorpcruz/sso/internal/infra/db"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	config.LoadEnvConfig(dir)

	db.InitDB()
}
