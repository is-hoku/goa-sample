package datastore

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/is-hoku/goa-template/webapi/model"
	"github.com/is-hoku/goa-template/webapi/repository"
)

type DBHandler struct {
	DB *sql.DB
}

type Config struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

func (c Config) DNS() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.DBName,
	)
}

func New(config *Config) (*DBHandler, error) {
	db, err := sql.Open("mysql", config.DNS())
	if err != nil {
		log.Printf("Cannot connect to DB: %s", err)
		return nil, err
	}
	return &DBHandler{db}, err
}

var _ repository.StudentRepository = (*DBHandler)(nil)

func (db *DBHandler) Get(ctx context.Context, number int64) (*model.Student, error) {
	var s model.Student
	row := db.DB.QueryRowContext(ctx,
		"SELECT `id`, `name`, `ruby`, `student_number`, `date_of_birth`, `address`, `expiration_date`, `created_at`, `updated_at` FROM `students` WHERE `student_number`=?",
		number,
	)
	err := row.Scan(&s.ID, &s.Name, &s.Ruby, &s.StudentNumber, &s.DateOfBirth, &s.Address, &s.ExpirationDate, &s.CreatedAt, &s.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (db *DBHandler) Set(ctx context.Context, s *model.Student) (*model.Student, error) {
	result, err := db.DB.ExecContext(ctx,
		"INSERT INTO `students` (`name`, `ruby`, `student_number`, `date_of_birth`, `address`, `expiration_date`) VALUES(?, ?, ?, ?, ?, ?)",
		s.Name, s.Ruby, s.StudentNumber, s.DateOfBirth, s.Address, s.ExpirationDate,
	)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	inserted := &model.Student{
		ID:             id,
		Name:           s.Name,
		Ruby:           s.Ruby,
		StudentNumber:  s.StudentNumber,
		DateOfBirth:    s.DateOfBirth,
		Address:        s.Address,
		ExpirationDate: s.ExpirationDate,
	}
	return inserted, nil
}

func (db *DBHandler) GetAll(ctx context.Context) ([]*model.Student, error) {
	var students []*model.Student
	rows, err := db.DB.QueryContext(ctx,
		"SELECT `id`, `name`, `ruby`, `student_number`, `date_of_birth`, `address`, `expiration_date`, `created_at`, `updated_at` FROM `students` ORDER BY `student_number` ASC",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var s model.Student
		if err := rows.Scan(&s.ID, &s.Name, &s.Ruby, &s.StudentNumber, &s.DateOfBirth, &s.Address, &s.ExpirationDate, &s.CreatedAt, &s.UpdatedAt); err != nil {
			return nil, err
		}
		students = append(students, &s)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return students, err
}

func (db *DBHandler) Close() error {
	return db.DB.Close()
}
