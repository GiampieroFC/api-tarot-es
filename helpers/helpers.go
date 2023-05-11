package helpers

const (
	Get    = "GET"
	Post   = "POST"
	Put    = "PUT"
	Delete = "DELETE"
)

func MethodChecker(r string, m string) string {
	if m == "" {
		m = Get
	}

	if r != m {
		return "Solo método " + m
	}
	return ""
}
