package http

import (
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/sokolovvs/bank-currencies/internal/http/api/v1"
	"net/http"
)

type Server struct {
	router *mux.Router
	api    *Api
	port   string
}

func NewServer(appPort string) *Server {
	return &Server{router: mux.NewRouter(), api: &Api{v1: new(v1.HttpApiV1Controller)}, port: appPort}
}

func (s *Server) Serve() {

	s.router.HandleFunc("/health-check", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "running")
	})
	s.router.HandleFunc("/api/v1/banks", s.api.v1.GetBanks).Methods("GET")
	s.router.HandleFunc("/api/v1/currencies", s.api.v1.GetCurrencies).Methods("GET")

	err := http.ListenAndServe(fmt.Sprintf(":%s", s.port), s.router)

	if err != nil {
		log.Error("Listen and serve was failed, err: ", err, " port: ", s.port)
		panic(err)
	}
}

type Api struct {
	v1 *v1.HttpApiV1Controller
}
