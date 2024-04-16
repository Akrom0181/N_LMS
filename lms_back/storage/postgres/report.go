package postgres

import (
	"context"
	"time"

	"lms_back/api/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type adminReportRepo struct {
	db *pgxpool.Pool
}

func AdminPayment(db *pgxpool.Pool) adminReportRepo {
	return adminReportRepo{
		db: db,
	}
}

func (c *adminReportRepo) GetByIDAdminPayment(ctx context.Context, req models.AdminKey) ([]models.AdminPayment, error) {
	var adminPayments []models.AdminPayment

	rows, err := c.db.Query(ctx, `SELECT 
        a.id AS admin_id,
        a.full_name AS admin_fullname,
        p.id AS payment_id,
        p.price,
        p.student_id,
        p.branch_id,
        p.created_at::date AS payment_created_at,
        p.updated_at AS payment_updated_at
    FROM 
        admin a
    JOIN 
        payment p ON a.id = p.admin_id
    WHERE 
        a.id = $1`, req.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var admin models.Admin
		var payment models.Payment
		var CreatedAt time.Time
		var Updated_At time.Time

		if err := rows.Scan(
			&admin.Id,
			&admin.Full_Name,
			&payment.Id,
			&payment.Price,
			&payment.Student_id,
			&payment.Branch_id,
			&CreatedAt,
			&Updated_At,
		); err != nil {
			return nil, err
		}

		adminPayment := models.AdminPayment{
			Admin:   admin,
			Payment: payment,
		}
		adminPayments = append(adminPayments, adminPayment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return adminPayments, nil
}
