package main

import "github.com/rodrigopmatias/ligistic/framework/db"

func main() {
	db.Open(db.OpenConfig{
		Migrate: true,
	})
}
