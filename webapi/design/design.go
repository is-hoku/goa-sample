package design

import (
	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
)

var _ = API("sample", func() {
	Title("Sample API")
	Description("Sample API for student information management")
	Server("sample", func() {
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
