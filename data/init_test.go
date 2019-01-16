package data

import (
	"database/sql"
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jmoiron/sqlx"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

var mock sqlmock.Sqlmock

//main-thread/entry point of go-test
func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	var err error
	var dbx *sql.DB
	dbx, mock, err = sqlmock.New()
	if err != nil {
		panic(err.Error())
	}

	db = sqlx.NewDb(dbx, "postgres")
	os.Exit(m.Run())
}
func TestGetUserByID(t *testing.T) {
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		args    args
		want    User
		wantErr bool
	}{
		{
			name: "tc1-normal",
			args: args{
				id: 1,
			},
			want: User{
				ID:      1,
				Name:    "Patrick",
				Email:   "patrick@bikinibottom.oc",
				Phone:   "-",
				OVOCash: 150000,
			},
			wantErr: false,
		},
		{
			name: "tc2-err",
			args: args{
				id: 2,
			},
			want: User{
				ID:      2,
				Name:    "sandy",
				Email:   "sandy@bikinibottom.oc",
				Phone:   "-",
				OVOCash: 150000,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			rows := sqlmock.NewRows([]string{"id", "name", "email", "phone", "ovo_cash"})
			rows.AddRow(tt.args.id, tt.want.Name, tt.want.Email, tt.want.Phone, tt.want.OVOCash)

			eq := mock.ExpectQuery("^SELECT (.+) FROM ws_user (.+)$").
				WithArgs(tt.args.id).WillReturnRows(rows)

			if tt.wantErr {
				eq.WillReturnError(errors.New("sql err"))
			}

			got, err := GetUserByID(tt.args.id)
			t.Logf("input data %+v", got)
			t.Logf("input param %+v", tt.want)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Error(err)
				return
			}
			assert.Equal(t, err != nil, tt.wantErr, "err false")

			if tt.wantErr {
				return
			}

			assert.Equal(t, tt.want, got, "data failure")
		})
	}
}
