package design

import (
	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
)

var _ = API("students", func() {
	Title("Students API")
	Description("API for student information management")
	Server("student", func() {
		Host("0.0.0.0", func() {
			URI("http://0.0.0.0:8080")
		})
	})
	cors.Origin("http://localhost:8017", func() {
		cors.Expose("X-Time")
		cors.Methods("GET", "POST", "PUT", "DELETE")
		cors.MaxAge(600)
		cors.Credentials()
	})
})

var _ = Service("students", func() {
	Method("get student", func() {
		Description("id から学生を取得する。")
		Result(StudentType)
		Error("internal_error", CustomErrorType)
		Error("not_found", CustomErrorType)
		Payload(func() {
			Attribute("id", Int64, "Student's unique ID")
		})
		HTTP(func() {
			GET("students/{id}")
			Response(StatusOK)
			Response("internal_error", StatusInternalServerError)
			Response("not_found", StatusNotFound)
		})
	})
	Method("get students", func() {
		Description("学籍番号で昇順にソートされた全ての学生を取得する。")
		Result(StudentsType)
		Error("internal_error", CustomErrorType)
		HTTP(func() {
			GET("/students")
			Response(StatusOK)
			Response("internal_error", StatusInternalServerError)
		})
	})
	Method("create student", func() {
		Description("学生を登録する。")
		Result(StudentType)
		Error("internal_error", CustomErrorType)
		Error("bad_request", CustomErrorType)
		HTTP(func() {
			POST("/students")
			Response(StatusOK)
			Response("internal_error", StatusInternalServerError)
			Response("bad_request", StatusBadRequest)
		})
	})
})
