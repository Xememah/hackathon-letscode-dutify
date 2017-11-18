package main

import (
	"log"

	"os"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"repo.letscode.sii.pl/wroclaw/three/backend/app"
)

func main() {
	logger := log.New(os.Stdout, "LetsCode", log.Ldate|log.Lshortfile)
	app := &app.Application{Logger: logger}

	err := app.Init()
	if err != nil {
		logger.Fatal(err)
	}

	err = app.Serve()
	if err != nil {
		logger.Fatal(err)
	}
}
