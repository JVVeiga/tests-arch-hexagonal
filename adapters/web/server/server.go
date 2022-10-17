package server

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/jvveiga/tests-arch-hexagonal/adapters/web/handler"
	"github.com/jvveiga/tests-arch-hexagonal/app"
	"log"
	"net/http"
	"os"
	"time"
)

type Webserver struct {
	Service app.ProductServiceInterface
}

func MakeNewWebserver() *Webserver {
	return &Webserver{}
}

func (w Webserver) Serve() {

	rt := mux.NewRouter()
	ng := negroni.New(
		negroni.NewLogger(),
	)

	handler.MakeProductHandlers(rt, ng, w.Service)
	http.Handle("/", rt)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":9000",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log: ", log.Lshortfile),
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
