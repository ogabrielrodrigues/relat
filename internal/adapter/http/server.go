package http

import (
	"net/http"

	"relat/internal/adapter/postgres"
	"relat/internal/adapter/usecase"
	"relat/internal/platform/config"
	"relat/internal/platform/storage"
)

type Server struct {
	mux           *http.ServeMux
	storage       *storage.Storage
	reportUseCase *usecase.ReportUseCase
}

func NewServer() (*Server, error) {
	storage, err := storage.NewStorage(config.EnvReportFolder)
	if err != nil {
		return nil, err
	}

	db, err := postgres.NewPostgres()
	if err != nil {
		return nil, err
	}

	return &Server{
		mux:           http.NewServeMux(),
		storage:       storage,
		reportUseCase: usecase.NewReportUseCase(db),
	}, nil
}

func (s *Server) MountHandlers() http.Handler {
	s.mux.HandleFunc("GET /report/{id}", s.findReport)
	s.mux.HandleFunc("POST /report", s.createReport)
	s.mux.HandleFunc("GET /reports", s.listReports)

	return chain(s.mux, logMiddleware)
}
