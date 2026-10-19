package http

import (
	"encoding/json/v2"
	"errors"
	"log/slog"
	"net/http"
	"strings"

	"relat/internal/platform/httperr"
)

type Response struct {
	StatusCode int    `json:"status_code"`
	StatusText string `json:"status_text"`
	Result     any    `json:"result"`
}

var fallbackBody = []byte(`{
	"status_code":500,
	"status_text":"internal server error",
	"result": {
		"cause":"server internal error occurred"
	}
}`)

func SendJson(w http.ResponseWriter, statusCode int, result any) {
	body, err := json.Marshal(Response{
		StatusCode: statusCode,
		StatusText: strings.ToLower(http.StatusText(statusCode)),
		Result:     result,
	})
	if err != nil {
		slog.Error("failed to marshal response", slog.Any("error", err))
		write(w, http.StatusInternalServerError, fallbackBody)
		return
	}

	write(w, statusCode, body)
}

func SendError(w http.ResponseWriter, err error) {
	var httpErr *httperr.HttpError
	if !errors.As(err, &httpErr) {
		httpErr = httperr.ErrInternalServerError()
	}

	if httpErr.StatusCode >= http.StatusInternalServerError {
		slog.Error("request failed", slog.Any("error", err))
	}

	SendJson(w, httpErr.StatusCode, httpErr)
}

func SendStatus(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
}

func write(w http.ResponseWriter, statusCode int, body []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if _, err := w.Write(body); err != nil {
		slog.Warn("failed to write response", slog.Any("error", err))
	}
}
