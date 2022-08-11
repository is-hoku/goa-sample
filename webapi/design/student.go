package design

import . "goa.design/goa/v3/dsl"

var _ = Service("student", func() {
	Method("get_student", func() {
		Description("学籍番号から学生を取得する。")
		Result(StudentType)
		Error("internal_error", CustomErrorType)
		Error("not_found", CustomErrorType)
		Error("bad_request", CustomErrorType)
		Payload(func() {
			Attribute("student_number", UInt32, "Student's unique number")
		})
		HTTP(func() {
			GET("students/{student_number}")
			Response(StatusOK)
			Response("internal_error", StatusInternalServerError)
			Response("not_found", StatusNotFound)
			Response("bad_request", StatusBadRequest)
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
			Response(StatusOK)
			Response("internal_error", StatusInternalServerError)
			Response("bad_request", StatusBadRequest)
		})
	})
})

var StudentType = ResultType("application/vnd.student+json", "Student", func() {
	Description("One student")
	Attributes(func() {
		Attribute("id", UInt64, "学生を一意に表す ID", func() {
			Example(1)
		})
		Attribute("name", String, "学生の氏名", func() {
			Example("鈴木太郎")
		})
		Attribute("ruby", String, "学生の氏名のフリガナ", func() {
			Example("スズキタロウ")
		})
		Attribute("student_number", UInt32, "学生の学籍番号", func() {
			Example(12345)
		})
		Attribute("date_of_birth", String, "学生の生年月日 (RFC3339)", func() {
			Format(FormatDateTime)
			Example("2022-04-01T13:30:00+09:00")
		})
		Attribute("address", String, "学生の住所", func() {
			Example("名古屋市中区三の丸三丁目1番2号")
		})
		Attribute("expiration_date", String, "学生証の有効期間 (RFC3339)", func() {
			Format(FormatDateTime)
			Example("2027-03-31T00:00:00+09:00")
		})
		Required("id", "name", "ruby", "student_number", "date_of_birth", "address", "expiration_date")
		View("default", func() {
			Attribute("id")
			Attribute("name")
			Attribute("ruby")
			Attribute("student_number")
			Attribute("date_of_birth")
			Attribute("address")
			Attribute("expiration_date")
		})
	})
})

var StudentsType = ResultType("application/vnd.students+json", "Students", func() {
	Description("All students")
	Attribute("students", ArrayOf(StudentType), func() {
		ArrayOf(StudentType)
	})
	Required("students")
	View("default", func() {
		Attribute("students")
	})
})

var StudentBodyType = Type("StudentBody", func() {
	Description("Student Body")
	Attribute("name", String, "学生の氏名", func() {
		Example("鈴木太郎")
	})
	Attribute("ruby", String, "学生の氏名のフリガナ", func() {
		Example("スズキタロウ")
	})
	Attribute("student_number", UInt32, "学生の学籍番号", func() {
		Example(12345)
	})
	Attribute("date_of_birth", String, "学生の生年月日 (RFC3339)", func() {
		Format(FormatDateTime)
		Example("2022-04-01T13:30:00+09:00")
	})
	Attribute("address", String, "学生の住所", func() {
		Example("名古屋市中区三の丸三丁目1番2号")
	})
	Attribute("expiration_date", String, "学生証の有効期間 (RFC3339)", func() {
		Format(FormatDateTime)
		Example("2027-03-31T00:00:00+09:00")
	})
	Required("name", "ruby", "student_number", "date_of_birth", "address", "expiration_date")
})
