package http

import (
	"fmt"
	"net/http"

	"relat/internal/platform/httperr"
)

func (s *Server) createReport(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(256); err != nil {
		SendError(w, httperr.NewHttpError(err.Error(), http.StatusInternalServerError))
		return
	}

	file := r.MultipartForm.File["file"][0]

	fileUrl, err := s.storage.SaveFile(file)
	if err != nil {
		SendError(w, err)
		return
	}

	id, err := s.reportUseCase.CreateReport(fileUrl)
	if err != nil {
		SendError(w, err)
		return
	}

	w.Header().Add("Location", fmt.Sprintf("/report/%s", id))
	SendStatus(w, http.StatusCreated)
}

func (s *Server) findReport(w http.ResponseWriter, r *http.Request) {
	id, err := extractID(r)
	if err != nil {
		SendError(w, err)
		return
	}

	report, err := s.reportUseCase.FindReport(id)
	if err != nil {
		SendError(w, err)
		return
	}

	SendJson(w, http.StatusOK, report)
}

func (s *Server) listReports(w http.ResponseWriter, r *http.Request) {
	reports, err := s.reportUseCase.ListReports()
	if err != nil {
		SendError(w, httperr.NewHttpError("error listing reports", http.StatusInternalServerError))
		return
	}

	SendJson(w, http.StatusOK, reports)
}
