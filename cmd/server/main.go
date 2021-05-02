package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rodrigopmatias/ligistic/framework/config"
	"github.com/rodrigopmatias/ligistic/framework/db"
	"github.com/rodrigopmatias/ligistic/framework/router"
	"github.com/rodrigopmatias/ligistic/routes"
)

type Version struct {
	Major int8
	Minor int8
	Build int16
}

func (version Version) String() string {
	return fmt.Sprintf("%d.%d.%d", version.Major, version.Minor, version.Build)
}

var version = Version{
	Major: 0,
	Minor: 0,
	Build: 5,
}

func main() {
	http.HandleFunc("/", router.RouterHandler)
	cnf := config.New()

	routes.Setup()
	db.Open(db.OpenConfig{
		Migrate: true,
	})

	log.Printf("Running version %s", version.String())
	log.Printf("Server are listen http://%s:%d", cnf.Addr, cnf.Port)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf("%s:%d", cnf.Addr, cnf.Port), nil))
}
