package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

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
	Build: 99,
}

func server(ctx context.Context) (err error) {
	cnf := config.New()

	mux := http.NewServeMux()
	server := http.Server{
		Addr:    fmt.Sprintf("%s:%d", cnf.Addr, cnf.Port),
		Handler: mux,
	}

	db.Open(db.OpenConfig{
		Migrate: true,
	})

	mux.HandleFunc("/", router.RouterHandler)
	routes.Setup()

	log.Printf("Running version %s", version.String())

	go func() {
		if err = server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen: %+s\n", err)
		}
	}()

	log.Printf("Server are listen http://%s:%d", cnf.Addr, cnf.Port)

	<-ctx.Done()

	log.Println("Server is stoped!!!")

	ctxShutdown, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer func() {
		cancel()
	}()

	if err = server.Shutdown(ctxShutdown); err != nil {
		log.Fatalf("Shutdown failed: %+s", err)
	}

	log.Println("Server shutdown with success")

	if err == http.ErrServerClosed {
		err = nil
	}

	return
}

func main() {
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		oscall := <-channel
		log.Printf("System call: %+v", oscall)
		cancel()
	}()

	if err := server(ctx); err != nil {
		log.Fatalln(err)
	}
}
