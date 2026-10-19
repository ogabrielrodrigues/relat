package http

import (
	"net/http"
	"uuid"

	"relat/internal/platform/httperr"
)

func extractID(r *http.Request) (uuid.UUID, error) {
	if r.PathValue("id") == "" {
		return uuid.Nil(), httperr.NewHttpError("param id is required", http.StatusBadRequest)
	}

	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		return uuid.Nil(), httperr.NewHttpError("param id invalid format", http.StatusBadRequest)
	}

	return id, nil
}
