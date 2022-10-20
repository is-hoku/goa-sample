package sample

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/is-hoku/goa-sample/webapi/gen/student"
	interactor "github.com/is-hoku/goa-sample/webapi/interactor/mock"
	"github.com/is-hoku/goa-sample/webapi/model"
	"github.com/is-hoku/goa-sample/webapi/usecase"
)

func TestGetStudent(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	t.Run("学籍番号に対応する学生を返却", func(t *testing.T) {
		getter := interactor.GetStudentByNumber(func(ctx context.Context, input *usecase.GetStudentByNumberInput) (*usecase.GetStudentByNumberOutput, error) {
			date1 := time.Date(2003, 6, 27, 15, 0, 0, 0, time.UTC)
			edate1 := time.Date(2024, 3, 31, 0, 0, 0, 0, time.UTC)
			return &usecase.GetStudentByNumberOutput{
				Student: &model.Student{
					ID:             10001,
					Name:           "太郎",
					Ruby:           "たろう",
					StudentNumber:  12345,
					DateOfBirth:    date1,
					Address:        "愛知県名古屋市中区三の丸三丁目1番2号",
					ExpirationDate: edate1,
				},
			}, nil
		})

		date1 := time.Date(2003, 6, 27, 15, 0, 0, 0, time.UTC)
		edate1 := time.Date(2024, 3, 31, 0, 0, 0, 0, time.UTC)
		want := &student.Student{
			ID:             10001,
			Name:           "太郎",
			Ruby:           "たろう",
			StudentNumber:  12345,
			DateOfBirth:    date1.Format(time.RFC3339),
			Address:        "愛知県名古屋市中区三の丸三丁目1番2号",
			ExpirationDate: edate1.Format(time.RFC3339),
		}
		studentNumber := uint32(12345)
		payload := &student.GetStudentPayload{StudentNumber: &studentNumber}

		logger := log.New(os.Stderr, "[test] ", log.Ltime)
		studentApp, err := newStudentApp(ctx, getter, nil, nil)
		if err != nil {
			t.Errorf("unexpected error")
		}
		srv := NewStudent(logger, studentApp)
		got, err := srv.GetStudent(ctx, payload)
		if err != nil {
			t.Errorf("unexpected error")
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("unexpected response (-want +got):\n%s", diff)
		}
	})

	t.Run("存在しない学籍番号", func(t *testing.T) {
		getter := interactor.GetStudentByNumber(func(ctx context.Context, input *usecase.GetStudentByNumberInput) (*usecase.GetStudentByNumberOutput, error) {
			return nil, &student.CustomError{Name: "not_found", Message: "Student Not Found"}
		})

		wantedError := &student.CustomError{Name: "not_found", Message: "Student Not Found"}

		studentNumber := uint32(12346)
		payload := &student.GetStudentPayload{StudentNumber: &studentNumber}

		logger := log.New(os.Stderr, "[test] ", log.Ltime)
		studentApp, err := newStudentApp(ctx, getter, nil, nil)
		if err != nil {
			t.Errorf("unexpected error")
		}
		srv := NewStudent(logger, studentApp)
		_, err = srv.GetStudent(ctx, payload)
		if err.(*student.CustomError).Name != wantedError.Name {
			t.Errorf("unexpected response:\nwant: %v\ngot: %v", wantedError.Name, err.(*student.CustomError).Name)
		}
	})
}

func TestGetStudents(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	t.Run("学生を登録", func(t *testing.T) {
		multigetter := interactor.GetStudents(func(ctx context.Context) (*usecase.GetStudentsOutput, error) {
			date1 := time.Date(2003, 6, 27, 15, 0, 0, 0, time.UTC)
			date2 := time.Date(2004, 11, 4, 19, 0, 0, 0, time.UTC)
			edate1 := time.Date(2024, 3, 31, 0, 0, 0, 0, time.UTC)
			edate2 := time.Date(2025, 3, 31, 0, 0, 0, 0, time.UTC)
			return &usecase.GetStudentsOutput{
				Students: []*model.Student{
					{
						ID:             10001,
						Name:           "太郎",
						Ruby:           "たろう",
						StudentNumber:  12345,
						DateOfBirth:    date1,
						Address:        "愛知県名古屋市中区三の丸三丁目1番2号",
						ExpirationDate: edate1,
					},
					{
						ID:             10002,
						Name:           "次郎",
						Ruby:           "じろう",
						StudentNumber:  12346,
						DateOfBirth:    date2,
						Address:        "東京都新宿区西新宿2-8-1",
						ExpirationDate: edate2,
					},
				},
			}, nil
		})

		date1 := time.Date(2003, 6, 27, 15, 0, 0, 0, time.UTC)
		date2 := time.Date(2004, 11, 4, 19, 0, 0, 0, time.UTC)
		edate1 := time.Date(2024, 3, 31, 0, 0, 0, 0, time.UTC)
		edate2 := time.Date(2025, 3, 31, 0, 0, 0, 0, time.UTC)
		want := &student.Students{
			Students: []*student.Student{
				{
					ID:             10001,
					Name:           "太郎",
					Ruby:           "たろう",
					StudentNumber:  12345,
					DateOfBirth:    date1.Format(time.RFC3339),
					Address:        "愛知県名古屋市中区三の丸三丁目1番2号",
					ExpirationDate: edate1.Format(time.RFC3339),
				},
				{
					ID:             10002,
					Name:           "次郎",
					Ruby:           "じろう",
					StudentNumber:  12346,
					DateOfBirth:    date2.Format(time.RFC3339),
					Address:        "東京都新宿区西新宿2-8-1",
					ExpirationDate: edate2.Format(time.RFC3339),
				},
			},
		}

		logger := log.New(os.Stderr, "[test] ", log.Ltime)
		studentApp, err := newStudentApp(ctx, nil, multigetter, nil)
		if err != nil {
			t.Errorf("unexpected error")
		}
		srv := NewStudent(logger, studentApp)
		got, err := srv.GetStudents(ctx, &student.GetStudentsPayload{})
		if err != nil {
			t.Fatalf("unexpected error: %s\n", err)
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("unexpected response (-want +got):\n%s", diff)
		}
	})
}

func TestCreate(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	t.Run("学生を登録", func(t *testing.T) {
		creator := interactor.CreateStudent(func(ctx context.Context, input *usecase.CreateStudentInput) (*usecase.CreateStudentOutput, error) {
			date1 := time.Date(2003, 6, 27, 15, 0, 0, 0, time.UTC)
			edate1 := time.Date(2024, 3, 31, 0, 0, 0, 0, time.UTC)
			return &usecase.CreateStudentOutput{
				Student: &model.Student{
					ID:             10001,
					Name:           "太郎",
					Ruby:           "たろう",
					StudentNumber:  12345,
					DateOfBirth:    date1,
					Address:        "愛知県名古屋市中区三の丸三丁目1番2号",
					ExpirationDate: edate1,
				},
			}, nil
		})

		date1 := time.Date(2003, 6, 27, 15, 0, 0, 0, time.UTC)
		edate1 := time.Date(2024, 3, 31, 0, 0, 0, 0, time.UTC)
		want := &student.Student{
			ID:             10001,
			Name:           "太郎",
			Ruby:           "たろう",
			StudentNumber:  12345,
			DateOfBirth:    date1.Format(time.RFC3339),
			Address:        "愛知県名古屋市中区三の丸三丁目1番2号",
			ExpirationDate: edate1.Format(time.RFC3339),
		}
		body := &student.CreateStudentPayload{
			Name:           "太郎",
			Ruby:           "たろう",
			StudentNumber:  12345,
			DateOfBirth:    date1.Format(time.RFC3339),
			Address:        "愛知県名古屋市中区三の丸三丁目1番2号",
			ExpirationDate: edate1.Format(time.RFC3339),
		}

		logger := log.New(os.Stderr, "[test] ", log.Ltime)
		studentApp, err := newStudentApp(ctx, nil, nil, creator)
		if err != nil {
			t.Errorf("unexpected error")
		}
		srv := NewStudent(logger, studentApp)
		got, err := srv.CreateStudent(ctx, body)
		if err != nil {
			t.Errorf("unexpected error")
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("unexpected response (-want +got):\n%s", diff)
		}
	})
}
