package datastore

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/is-hoku/goa-sample/webapi/gen/student"
	"github.com/is-hoku/goa-sample/webapi/model"
	"github.com/is-hoku/goa-sample/webapi/repository"
)

func TestGetByNumber(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	testDB, err := newTestDB(ctx)
	if err != nil {
		t.Errorf("unexpected error: %s\n", err)
	}
	defer testDB.Close()
	defer deleteTestDB(ctx, testDB)
	if err := createTestStudentsTable(ctx, testDB); err != nil {
		t.Errorf("unexpected error: %s\n", err)
	}

	t.Run("学籍番号に対応する学生を返却", func(t *testing.T) {
		if err := truncateAll(ctx, testDB); err != nil {
			t.Fatalf("Could not remove test data from test db: %s\n", err)
		}
		date1 := time.Date(2003, 6, 27, 15, 0, 0, 0, time.UTC)
		edate1 := time.Date(2024, 3, 31, 0, 0, 0, 0, time.UTC)
		createStudentMedia := NewCreateStudentMedia(testDB)
		createStudentMedia.CreateStudent(ctx, &repository.CreateStudentInput{
			Student: &model.Student{
				ID:             1,
				Name:           "太郎",
				Ruby:           "たろう",
				StudentNumber:  12345,
				DateOfBirth:    date1,
				Address:        "愛知県名古屋市中区三の丸三丁目1番2号",
				ExpirationDate: edate1,
			},
		})

		want := &repository.GetStudentByNumberOutput{
			Student: &model.Student{
				ID:             1,
				Name:           "太郎",
				Ruby:           "たろう",
				StudentNumber:  12345,
				DateOfBirth:    date1,
				Address:        "愛知県名古屋市中区三の丸三丁目1番2号",
				ExpirationDate: edate1,
			},
		}

		getStudentMedia := NewGetStudentMedia(testDB)
		got, err := getStudentMedia.GetStudentByNumber(ctx, &repository.GetStudentByNumberInput{
			StudentNumber: 12345,
		})
		if err != nil {
			t.Fatalf("Could not get test data from test db: %s\n", err)
		}
		option := cmpopts.IgnoreFields(got, "CreatedAt", "UpdatedAt")
		if diff := cmp.Diff(want, got, option); diff != "" {
			t.Errorf("unexpected response (-want +got):\n%s", diff)
		}
	})

	t.Run("存在しない学籍番号の場合はエラーとして not_found を返す", func(t *testing.T) {
		if err := truncateAll(ctx, testDB); err != nil {
			t.Fatalf("Could not remove test data from test db: %s\n", err)
		}

		getStudentMedia := NewGetStudentMedia(testDB)
		_, err := getStudentMedia.GetStudentByNumber(ctx, &repository.GetStudentByNumberInput{
			StudentNumber: 12345,
		})
		wantedError := &student.CustomError{Name: "not_found", Message: "Student Not Found"}
		if err.(*student.CustomError).Name != wantedError.Name {
			t.Errorf("error response mismatch:\nwant: %v\ngot: %v", wantedError.Name, err.(*student.CustomError).Name)
		}
	})
}

func TestGetByID(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	testDB, err := newTestDB(ctx)
	if err != nil {
		t.Errorf("unexpected error: %s\n", err)
	}
	defer testDB.Close()
	defer deleteTestDB(ctx, testDB)
	if err := createTestStudentsTable(ctx, testDB); err != nil {
		t.Errorf("unexpected error: %s\n", err)
	}

	t.Run("ID に対応する学生を返却", func(t *testing.T) {
		if err := truncateAll(ctx, testDB); err != nil {
			t.Fatalf("Could not remove test data from test db: %s\n", err)
		}
		date1 := time.Date(2003, 6, 27, 15, 0, 0, 0, time.UTC)
		edate1 := time.Date(2024, 3, 31, 0, 0, 0, 0, time.UTC)
		createStudentMedia := NewCreateStudentMedia(testDB)
		createStudentMedia.CreateStudent(ctx, &repository.CreateStudentInput{
			Student: &model.Student{
				Name:           "太郎",
				Ruby:           "たろう",
				StudentNumber:  12345,
				DateOfBirth:    date1,
				Address:        "愛知県名古屋市中区三の丸三丁目1番2号",
				ExpirationDate: edate1,
			},
		})

		want := &repository.GetStudentByIDOutput{
			Student: &model.Student{
				ID:             1,
				Name:           "太郎",
				Ruby:           "たろう",
				StudentNumber:  12345,
				DateOfBirth:    date1,
				Address:        "愛知県名古屋市中区三の丸三丁目1番2号",
				ExpirationDate: edate1,
			},
		}

		gotStudentMedia := NewGetStudentByIDMedia(testDB)
		got, err := gotStudentMedia.GetStudentByID(ctx, &repository.GetStudentByIDInput{
			ID: 1,
		})
		if err != nil {
			t.Fatalf("Could not get test data from test db: %s\n", err)
		}
		option := cmpopts.IgnoreFields(got, "CreatedAt", "UpdatedAt")
		if diff := cmp.Diff(want, got, option); diff != "" {
			t.Errorf("unexpected response (-want +got):\n%s", diff)
		}
	})

	t.Run("存在しない ID の場合はエラーとして not_found を返す", func(t *testing.T) {
		if err := truncateAll(ctx, testDB); err != nil {
			t.Fatalf("Could not remove test data from test db: %s\n", err)
		}

		gotStudentMedia := NewGetStudentByIDMedia(testDB)
		_, err := gotStudentMedia.GetStudentByID(ctx, &repository.GetStudentByIDInput{
			ID: 1,
		})
		wantedError := &student.CustomError{Name: "not_found", Message: "Student Not Found"}
		if err.(*student.CustomError).Name != wantedError.Name {
			t.Errorf("error response mismatch:\nwant: %v\ngot: %v", wantedError.Name, err.(*student.CustomError).Name)
		}
	})
}

func TestGetAll(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	testDB, err := newTestDB(ctx)
	if err != nil {
		t.Errorf("unexpected error: %s\n", err)
	}
	defer testDB.Close()
	defer deleteTestDB(ctx, testDB)
	if err := createTestStudentsTable(ctx, testDB); err != nil {
		t.Errorf("unexpected error: %s\n", err)
	}

	t.Run("全ての学生を返却 (学籍番号で昇順にソート)", func(t *testing.T) {
		if err := truncateAll(ctx, testDB); err != nil {
			t.Fatalf("Could not remove test data from test db: %s\n", err)
		}
		date1 := time.Date(2003, 6, 27, 15, 0, 0, 0, time.UTC)
		edate1 := time.Date(2024, 3, 31, 0, 0, 0, 0, time.UTC)
		createStudentMedia := NewCreateStudentMedia(testDB)
		createStudentMedia.CreateStudent(ctx, &repository.CreateStudentInput{
			Student: &model.Student{
				Name:           "太郎",
				Ruby:           "たろう",
				StudentNumber:  12345,
				DateOfBirth:    date1,
				Address:        "愛知県名古屋市中区三の丸三丁目1番2号",
				ExpirationDate: edate1,
			},
		})
		createStudentMedia.CreateStudent(ctx, &repository.CreateStudentInput{
			Student: &model.Student{
				ID:             2,
				Name:           "次郎",
				Ruby:           "じろう",
				StudentNumber:  67890,
				DateOfBirth:    date1,
				Address:        "愛知県名古屋市中区三の丸三丁目1番2号",
				ExpirationDate: edate1,
			},
		})

		want := &repository.GetStudentsOutput{
			Students: []*model.Student{
				{
					ID:             1,
					Name:           "太郎",
					Ruby:           "たろう",
					StudentNumber:  12345,
					DateOfBirth:    date1,
					Address:        "愛知県名古屋市中区三の丸三丁目1番2号",
					ExpirationDate: edate1,
				},
				{
					ID:             2,
					Name:           "次郎",
					Ruby:           "じろう",
					StudentNumber:  67890,
					DateOfBirth:    date1,
					Address:        "愛知県名古屋市中区三の丸三丁目1番2号",
					ExpirationDate: edate1,
				},
			},
		}

		getStudentsMedia := NewGetStudentsMedia(testDB)
		got, err := getStudentsMedia.GetStudents(ctx)
		if err != nil {
			t.Fatalf("Could not get test data from test db: %s\n", err)
		}
		option := cmpopts.IgnoreFields(got, "CreatedAt", "UpdatedAt")
		if diff := cmp.Diff(want, got, option); diff != "" {
			t.Errorf("unexpected response (-want +got):\n%s", diff)
		}
	})
}

func TestSet(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	testDB, err := newTestDB(ctx)
	if err != nil {
		t.Errorf("unexpected error: %s\n", err)
	}
	defer testDB.Close()
	defer deleteTestDB(ctx, testDB)
	if err := createTestStudentsTable(ctx, testDB); err != nil {
		t.Errorf("unexpected error: %s\n", err)
	}

	t.Run("学生を登録", func(t *testing.T) {
		if err := truncateAll(ctx, testDB); err != nil {
			t.Fatalf("Could not remove test data from test db: %s\n", err)
		}
		date1 := time.Date(2003, 6, 27, 15, 0, 0, 0, time.UTC)
		edate1 := time.Date(2024, 3, 31, 0, 0, 0, 0, time.UTC)

		want := &repository.CreateStudentOutput{
			ID: 1,
		}

		createStudentMedia := NewCreateStudentMedia(testDB)
		gotID, err := createStudentMedia.CreateStudent(ctx, &repository.CreateStudentInput{
			Student: &model.Student{
				Name:           "太郎",
				Ruby:           "たろう",
				StudentNumber:  12345,
				DateOfBirth:    date1,
				Address:        "愛知県名古屋市中区三の丸三丁目1番2号",
				ExpirationDate: edate1,
			},
		})
		if err != nil {
			t.Fatalf("Could not set test data to test db: %s\n", err)
		}

		getStudentByIDMedia := NewGetStudentByIDMedia(testDB)
		got, err := getStudentByIDMedia.GetStudentByID(ctx, &repository.GetStudentByIDInput{
			ID: gotID.ID,
		})
		if err != nil {
			t.Fatalf("Could not get test data from test db: %s\n", err)
		}
		option := cmpopts.IgnoreFields(got, "CreatedAt", "UpdatedAt")
		if diff := cmp.Diff(want, got, option); diff != "" {
			t.Errorf("unexpected response (-want +got):\n%s", diff)
		}
	})
}
