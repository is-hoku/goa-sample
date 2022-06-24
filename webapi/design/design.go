package design

import . "goa.design/goa/v3/dsl"

var _ = API("students", func() {
	Title("Students API")
	Description("API for student information management")
	Server("student", func() {
		Host("localhost", func() {
			URI("http://localhost:8080")
		})
	})
})

var _ = Service("student", func() {
	Method("students", func() {
		Result(StudentsType)
		HTTP(func() {
			GET("/students")
			Response(StatusOK, StudentsType)
			Response(StatusInternalServerError)
		})
	})
})
