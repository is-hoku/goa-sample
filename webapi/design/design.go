package design

import (
	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
)

var _ = API("student", func() {
	Title("Student API")
	Description("API for student information management")
	Server("student", func() {
		Host("0.0.0.0", func() {
			URI("http://0.0.0.0:8080")
		})
	})

	cors.Origin("http://localhost:8017", func() {
		cors.Expose("X-Time")
		cors.Methods("GET", "POST", "PUT", "DELETE")
		cors.Headers("*")
		cors.MaxAge(600)
		cors.Credentials()
	})
})

var _ = Service("student", func() {
	Method("get_student", func() {
		Description("学籍番号から学生を取得する。")
		Result(StudentType)
		Error("internal_error", CustomErrorType)
		Error("not_found", CustomErrorType)
		Payload(func() {
			Attribute("student_number", Int64, "Student's unique number")
		})
		HTTP(func() {
			GET("students/{student_number}")
			Response(StatusOK)
			Response("internal_error", StatusInternalServerError)
			Response("not_found", StatusNotFound)
		})
	})
	Method("get_students", func() {
		Description("学籍番号で昇順にソートされた全ての学生を取得する。")
		Result(StudentsType)
		Error("internal_error", CustomErrorType)
		HTTP(func() {
			GET("/students")
			Response(StatusOK)
			Response("internal_error", StatusInternalServerError)
		})
	})
	Method("create_student", func() {
		Description("学生を登録する。")
		Result(StudentType)
		Error("internal_error", CustomErrorType)
		Error("bad_request", CustomErrorType)
		Payload(StudentBodyType)
		HTTP(func() {
			POST("/students")
			Body(StudentBodyType)
			Response(StatusCreated)
			Response("internal_error", StatusInternalServerError)
			Response("bad_request", StatusBadRequest)
		})
	})
})
