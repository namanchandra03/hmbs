package dbHelper

import (
	"golang.org/x/net/context"
	"hms/database"
	"hms/models"
)

func AddComplain(ctx context.Context, complainInfo models.ComplainDetails) (string, error) {
	args := []interface{}{
		complainInfo.StudentID,
		complainInfo.HostelID,
		complainInfo.ComplainType,
		complainInfo.Description,
	}
	SQL := `INSERT INTO complains (student_id, hostel_id, complain_type, description)
VALUES ($1, $2, $3, $4) RETURNING id`

	var complainID string
	err := database.DB.GetContext(ctx, &complainID, SQL, args...)
	return complainID, err
}

func UpdateComplain(ctx context.Context, updatedComplain models.ComplainDetails, complainID string) error {
	args := []interface{}{
		updatedComplain.HostelID,
		updatedComplain.ComplainType,
		updatedComplain.Description,
		complainID,
	}
	SQL := `UPDATE complains
SET hostel_id     = $1,
    complain_type = $2,
    description   =$3,
    updated_at    = now()
WHERE id = $4
  AND archived_at IS NULL;`

	_, err := database.DB.ExecContext(ctx, SQL, args...)
	return err
}

func DeleteComplain(ctx context.Context, complainID string) error {
	SQL := `UPDATE complains SET archived_at = now() WHERE
            id = $1 AND archived_at IS NULL;`

	_, err := database.DB.ExecContext(ctx, SQL, complainID)
	return err
}

func GetComplains(ctx context.Context, filters *models.ComplainFilters) ([]models.ComplainInfo, error) {
	args := []interface{}{
		filters.ComplainID,
		filters.StudentID,
		filters.HostelID,
		filters.ComplainType,
		filters.GenericFilters.Limit,
		filters.GenericFilters.Limit * filters.GenericFilters.Page,
	}

	SQL := `SELECT sd.first_name || '' || sd.last_name as student_name,
       h.name as hostel_name,
       ct.name as complain_type,
       c.description ,
       hrd.floor_no as floor,
       hrd.room_no as room
FROM complains c
         JOIN complains_type ct on c.complain_type = ct.id AND ct.archived_at IS NULL
         JOIN student_details sd on c.student_id = sd.id AND sd.archived_at IS NULL
         JOIN hostel_rooms_details hrd on sd.id = hrd.student_id AND hrd.archived_at IS NULL
         JOIN hostels h on c.hostel_id = h.id AND h.archived_at IS NULL
WHERE c.archived_at IS NULL
 AND (c.id::text = '' OR c.id = $1::uuid)
 AND (sd.id::text = '' OR sd.id = $2::uuid)
 AND (h.id::text = '' OR h.id = $3::uuid)
 AND (ct.id::text = '' OR ct.id = $4::uuid)
 ORDER BY c.created_at DESC
 LIMIT $5 OFFSET $6;`

	complains := make([]models.ComplainInfo, 0)
	err := database.DB.SelectContext(ctx, &complains, SQL, args...)
	return complains, err
}
