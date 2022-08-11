package interactor

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/is-hoku/goa-sample/webapi/gen/student"
	"github.com/is-hoku/goa-sample/webapi/model"
	"github.com/is-hoku/goa-sample/webapi/repository"
)

type MockStudent struct {
	GetByNumberFunc func(context.Context, int32) (*model.Student, error)
	GetByIDFunc     func(context.Context, int64) (*model.Student, error)
	SetFunc         func(context.Context, *model.Student) (int64, error)
	GetAllFunc      func(context.Context) ([]*model.Student, error)
	repository.TxBeginner
}

func (m *MockStudent) GetByNumber(ctx context.Context, number int32) (*model.Student, error) {
	return m.GetByNumberFunc(ctx, number)
}

func (m *MockStudent) GetByID(ctx context.Context, id int64) (*model.Student, error) {
	return m.GetByIDFunc(ctx, id)
}

func (m *MockStudent) Set(ctx context.Context, student *model.Student) (int64, error) {
	return m.SetFunc(ctx, student)
}

func (m *MockStudent) GetAll(ctx context.Context) ([]*model.Student, error) {
	return m.GetAllFunc(ctx)
}

type mockTx string

func (tx mockTx) Commit() error {
	return nil
}

func (tx mockTx) Rollback() error {
	return nil
}

func (m *MockStudent) BeginTx(ctx context.Context) (repository.Tx, error) {
	var tx mockTx
	return tx, nil
}

var _ repository.StudentRepository = (*MockStudent)(nil)

func TestGetByNumber(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	mock := &MockStudent{}
	i := &StudentInteractor{Repo: mock}
	t.Run("学籍番号に対応する学生を返却", func(t *testing.T) {
		mock.GetByNumberFunc = func(ctx context.Context, number int32) (*model.Student, error) {
			wantedNumber := int32(12345)
			if diff := cmp.Diff(wantedNumber, number); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
			date1 := time.Date(2003, 6, 27, 15, 0, 0, 0, time.UTC)
			edate1 := time.Date(2024, 3, 31, 0, 0, 0, 0, time.UTC)
			s := &model.Student{
				ID:             10001,
				Name:           "太郎",
				Ruby:           "たろう",
				StudentNumber:  12345,
				DateOfBirth:    date1,
				Address:        "愛知県名古屋市中区三の丸三丁目1番2号",
				ExpirationDate: edate1,
			}
			return s, nil
		}
		date1 := time.Date(2003, 6, 27, 15, 0, 0, 0, time.UTC)
		edate1 := time.Date(2024, 3, 31, 0, 0, 0, 0, time.UTC)
		want := &model.Student{
			ID:             10001,
			Name:           "太郎",
			Ruby:           "たろう",
			StudentNumber:  12345,
			DateOfBirth:    date1,
			Address:        "愛知県名古屋市中区三の丸三丁目1番2号",
			ExpirationDate: edate1,
		}
		got, err := i.GetByNumber(ctx, 12345)
		if err != nil {
			t.Fatalf("Could not get mock data: %s\n", err)
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("student mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("存在しない学籍番号", func(t *testing.T) {
		mock.GetByNumberFunc = func(ctx context.Context, number int32) (*model.Student, error) {
			wantedNumber := int32(12346)
			if diff := cmp.Diff(wantedNumber, number); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
			return nil, &student.CustomError{Name: "not_found", Message: "Student Not Found"}
		}
		_, err := i.GetByNumber(ctx, 12346)
		wantedError := &student.CustomError{Name: "not_found", Message: "Student Not Found"}
		if err.(*student.CustomError).Name != wantedError.Name {
			t.Errorf("error response mismatch:\nwant: %v\ngot: %v", wantedError.Name, err.(*student.CustomError).Name)
		}
	})
}

func TestCreate(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	mock := &MockStudent{}
	i := &StudentInteractor{Repo: mock}
	t.Run("学生を登録", func(t *testing.T) {
		mock.GetByIDFunc = func(ctx context.Context, id int64) (*model.Student, error) {
			wantedID := int64(10001)
			if diff := cmp.Diff(wantedID, id); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
			date1 := time.Date(2003, 6, 27, 15, 0, 0, 0, time.UTC)
			edate1 := time.Date(2024, 3, 31, 0, 0, 0, 0, time.UTC)
			s := &model.Student{
				ID:             10001,
				Name:           "太郎",
				Ruby:           "たろう",
				StudentNumber:  12345,
				DateOfBirth:    date1,
				Address:        "愛知県名古屋市中区三の丸三丁目1番2号",
				ExpirationDate: edate1,
			}
			return s, nil
		}
		mock.SetFunc = func(ctx context.Context, gotStudent *model.Student) (int64, error) {
			date1 := time.Date(2003, 6, 27, 15, 0, 0, 0, time.UTC)
			edate1 := time.Date(2024, 3, 31, 0, 0, 0, 0, time.UTC)
			wantedStudent := &model.Student{
				ID:             10001,
				Name:           "太郎",
				Ruby:           "たろう",
				StudentNumber:  12345,
				DateOfBirth:    date1,
				Address:        "愛知県名古屋市中区三の丸三丁目1番2号",
				ExpirationDate: edate1,
			}
			if diff := cmp.Diff(wantedStudent, gotStudent); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
			return 10001, nil
		}
		date1 := time.Date(2003, 6, 27, 15, 0, 0, 0, time.UTC)
		edate1 := time.Date(2024, 3, 31, 0, 0, 0, 0, time.UTC)
		want := &model.Student{
			ID:             10001,
			Name:           "太郎",
			Ruby:           "たろう",
			StudentNumber:  12345,
			DateOfBirth:    date1,
			Address:        "愛知県名古屋市中区三の丸三丁目1番2号",
			ExpirationDate: edate1,
		}
		got, err := i.Create(ctx, want)
		if err != nil {
			t.Fatalf("Could not get mock data: %s\n", err)
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("student mismatch (-want +got):\n%s", diff)
		}
	})
}

func TestGetAll(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	mock := &MockStudent{}
	i := &StudentInteractor{Repo: mock}
	t.Run("全ての学生を返却", func(t *testing.T) {
		mock.GetAllFunc = func(ctx context.Context) ([]*model.Student, error) {
			date1 := time.Date(2003, 6, 27, 15, 0, 0, 0, time.UTC)
			date2 := time.Date(2004, 11, 4, 19, 0, 0, 0, time.UTC)
			edate1 := time.Date(2024, 3, 31, 0, 0, 0, 0, time.UTC)
			edate2 := time.Date(2025, 3, 31, 0, 0, 0, 0, time.UTC)
			wantedStudents := []*model.Student{
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
			}
			return wantedStudents, nil
		}
		date1 := time.Date(2003, 6, 27, 15, 0, 0, 0, time.UTC)
		date2 := time.Date(2004, 11, 4, 19, 0, 0, 0, time.UTC)
		edate1 := time.Date(2024, 3, 31, 0, 0, 0, 0, time.UTC)
		edate2 := time.Date(2025, 3, 31, 0, 0, 0, 0, time.UTC)
		want := []*model.Student{
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
		}
		got, err := i.GetAll(ctx)
		if err != nil {
			t.Fatalf("Could not get mock data: %s\n", err)
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("student mismatch (-want +got):\n%s", diff)
		}
	})
}
