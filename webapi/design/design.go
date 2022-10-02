package design

//a
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
	Security(JWT)

	cors.Origin("http://localhost:8017", func() {
		cors.Expose("X-Time")
		cors.Methods("GET", "POST", "PUT", "DELETE")
		cors.Headers("*")
		cors.MaxAge(600)
		cors.Credentials()
	})
})

var _ = Service("health", func() {
	Method("check", func() {
		Description("ヘルスチェック")
		Result(HealthResultType)
		Error("internal_error", CustomErrorType)
		NoSecurity()
		HTTP(func() {
			GET("/health")
			Response(StatusOK)
			Response("internal_error", StatusInternalServerError)
		})
	})
})

var HealthResultType = ResultType("application/vnd.health+json", "HealthResult", func() {
	Description("OK.")
	Attribute("message", String, "health message", func() {
		Example("OK.")
	})
	Required("message")
})

var JWT = JWTSecurity("jwt", func() {
	Description("Use firebase Authentication")
	Scope("api:read", "Read access")
	Scope("api:write", "Write access")
})

var Authorization = Type("Authorization", func() {
	Token("Authorization", String, "Firebase JWT Token")
	Required("Authorization")
})
