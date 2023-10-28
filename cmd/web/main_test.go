package main

import (
	"context"
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"subscription-service/data"
	"sync"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
)

var testConfig Config

// TestMain is the entry point for all tests
func TestMain(m *testing.M) {
	gob.Register(data.User{})

	tmpPath = "./../../tmp"
	pathToManual = "./../../pdf"

	// set up session
	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = true

	testConfig = Config{
		Session:       session,
		DB:            nil,
		InfoLog:       log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		ErrorLog:      log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		Models:        data.TestNew(nil),
		Wait:          &sync.WaitGroup{},
		ErrorChan:     make(chan error),
		ErrorChanDone: make(chan bool),
	}

	// dummy mailer
	errorChan := make(chan error)
	mailerChan := make(chan Message, 100)
	mailerDoneChan := make(chan bool)

	testConfig.Mailer = Mail{
		Wait:       testConfig.Wait,
		ErrorChan:  errorChan,
		MailerChan: mailerChan,
		DoneChan:   mailerDoneChan,
	}

	go func() {
		for {
			select {
			case <-testConfig.Mailer.MailerChan:
				testConfig.Wait.Done()
			case <-testConfig.Mailer.ErrorChan:
			case <-testConfig.Mailer.DoneChan:
				return
			}
		}
	}()

	go func() {
		for {
			select {
			case err := <-testConfig.ErrorChan:
				testConfig.ErrorLog.Println(err)
			case <-testConfig.ErrorChanDone:
				return
			}
		}
	}()

	os.Exit(m.Run())
}

func ctx(req *http.Request) context.Context {
	ctx, err := testConfig.Session.Load(req.Context(), req.Header.Get("X-Session"))
	if err != nil {
		log.Println(err)
	}
	return ctx
}
