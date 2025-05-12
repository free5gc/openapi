package openapi

// RequestMetricsHook is called after each HTTP Submit.
//
//	method:       HTTP method (GET, POST, ...)
//	path:         URL path pattern (e.g. "/namf-comm/v1/...")
//	status:       HTTP status code (0 if no response)
type RequestMetricsHook func(
	method string,
	path string,
	status int,
	duration float64,
)
