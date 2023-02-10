package apiserver

import (
	"WB0/internal/apiserver/orderservice"
	"encoding/json"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type httpServer struct {
	router  *mux.Router
	logger  *logrus.Logger
	service *orderservice.OrderService
}

func newHttpServer(logger *logrus.Logger, service *orderservice.OrderService) *httpServer {
	s := &httpServer{
		router:  mux.NewRouter(),
		logger:  logger,
		service: service,
	}

	s.configureRouter()

	return s
}

func (s *httpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *httpServer) configureRouter() {
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	s.router.HandleFunc("/test/{id}", s.testGET).Methods("GET")
	s.router.HandleFunc("/orders/{id}", s.GetOrderByUID).Methods("GET")
}

func (s *httpServer) GetOrderByUID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["id"]
	response, err := s.service.GetByUid(uid)
	if err != nil {
		s.logger.Info(err)
		respondWithError(w, http.StatusInternalServerError, "failed to find order with given uid")
	} else {
		respondWithJSON(w, http.StatusOK, response)
	}
}

func (s *httpServer) testGET(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	type resp struct {
		ID string `json:"id"`
	}
	response := &resp{
		ID: id,
	}
	respondWithJSON(w, http.StatusOK, response)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}