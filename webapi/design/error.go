package design

import . "goa.design/goa/v3/dsl"

var CustomErrorType = Type("CustomError", func() {
	Description("CustomError is the error returned error name and message.")
	ErrorName("name", String, "Name of error", func() {
		Example("internal_error")
	})
	Attribute("message", String, "Message of error", func() {
		Example("This is an error message.")
	})
	Required("name", "message")
})
