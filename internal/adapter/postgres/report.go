package postgres

import (
	"errors"
	"net/http"
	"uuid"

	"relat/internal/adapter/dto"
	"relat/internal/domain"
	"relat/internal/platform/httperr"

	"github.com/jackc/pgx/v5"
)

func (p *Postgres) CreateReport(report *domain.Report) error {
	if _, err := p.pool.Exec(p.ctx,
		`insert into "report" values ($1, $2, $3, $4, $5, $6)`,
		report.ID, report.FileUrl, report.ReportedAt, report.Active, report.CreatedAt, report.UpdatedAt,
	); err != nil {
		return httperr.NewHttpError("error creating report", http.StatusInternalServerError)
	}

	return nil
}

func (p *Postgres) FindReport(id uuid.UUID) (*dto.ReportOutput, error) {
	row := p.pool.QueryRow(p.ctx, `select * from "report" where id = $1`, id)

	var report dto.ReportOutput
	if err := row.Scan(
		&report.ID,
		&report.FileUrl,
		&report.ReportedAt,
		&report.Active,
		&report.CreatedAt,
		&report.UpdatedAt,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, httperr.ErrNotFound("report")
		}

		return nil, httperr.NewHttpError("unable to find report", http.StatusInternalServerError)
	}

	return &report, nil
}

func (p *Postgres) ListReports() (*[]dto.ReportOutput, error) {
	rows, err := p.pool.Query(p.ctx, `select * from "report"`)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, httperr.NewHttpError("unable to list reports", http.StatusInternalServerError)
	}

	var report dto.ReportOutput
	var reports []dto.ReportOutput

	for rows.Next() {
		defer rows.Close()

		if err := rows.Scan(
			&report.ID,
			&report.FileUrl,
			&report.ReportedAt,
			&report.Active,
			&report.CreatedAt,
			&report.UpdatedAt,
		); err != nil {
			return nil, httperr.NewHttpError("error scanning reports", http.StatusInternalServerError)
		}

		reports = append(reports, report)
	}

	return &reports, nil
}
