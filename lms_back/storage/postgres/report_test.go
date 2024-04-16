package postgres

import (
	"context"
	"lms_back/api/models"
	"reflect"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Test_adminReportRepo_GetByIDAdminPayment(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		ctx context.Context
		req models.AdminKey
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.AdminPayment
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &adminReportRepo{
				db: tt.fields.db,
			}
			got, err := c.GetByIDAdminPayment(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("adminReportRepo.GetByIDAdminPayment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("adminReportRepo.GetByIDAdminPayment() = %v, want %v", got, tt.want)
			}
		})
	}
}
