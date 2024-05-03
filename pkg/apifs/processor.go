package apifs

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

type Processor struct {
	rootFs    string
	apiPrefix string
	port      int

	router *mux.Router
}

func NewProcessor(path string, prefix string, port int) (*Processor, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, ErrWithError(ErrorInvalid, err)
	}

	if !info.IsDir() {
		return nil, ErrWithString(ErrorInvalid, "not a directory")
	}

	return &Processor{
		rootFs:    path,
		apiPrefix: prefix,
		port:      port,
		router:    mux.NewRouter(),
	}, nil
}

func (p *Processor) Run() error {
	p.router.PathPrefix(p.apiPrefix).Handler(http.StripPrefix(p.apiPrefix, http.FileServer(http.Dir(p.rootFs))))
	p.router.Use(p.loggingMiddleware)

	addr := fmt.Sprintf("0.0.0.0:%d", p.port)
	srv := &http.Server{
		Handler:      p.router,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return srv.ListenAndServe()
}

func (p *Processor) finish(w http.ResponseWriter, err error) {
	w.Header().Add("content-type", "application/json")

	d := json.NewEncoder(w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		d.Encode(&Status{
			Status: false,
			Error:  err.Error(),
		})
	} else {
		d.Encode(&Status{
			Status: true,
		})
	}
}

func (p *Processor) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
