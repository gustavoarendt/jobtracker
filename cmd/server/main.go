package main

import (
	"github.com/gustavoarendt/jobtracker/configs"
	"github.com/gustavoarendt/jobtracker/internal/database"
	"github.com/gustavoarendt/jobtracker/internal/routes"
)

func main() {
	config, err := configs.Configure("C:\\Users\\Gus\\projects\\go\\jobtracker\\cmd\\server")
	if err != nil {
		panic(err)
	}
	database.DbConnection(config)
	database.JwtConnection(config)
	routes.HandleRequest(config)
}
