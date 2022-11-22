package interactor

import (
	"github.com/google/wire"
	"github.com/is-hoku/goa-sample/webapi/usecase"
)

var StudentSet = wire.NewSet(
	NewGetStudentByNumber,
	wire.Struct(new(GetStudentByNumberOption), "*"),
	wire.Bind(new(usecase.StudentByNumberGetter), new(*GetStudentByNumber)),
	NewGetStudents,
	wire.Struct(new(GetStudentsOption), "*"),
	wire.Bind(new(usecase.StudentsGetter), new(*GetStudents)),
	NewCreateStudent,
	wire.Struct(new(CreateStudentOption), "*"),
	wire.Bind(new(usecase.StudentCreator), new(*CreateStudent)),
	NewUpdateStudent,
	wire.Struct(new(UpdateStudentOption), "*"),
	wire.Bind(new(usecase.StudentUpdater), new(*UpdateStudent)),
	NewDeleteStudent,
	wire.Struct(new(DeleteStudentOption), "*"),
	wire.Bind(new(usecase.StudentDeleter), new(*DeleteStudent)),
)
