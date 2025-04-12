package server

import (
	"net/http"

	"github.com/RamiroCyber/gateway-go/internal/service"
	"github.com/RamiroCyber/gateway-go/internal/web/handler"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	router         *chi.Mux
	server         *http.Server
	accountService *service.AccountService
	port           string
}

func NewServer(accountService *service.AccountService, port string) *Server {
	server := &Server{
		router:         chi.NewRouter(),
		accountService: accountService,
		port:           port,
	}
	server.configureRoutes()
	return server
}

func (s *Server) configureRoutes() {
	accountHandler := handler.NewAccountHandler(s.accountService)
	s.router.Post("/accounts", accountHandler.Create)
	s.router.Get("/accounts", accountHandler.Get)
}

func (s *Server) Start() error {
	s.server = &http.Server{
		Addr:    ":" + s.port,
		Handler: s.router,
	}

	return s.server.ListenAndServe()
}
