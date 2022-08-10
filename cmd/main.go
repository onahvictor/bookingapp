package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/onahvictor/BookingApp/internal/config"
	"github.com/onahvictor/BookingApp/internal/handlers"
	"github.com/onahvictor/BookingApp/internal/render"
)

var app config.AppConfig
var portNumber string = ":8080"
var session *scs.SessionManager

func main() {

	app.InProduction = false
	//change this to true when in production

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	fmt.Printf("Server is running on Port %v\n", portNumber)
	err = srv.ListenAndServe()
	log.Fatal(err)
}