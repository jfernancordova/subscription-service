package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"subscription-service/data"
	"subscription-service/internal/database"
	"subscription-service/internal/session"
	"sync"
	"syscall"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	// connect to the database
	db := database.Init()

	// create sessions
	session := session.Init()

	// create loggers
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// create waitgroup
	wg := sync.WaitGroup{}

	// set up the application config
	app := config{
		session:  session,
		db:       db,
		infoLog:  infoLog,
		errorLog: errorLog,
		wait:     &wg,
		models:   data.New(db),
	}

	// listen for signals
	go app.listenForShutdown()

	// listen for web connections
	app.serve()
}

func (app *config) serve() {
	// start http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", "80"),
		Handler: app.routes(),
	}

	app.infoLog.Println("Starting web server...")
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func (app *config) listenForShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	app.shutdown()
	os.Exit(0)
}

func (app *config) shutdown() {
	// perform any cleanup tasks
	app.infoLog.Println("would run cleanup tasks...")

	// block until waitgroup is empty
	app.wait.Wait()

	app.infoLog.Println("closing channels and shutting down application...")
}
