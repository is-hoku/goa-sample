package interactor

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	datastore "github.com/is-hoku/goa-sample/webapi/datastore/mock"
	"github.com/is-hoku/goa-sample/webapi/gen/student"
	"github.com/is-hoku/goa-sample/webapi/model"
	"github.com/is-hoku/goa-sample/webapi/repository"
	"github.com/is-hoku/goa-sample/webapi/usecase"
)

func TestGetByNumber(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	t.Run("学籍番号に対応する学生を返却", func(t *testing.T) {
		opt := datastore.GetStudentByNumber(func(ctx context.Context, input *repository.GetStudentByNumberInput) (*repository.GetStudentByNumberOutput, error) {
			want := &repository.GetStudentByNumberInput{
				StudentNumber: uint32(12345),
			}
			if diff := cmp.Diff(want, input); diff != "" {
				t.Errorf("unexpected input (-want +got):\n%s", diff)
			}
			date1 := time.Date(2003, 6, 27, 15, 0, 0, 0, time.UTC)
			edate1 := time.Date(2024, 3, 31, 0, 0, 0, 0, time.UTC)
			return &repository.GetStudentByNumberOutput{
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
		want := &usecase.GetStudentByNumberOutput{
			Student: &model.Student{
				ID:             10001,
				Name:           "太郎",
				Ruby:           "たろう",
				StudentNumber:  12345,
				DateOfBirth:    date1,
				Address:        "愛知県名古屋市中区三の丸三丁目1番2号",
				ExpirationDate: edate1,
			},
		}
		getter := NewGetStudentByNumber(&GetStudentByNumberOption{
			opt,
		})
		got, err := getter.GetStudentByNumber(ctx, &usecase.GetStudentByNumberInput{
			StudentNumber: 12345,
		})
		if err != nil {
			t.Errorf("unexpected error:%s", err)
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("unexpected response (-want +got):\n%s", diff)
		}
	})

	t.Run("存在しない学籍番号", func(t *testing.T) {
		opt := datastore.GetStudentByNumber(func(ctx context.Context, input *repository.GetStudentByNumberInput) (*repository.GetStudentByNumberOutput, error) {
			want := &repository.GetStudentByNumberInput{
				StudentNumber: uint32(12345),
			}
			if diff := cmp.Diff(want, input); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
			return nil, &student.CustomError{Name: "not_found", Message: "Student Not Found"}
		})

		wantedError := &student.CustomError{Name: "not_found", Message: "Student Not Found"}

		getter := NewGetStudentByNumber(&GetStudentByNumberOption{
			opt,
		})
		_, err := getter.GetStudentByNumber(ctx, &usecase.GetStudentByNumberInput{
			StudentNumber: 12345,
		})
		if err.(*student.CustomError).Name != wantedError.Name {
			t.Errorf("error response mismatch:\nwant: %v\ngot: %v", wantedError.Name, err.(*student.CustomError).Name)
		}
	})
}

func TestCreate(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	t.Run("学生を登録", func(t *testing.T) {
		opt1 := datastore.GetStudentByID(func(ctx context.Context, input *repository.GetStudentByIDInput) (*repository.GetStudentByIDOutput, error) {
			want := &repository.GetStudentByIDInput{
				ID: uint64(10001),
			}
			if diff := cmp.Diff(want, input); diff != "" {
				t.Errorf("unexpected input (-want +got):\n%s", diff)
			}
			date1 := time.Date(2003, 6, 27, 15, 0, 0, 0, time.UTC)
			edate1 := time.Date(2024, 3, 31, 0, 0, 0, 0, time.UTC)
			return &repository.GetStudentByIDOutput{
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
		opt2 := datastore.CreateStudent(func(ctx context.Context, input *repository.CreateStudentInput) (*repository.CreateStudentOutput, error) {
			date1 := time.Date(2003, 6, 27, 15, 0, 0, 0, time.UTC)
			edate1 := time.Date(2024, 3, 31, 0, 0, 0, 0, time.UTC)
			want := &repository.CreateStudentInput{
				Student: &model.Student{
					ID:             10001,
					Name:           "太郎",
					Ruby:           "たろう",
					StudentNumber:  12345,
					DateOfBirth:    date1,
					Address:        "愛知県名古屋市中区三の丸三丁目1番2号",
					ExpirationDate: edate1,
				},
			}
			if diff := cmp.Diff(want, input); diff != "" {
				t.Errorf("unexpected input (-want +got):\n%s", diff)
			}
			return &repository.CreateStudentOutput{
				ID: uint64(10001),
			}, nil
		})

		date1 := time.Date(2003, 6, 27, 15, 0, 0, 0, time.UTC)
		edate1 := time.Date(2024, 3, 31, 0, 0, 0, 0, time.UTC)
		want := &usecase.CreateStudentOutput{
			Student: &model.Student{
				ID:             10001,
				Name:           "太郎",
				Ruby:           "たろう",
				StudentNumber:  12345,
				DateOfBirth:    date1,
				Address:        "愛知県名古屋市中区三の丸三丁目1番2号",
				ExpirationDate: edate1,
			},
		}
		creator := NewCreateStudent(&CreateStudentOption{
			opt1,
			opt2,
		})
		got, err := creator.CreateStudent(ctx, &usecase.CreateStudentInput{
			Student: &model.Student{
				ID:             10001,
				Name:           "太郎",
				Ruby:           "たろう",
				StudentNumber:  12345,
				DateOfBirth:    date1,
				Address:        "愛知県名古屋市中区三の丸三丁目1番2号",
				ExpirationDate: edate1,
			},
		})
		if err != nil {
			t.Fatalf("unexpected error:%s", err)
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("unexpected response (-want +got):\n%s", diff)
		}
	})
}

func TestGetAll(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	t.Run("全ての学生を返却", func(t *testing.T) {
		opt := datastore.GetStudents(func(ctx context.Context) (*repository.GetStudentsOutput, error) {
			date1 := time.Date(2003, 6, 27, 15, 0, 0, 0, time.UTC)
			date2 := time.Date(2004, 11, 4, 19, 0, 0, 0, time.UTC)
			edate1 := time.Date(2024, 3, 31, 0, 0, 0, 0, time.UTC)
			edate2 := time.Date(2025, 3, 31, 0, 0, 0, 0, time.UTC)
			return &repository.GetStudentsOutput{
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
		want := &usecase.GetStudentsOutput{
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
		}
		getter := NewGetStudents(&GetStudentsOption{
			opt,
		})
		got, err := getter.GetStudents(ctx)
		if err != nil {
			t.Fatalf("Could not get mock data: %s\n", err)
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("student mismatch (-want +got):\n%s", diff)
		}
	})
}
