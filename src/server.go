package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mariacalinoiu/salesApp/src/datasources"
	"github.com/mariacalinoiu/salesApp/src/handlers"
)

type server struct {
	mux    *http.ServeMux
	logger *log.Logger
}

type option func(*server)

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.log("Method: %s, Path: %s", r.Method, r.URL.Path)
	s.mux.ServeHTTP(w, r)
}

func (s *server) log(format string, v ...interface{}) {
	s.logger.Printf(format+"\n", v...)
}

func logWith(logger *log.Logger) option {
	return func(s *server) {
		s.logger = logger
	}
}

func setup(logger *log.Logger, db datasources.DBClient) *http.Server {
	server := newServer(db, logWith(logger))
	return &http.Server{
		Addr:         ":8081",
		Handler:      server,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  600 * time.Second,
	}
}

func newServer(db datasources.DBClient, options ...option) *server {
	s := &server{logger: log.New(ioutil.Discard, "", 0)}

	for _, o := range options {
		o(s)
	}

	s.mux = http.NewServeMux()

	s.mux.HandleFunc("/articole",
		func(w http.ResponseWriter, r *http.Request) {
			handlers.HandleArticole(w, r, db, s.logger)
		},
	)
	s.mux.HandleFunc("/parteneri",
		func(w http.ResponseWriter, r *http.Request) {
			handlers.HandleParteneri(w, r, db, s.logger)
		},
	)
	s.mux.HandleFunc("/vanzatori",
		func(w http.ResponseWriter, r *http.Request) {
			handlers.HandleVanzatori(w, r, db, s.logger)
		},
	)
	s.mux.HandleFunc("/vanzari",
		func(w http.ResponseWriter, r *http.Request) {
			handlers.HandleVanzari(w, r, db, s.logger)
		},
	)
	s.mux.HandleFunc("/sucursale",
		func(w http.ResponseWriter, r *http.Request) {
			handlers.HandleSucursale(w, r, db, s.logger)
		},
	)
	s.mux.HandleFunc("/proiecte",
		func(w http.ResponseWriter, r *http.Request) {
			handlers.HandleProiecte(w, r, db, s.logger)
		},
	)
	s.mux.HandleFunc("/cors",
		func(w http.ResponseWriter, r *http.Request) {
			if origin := r.Header.Get("Origin"); origin != "" {
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
				w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token")
				w.Header().Set("Access-Control-Expose-Headers", "Authorization")
			}
			return
		})

	return s
}

func main() {
	logger := log.New(os.Stdout, "", 0)
	db := datasources.GetClient("user", "password", "hostname", "salesapp")
	hs := setup(logger, db)

	logger.Printf("Listening on http://localhost%s\n", hs.Addr)
	go func() {
		if err := hs.ListenAndServe(); err != nil {
			logger.Println(err)
		}
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	<-signals

	logger.Println("Shutting down webserver.")
	os.Exit(0)
}
