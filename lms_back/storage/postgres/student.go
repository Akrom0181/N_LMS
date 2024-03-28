package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"lms_back/api/models"
	"lms_back/pkg"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type StudentRepo struct {
	db *pgxpool.Pool
}

func NewStudent(db *pgxpool.Pool) StudentRepo {
	return StudentRepo{
		db: db,
	}
}

func (c *StudentRepo) Create(student models.Student) (models.Student, error) {

	id := uuid.New()
	query := `INSERT INTO student (
		id,
		full_name,
		email,
		age,
		paid_sum,
		status,
		login,
		password,
		created_at)
		VALUES($1,$2,$3,$4,$5,$6,$7,$8,CURRENT_TIMESTAMP) 
	`

	_, err := c.db.Exec(context.Background(), query,
		id.String(),
		student.Full_Name,
		student.Email,
		student.Age,
		student.PaidSum,
		student.Status,
		student.Login,
		student.Password,
	)

	if err != nil {
		return models.Student{}, err
	}
	return models.Student{
		ID:         student.ID,
		Full_Name:  student.Full_Name,
		Email:      student.Email,
		Age:        student.Age,
		PaidSum:    student.PaidSum,
		Login:      student.Login,
		Password:   student.Password,
		Created_At: student.Created_At,
		Updated_At: student.Updated_At,
	}, nil
}

func (c *StudentRepo) Update(student models.Student) (models.Student, error) {
	query := `update student set 
		full_name=$1,
		email=$2,
		age=$3,
		paid_sum=$4,
		login=$5,
		password=$6,
		group_id=$7,
		status=$8,
		updated_at = CURRENT_TIMESTAMP
		WHERE id = $9
	`
	_, err := c.db.Exec(context.Background(), query,
		student.Full_Name,
		student.Email,
		student.Age,
		student.PaidSum,
		student.Login,
		student.Password,
		student.GroupID,
		student.Status,
		student.ID,
	)
	if err != nil {
		return models.Student{}, err
	}
	return models.Student{
		ID:         student.ID,
		Full_Name:  student.Full_Name,
		Email:      student.Email,
		Age:        student.Age,
		PaidSum:    student.PaidSum,
		Login:      student.Login,
		Password:   student.Password,
		Created_At: student.Created_At,
		Updated_At: student.Updated_At,
	}, nil
}

func (c *StudentRepo) GetAll(req models.GetAllStudentsRequest) (models.GetAllStudentsResponse, error) {
	var (
		resp   = models.GetAllStudentsResponse{}
		filter = ""
	)
	offset := (req.Page - 1) * req.Limit

	if req.Search != "" {
		filter += fmt.Sprintf(` and name ILIKE  '%%%v%%' `, req.Search)
	}

	filter += fmt.Sprintf(" OFFSET %v LIMIT %v", offset, req.Limit)
	fmt.Println("filter: ", filter)

	rows, err := c.db.Query(context.Background(), `select count(id) over(),
        id,
        full_name,
		email,
		age,
		paid_sum,
		status,
        login,
		password,
		group_id,
        created_at,
        updated_at FROM student`+filter+``)
	if err != nil {
		return resp, err
	}

	for rows.Next() {
		var (
			student    = models.Student{}
			full_name  sql.NullString
			email      sql.NullString
			age        sql.NullInt16
			paid_sum   sql.NullInt64
			status     sql.NullString
			login      sql.NullString
			password   sql.NullString
			group_id   sql.NullString
			created_at sql.NullString
			updated_at sql.NullString
		)
		if err := rows.Scan(
			&resp.Count,
			&student.ID,
			&full_name,
			&email,
			&age,
			&paid_sum,
			&status,
			&login,
			&password,
			&group_id,
			&created_at,
			&updated_at); err != nil {
			return resp, err
		}
		student.Updated_At = pkg.NullStringToString(updated_at)
		resp.Students = append(resp.Students, models.Student{
			Full_Name:  full_name.String,
			Email:      email.String,
			Age:        int(age.Int16),
			PaidSum:    float64(paid_sum.Int64),
			Status:     status.String,
			Login:      login.String,
			Password:   password.String,
			GroupID:    group_id.String,
			Created_At: created_at.String,
			Updated_At: updated_at.String,
		})
	}
	return resp, nil
}

func (c *StudentRepo) GetByID(id string) (models.Student, error) {
	var (
		student    = models.Student{}
		full_name  sql.NullString
		email      sql.NullString
		age        sql.NullInt16
		paid_sum   sql.NullInt64
		status     sql.NullString
		login      sql.NullString
		password   sql.NullString
		group_id   sql.NullString
		created_at sql.NullString
		updated_at sql.NullString
	)
	if err := c.db.QueryRow(context.Background(), `select id, full_name, email, age, paid_sum, status, login, password, group_id, created_at, updated_at from student where id = $1`, id).Scan(
		&student.ID,
		&full_name,
		&email,
		&age,
		&paid_sum,
		&status,
		&login,
		&password,
		&group_id,
		&created_at,
		&updated_at,
	); err != nil {
		return models.Student{}, err
	}
	return models.Student{
		Full_Name:  full_name.String,
		Email:      email.String,
		Age:        int(age.Int16),
		PaidSum:    float64(paid_sum.Int64),
		Status:     status.String,
		Login:      login.String,
		Password:   password.String,
		GroupID:    group_id.String,
		Created_At: created_at.String,
		Updated_At: updated_at.String,
	}, nil
}

func (c *StudentRepo) Delete(id string) error {
	query := `delete from student where id = $1`
	_, err := c.db.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	return nil
}
