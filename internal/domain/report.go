package domain

import (
	"time"
	"uuid"
)

type Report struct {
	ReportedAt time.Time
	FileUrl    string
	DefaultAttrs
}

func NewReport(fileUrl string) *Report {
	now := time.Now()

	return &Report{
		ReportedAt: now,
		FileUrl:    fileUrl,
		DefaultAttrs: DefaultAttrs{
			ID:        uuid.NewV7(),
			Active:    true,
			CreatedAt: now,
			UpdatedAt: now,
		},
	}
}
