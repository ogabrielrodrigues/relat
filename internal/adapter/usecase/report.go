package usecase

import (
	"uuid"

	"relat/internal/adapter/dto"
	"relat/internal/adapter/postgres"
	"relat/internal/domain"
)

var reports = []domain.Report{}

type ReportUseCase struct {
	db *postgres.Postgres
}

func NewReportUseCase(db *postgres.Postgres) *ReportUseCase {
	return &ReportUseCase{db}
}

func (u *ReportUseCase) CreateReport(fileUrl string) (uuid.UUID, error) {
	report := domain.NewReport(fileUrl)

	if err := u.db.CreateReport(report); err != nil {
		return uuid.Nil(), err
	}

	return report.ID, nil
}

func (u *ReportUseCase) FindReport(id uuid.UUID) (*dto.ReportOutput, error) {
	return u.db.FindReport(id)
}

func (u *ReportUseCase) ListReports() (*[]dto.ReportOutput, error) {
	return u.db.ListReports()
}
