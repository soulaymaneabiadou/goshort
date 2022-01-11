package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/soulaymaneabiadou/goshort"
)

type Server struct {
	*mux.Router

	urls []goshort.URL
}

func NewServer() *Server {
	s := &Server{
		Router: mux.NewRouter(),
		urls:   []goshort.URL{},
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.HandleFunc("/shorten", s.shortenUrl()).Methods("POST")
	s.HandleFunc("/{code}", s.getUrl()).Methods("GET")
}

func (s *Server) getUrl() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var u goshort.URL
		var err error

		code := mux.Vars(r)["code"]

		u, err = goshort.GetUrl(code)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		s.urls = append(s.urls, u)

		rw.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(rw).Encode(u); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) shortenUrl() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var u goshort.URL

		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		u, _ = goshort.ShortenUrl(u.LongUrl)
		s.urls = append(s.urls, u)

		rw.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(rw).Encode(u); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

	}
}
