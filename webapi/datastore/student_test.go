package datastore_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/is-hoku/goa-sample/webapi/datastore"
	"github.com/is-hoku/goa-sample/webapi/gen/student"
	"github.com/is-hoku/goa-sample/webapi/model"
	"github.com/is-hoku/goa-sample/webapi/testutil"
)

func TestGetByNumber(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	handler, err := testutil.NewTestDBHandler(ctx)
	if err != nil {
		t.Fatalf("Could not generate new test db handler: %s\n", err)
	}
	defer handler.Close()
	studentHandler := &datastore.StudentHandler{handler}
	defer testutil.DeleteTestDB(ctx, handler)
	if err := testutil.CreateTestStudentTable(ctx, handler); err != nil {
		t.Fatalf("Could not create test table: %s\n", err)
	}

	t.Run("学籍番号に対応する学生を返却", func(t *testing.T) {
		if err := testutil.TruncateAll(ctx, handler); err != nil {
			t.Fatalf("Could not remove test data from test db: %s\n", err)
		}
		date1 := time.Date(2003, 6, 27, 15, 0, 0, 0, time.UTC)
		date2 := time.Date(2004, 11, 4, 19, 0, 0, 0, time.UTC)
		edate1 := time.Date(2024, 3, 31, 0, 0, 0, 0, time.UTC)
		edate2 := time.Date(2025, 3, 31, 0, 0, 0, 0, time.UTC)
		inserted := []*datastore.SetStudentParams{
			{
				Name:           "太郎",
				Ruby:           "たろう",
				StudentNumber:  12345,
				DateOfBirth:    date1,
				Address:        "愛知県名古屋市中区三の丸三丁目1番2号",
				ExpirationDate: edate1,
			},
			{
				Name:           "次郎",
				Ruby:           "じろう",
				StudentNumber:  12346,
				DateOfBirth:    date2,
				Address:        "東京都新宿区西新宿2-8-1",
				ExpirationDate: edate2,
			},
		}
		if err := testutil.InsertTestStudentsData(ctx, handler, inserted); err != nil {
			t.Fatalf("Could not insert test students data: %s\n", err)
		}
		want := &model.Student{
			ID:             1,
			Name:           "太郎",
			Ruby:           "たろう",
			StudentNumber:  12345,
			DateOfBirth:    date1,
			Address:        "愛知県名古屋市中区三の丸三丁目1番2号",
			ExpirationDate: edate1,
		}

		got, err := studentHandler.GetByNumber(ctx, 12345)
		if err != nil {
			t.Fatalf("Could not get test data from test db: %s\n", err)
		}
		option := cmpopts.IgnoreFields(datastore.Student{}, "CreatedAt", "UpdatedAt")
		if diff := cmp.Diff(want, got, option); diff != "" {
			t.Errorf("student mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("存在しない学籍番号の場合はエラーとして not_found を返す", func(t *testing.T) {
		if err := testutil.TruncateAll(ctx, handler); err != nil {
			t.Fatalf("Could not remove test data from test db: %s\n", err)
		}

		_, err := studentHandler.GetByNumber(ctx, 12345)
		fmt.Println(err.(*student.CustomError))
		wantedError := &student.CustomError{Name: "not_found", Message: "Student Not Found"}
		if err.(*student.CustomError).Name != wantedError.Name {
			t.Errorf("error response mismatch:\nwant: %v\ngot: %v", wantedError.Name, err.(*student.CustomError).Name)
		}
	})
}

func TestGetByID(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	handler, err := testutil.NewTestDBHandler(ctx)
	if err != nil {
		t.Fatalf("Could not generate new test db handler: %s\n", err)
	}
	defer handler.Close()
	studentHandler := &datastore.StudentHandler{handler}
	defer testutil.DeleteTestDB(ctx, handler)
	if err := testutil.CreateTestStudentTable(ctx, handler); err != nil {
		t.Fatalf("Could not create test table: %s\n", err)
	}

	t.Run("ID に対応する学生を返却", func(t *testing.T) {
		if err := testutil.TruncateAll(ctx, handler); err != nil {
			t.Fatalf("Could not remove test data from test db: %s\n", err)
		}
		date1 := time.Date(2003, 6, 27, 15, 0, 0, 0, time.UTC)
		date2 := time.Date(2004, 11, 4, 19, 0, 0, 0, time.UTC)
		edate1 := time.Date(2024, 3, 31, 0, 0, 0, 0, time.UTC)
		edate2 := time.Date(2025, 3, 31, 0, 0, 0, 0, time.UTC)
		inserted := []*datastore.SetStudentParams{
			{
				Name:           "太郎",
				Ruby:           "たろう",
				StudentNumber:  12345,
				DateOfBirth:    date1,
				Address:        "愛知県名古屋市中区三の丸三丁目1番2号",
				ExpirationDate: edate1,
			},
			{
				Name:           "次郎",
				Ruby:           "じろう",
				StudentNumber:  12346,
				DateOfBirth:    date2,
				Address:        "東京都新宿区西新宿2-8-1",
				ExpirationDate: edate2,
			},
		}
		if err := testutil.InsertTestStudentsData(ctx, handler, inserted); err != nil {
			t.Fatalf("Could not insert test students data: %s\n", err)
		}
		want := &model.Student{
			ID:             1,
			Name:           "太郎",
			Ruby:           "たろう",
			StudentNumber:  12345,
			DateOfBirth:    date1,
			Address:        "愛知県名古屋市中区三の丸三丁目1番2号",
			ExpirationDate: edate1,
		}

		got, err := studentHandler.GetByID(ctx, 1)
		if err != nil {
			t.Fatalf("Could not get test data from test db: %s\n", err)
		}
		option := cmpopts.IgnoreFields(datastore.Student{}, "CreatedAt", "UpdatedAt")
		if diff := cmp.Diff(want, got, option); diff != "" {
			t.Errorf("student mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("存在しない ID の場合はエラーとして not_found を返す", func(t *testing.T) {
		if err := testutil.TruncateAll(ctx, handler); err != nil {
			t.Fatalf("Could not remove test data from test db: %s\n", err)
		}

		_, err := studentHandler.GetByID(ctx, 1)
		fmt.Println(err.(*student.CustomError))
		wantedError := &student.CustomError{Name: "not_found", Message: "Student Not Found"}
		if err.(*student.CustomError).Name != wantedError.Name {
			t.Errorf("error response mismatch:\nwant: %v\ngot: %v", wantedError.Name, err.(*student.CustomError).Name)
		}
	})
}

func TestGetAll(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	handler, err := testutil.NewTestDBHandler(ctx)
	if err != nil {
		t.Fatalf("Could not generate new test db handler: %s\n", err)
	}
	studentHandler := &datastore.StudentHandler{handler}
	defer handler.Close()
	defer testutil.DeleteTestDB(ctx, handler)
	if err := testutil.CreateTestStudentTable(ctx, handler); err != nil {
		t.Fatalf("Could not create test table: %s\n", err)
	}

	t.Run("全ての学生を返却 (学籍番号で昇順にソート)", func(t *testing.T) {
		if err := testutil.TruncateAll(ctx, handler); err != nil {
			t.Fatalf("Could not remove test data from test db: %s\n", err)
		}
		date1 := time.Date(2003, 6, 27, 15, 0, 0, 0, time.UTC)
		date2 := time.Date(2004, 11, 4, 19, 0, 0, 0, time.UTC)
		edate1 := time.Date(2024, 3, 31, 0, 0, 0, 0, time.UTC)
		edate2 := time.Date(2025, 3, 31, 0, 0, 0, 0, time.UTC)
		inserted := []*datastore.SetStudentParams{
			{
				Name:           "太郎",
				Ruby:           "たろう",
				StudentNumber:  12345,
				DateOfBirth:    date1,
				Address:        "愛知県名古屋市中区三の丸三丁目1番2号",
				ExpirationDate: edate1,
			},
			{
				Name:           "次郎",
				Ruby:           "じろう",
				StudentNumber:  12346,
				DateOfBirth:    date2,
				Address:        "東京都新宿区西新宿2-8-1",
				ExpirationDate: edate2,
			},
		}
		testutil.InsertTestStudentsData(ctx, handler, inserted)
		want := []*model.Student{
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
				StudentNumber:  12346,
				DateOfBirth:    date2,
				Address:        "東京都新宿区西新宿2-8-1",
				ExpirationDate: edate2,
			},
		}

		got, err := studentHandler.GetAll(ctx)
		if err != nil {
			t.Fatalf("Could not get test data from test db: %s\n", err)
		}
		option := cmpopts.IgnoreFields(datastore.Student{}, "CreatedAt", "UpdatedAt")
		if diff := cmp.Diff(want, got, option); diff != "" {
			t.Errorf("student mismatch (-want +got):\n%s", diff)
		}
	})
}

func TestSet(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	handler, err := testutil.NewTestDBHandler(ctx)
	if err != nil {
		t.Fatalf("Could not generate new test db handler: %s\n", err)
	}
	studentHandler := &datastore.StudentHandler{handler}
	defer handler.Close()
	defer testutil.DeleteTestDB(ctx, handler)
	if err := testutil.CreateTestStudentTable(ctx, handler); err != nil {
		t.Fatalf("Could not create test table: %s\n", err)
	}

	t.Run("学生を登録", func(t *testing.T) {
		if err := testutil.TruncateAll(ctx, handler); err != nil {
			t.Fatalf("Could not remove test data from test db: %s\n", err)
		}
		date1 := time.Date(2003, 6, 27, 15, 0, 0, 0, time.UTC)
		edate1 := time.Date(2024, 3, 31, 0, 0, 0, 0, time.UTC)
		inserted := &model.Student{
			Name:           "太郎",
			Ruby:           "たろう",
			StudentNumber:  12345,
			DateOfBirth:    date1,
			Address:        "愛知県名古屋市中区三の丸三丁目1番2号",
			ExpirationDate: edate1,
		}
		want := &model.Student{
			ID:             1,
			Name:           "太郎",
			Ruby:           "たろう",
			StudentNumber:  12345,
			DateOfBirth:    date1,
			Address:        "愛知県名古屋市中区三の丸三丁目1番2号",
			ExpirationDate: edate1,
		}

		gotID, err := studentHandler.Set(ctx, inserted)
		if err != nil {
			t.Fatalf("Could not set test data to test db: %s\n", err)
		}
		got, err := studentHandler.GetByID(ctx, gotID)
		if err != nil {
			t.Fatalf("Could not get test data from test db: %s\n", err)
		}
		option := cmpopts.IgnoreFields(datastore.Student{}, "CreatedAt", "UpdatedAt")
		if diff := cmp.Diff(want, got, option); diff != "" {
			t.Errorf("student mismatch (-want +got):\n%s", diff)
		}
	})
}
