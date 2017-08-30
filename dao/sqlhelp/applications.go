package sqlhelp

import (
	"database/sql"
	"time"

	. "github.com/ankyra/escape-registry/dao/types"
)

func (s *SQLHelper) AddApplication(app *Application) error {
	return s.PrepareAndExecInsert(s.AddApplicationQuery,
		app.Name,
		app.Project,
		app.Description,
		app.LatestVersion,
		app.Logo)
}
func (s *SQLHelper) UpdateApplication(app *Application) error {
	return s.PrepareAndExecUpdate(s.UpdateApplicationQuery,
		app.Description,
		app.LatestVersion,
		app.Logo,
		app.UploadedBy,
		app.UploadedAt.Unix(),
		app.Name,
		app.Project,
	)
}

func (s *SQLHelper) GetApplications(project string) (map[string]*Application, error) {
	rows, err := s.PrepareAndQuery(s.GetApplicationsQuery, project)
	if err != nil {
		return nil, err
	}
	return s.scanApplications(rows)
}

func (s *SQLHelper) GetApplication(project, name string) (*Application, error) {
	rows, err := s.PrepareAndQuery(s.GetApplicationQuery, project, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		return s.scanApplication(rows)
	}
	return nil, NotFound
}

func (s *SQLHelper) scanApplication(rows *sql.Rows) (*Application, error) {
	var name, project, description, latestVersion, logo, uploadedBy string
	var uploadedAt int64
	if err := rows.Scan(&name, &project, &description, &latestVersion, &logo, &uploadedBy, &uploadedAt); err != nil {
		return nil, err
	}
	return &Application{
		Name:          name,
		Project:       project,
		Description:   description,
		LatestVersion: latestVersion,
		Logo:          logo,
		UploadedBy:    uploadedBy,
		UploadedAt:    time.Unix(uploadedAt, 0),
	}, nil
}

func (s *SQLHelper) scanApplications(rows *sql.Rows) (map[string]*Application, error) {
	defer rows.Close()
	result := map[string]*Application{}
	for rows.Next() {
		app, err := s.scanApplication(rows)
		if err != nil {
			return nil, err
		}
		result[app.Name] = app
	}
	return result, nil
}