package repositories

import (
	"database/sql"
	"go-api/internal/models"
)

type ProspectRepository struct {
	DB *sql.DB
}

func (r *ProspectRepository) FetchAll() (models.ProspectCount, error) {
	rows, err := r.DB.Query("select count(Id) as RowCount from Prospects")
	if err != nil {
		return models.ProspectCount{}, err
	}
	defer rows.Close()

	var count models.ProspectCount

	if rows.Next() {
		if err := rows.Scan(&count.RowCount); err != nil {
			return models.ProspectCount{}, err
		}
	}

	if err := rows.Err(); err != nil {
		return models.ProspectCount{}, err
	}

	return count, nil
}
