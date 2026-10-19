package dto

import (
	"time"
	"uuid"

	"relat/internal/domain"
)

type ReportOutput struct {
	ID         uuid.UUID `json:"id"`
	FileUrl    string    `json:"file_url"`
	ReportedAt time.Time `json:"reported_at"`
	Active     bool      `json:"active"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func FromReport(report *domain.Report) *ReportOutput {
	return &ReportOutput{
		ID:         report.ID,
		FileUrl:    report.FileUrl,
		ReportedAt: report.ReportedAt,
		Active:     report.Active,
		CreatedAt:  report.CreatedAt,
		UpdatedAt:  report.UpdatedAt,
	}
}
